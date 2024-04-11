package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"

	"github.com/OlivierArgentieri/PrometheusExporter/responses"
	"github.com/OlivierArgentieri/PrometheusExporter/statics"
)

type Usage struct {
	Product  string
	Versions []string
	Users    []string
	Workers  []string
	Count    int
}

type Total struct {
	Product string
	Version string
	Count   int
}

// TODO: Move to a helper package
func getNamedGroups(reg *regexp.Regexp, value string) map[string]string {
	// create a  dictionary to hold the named groups
	result := make(map[string]string)
	// get the named groups
	match := reg.FindStringSubmatch(value)

	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sliceToString(slice []string) string {
	var str string
	for i := 0; i < len(slice); i++ {
		str += slice[i]
		if i < len(slice)-1 {
			str += ","
		}
	}
	return str
}

func GetUsages(value string) (map[string]Usage, error) {
	reg_usage_compiled := regexp.MustCompile(statics.REG_USAGE)
	usage_matchs := reg_usage_compiled.FindAllString(value, -1)

	usages := make(map[string]Usage)
	if len(usage_matchs) == 0 {
		return usages, fmt.Errorf("NO MATCH FOUND FOR USAGES")
	}

	for i := 0; i < len(usage_matchs); i++ {
		result := getNamedGroups(reg_usage_compiled, usage_matchs[i])

		if usage, ok := usages[result["product"]]; ok {
			// ignore if worker already exists because it's one license per worker
			if stringInSlice(result["worker"], usage.Workers) {
				continue
			}

			usage.Workers = append(usage.Workers, result["worker"])

			if !stringInSlice(result["user"], usage.Users) {
				usage.Users = append(usage.Users, result["user"])
			}

			if !stringInSlice(result["version"], usage.Versions) {
				usage.Versions = append(usage.Versions, result["version"])
			}

			usages[result["product"]] = Usage{
				Product:  result["product"],
				Versions: usage.Versions,
				Users:    usage.Users,
				Workers:  usage.Workers,
				Count:    usage.Count + 1,
			}
			continue
		}

		usage_obj := Usage{
			Product:  result["product"],
			Versions: []string{result["version"]},
			Users:    []string{result["user"]},
			Workers:  []string{result["worker"]},
			Count:    1,
		}
		// add to dictionary
		usages[usage_obj.Product] = usage_obj
	}
	return usages, nil
}

func GetTotals(value string) (map[string]Total, error) {
	reg_total_compiled := regexp.MustCompile(statics.REG_TOTAL)
	total_matchs := reg_total_compiled.FindAllString(value, -1)

	totals := make(map[string]Total)
	if len(total_matchs) == 0 {
		return totals, fmt.Errorf("NO MATCH FOUND FOR TOTAL")
	}

	for i := 0; i < len(total_matchs); i++ {
		result := getNamedGroups(reg_total_compiled, total_matchs[i])

		count, err := strconv.Atoi(result["count"])
		if err != nil {
			return nil, fmt.Errorf("ERROR CONVERTING STRING TO INT FROM RLM")
		}

		// ignore if key already exists
		if total, ok := totals[result["product"]]; ok {
			// update the count
			totals[result["product"]] = Total{
				Product: result["product"],
				Version: result["version"],
				Count:   total.Count + count,
			}
			continue
		}

		total_obj := Total{
			Product: result["product"],
			Version: result["version"],
			Count:   count,
		}
		// add to dictionary
		totals[total_obj.Product] = total_obj
	}
	return totals, nil
}

func GetMetrics(value string) (string, error) {
	usage, err := GetUsages(value)
	if err != nil {
		return "", err
	}

	total, err := GetTotals(value)
	if err != nil {
		return "", err
	}

	// Sort dict key
	// TODO: Move to a helper function/package
	total_keys := make([]string, 0, len(total))
	for k := range total {
		total_keys = append(total_keys, k)
	}
	sort.Strings(total_keys)

	str := ""
	for key := range total_keys {
		key := total_keys[key]
		value := total[key]
		str += fmt.Sprintf("# HELP %s_license_usage %s_license_usage\n", key, key)
		str += fmt.Sprintf("# TYPE %s_license_usage gauge\n", key)

		versions := "none"
		users := "none"
		workers := "none"
		count := 0

		if usage, ok := usage[key]; ok {
			versions = sliceToString(usage.Versions)
			users = sliceToString(usage.Users)
			workers = sliceToString(usage.Workers)
			count = usage.Count
		}

		str += fmt.Sprintf(
			"rlm_license_info_%s{id=\"%s_usage\", product=\"%s\", versions=\"%s\", users=\"%s\", workers=\"%s\", count=\"%d\"} %d\n",
			key,
			key,
			value.Product,
			versions,
			users,
			workers,
			count,
			count,
		)
	}

	for key := range total_keys {
		key := total_keys[key]
		value := total[key]
		str += fmt.Sprintf("# HELP %s_total_licenses %s_total_licenses\n", key, key)
		str += fmt.Sprintf("# TYPE %s_total_licenses gauge\n", key)
		str += fmt.Sprintf(
			"rlm_license_info_%s{id=\"%s_total\", product=\"%s\", version=\"%s\", count=\"%d\"} %d\n",
			key,
			key,
			value.Product,
			value.Version,
			value.Count,
			value.Count,
		)
	}

	return str, nil
}

func (server *Server) GetLicenses(w http.ResponseWriter, r *http.Request) {

	// get the rlmutil path
	rlmutil_root := os.Getenv(statics.RLMUTIL_BIN_ROOT)

	// check if folder exists
	if _, err := os.Stat(rlmutil_root); os.IsNotExist(err) {
		fmt.Println(err)
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	// check for rlmutil binary file
	rlmutil_bin_path := fmt.Sprintf(`%s/rlmutil`, rlmutil_root)

	rlm_port := os.Getenv(statics.RLM_SERVER_PORT)
	rlm_host := os.Getenv(statics.RLM_SERVER_HOST)
	cmd := fmt.Sprintf(`%s rlmstat -a -c %s@%s`, rlmutil_bin_path, rlm_port, rlm_host)

	get_processes := exec.Command("bash", "-c", cmd)
	stdout, _ := get_processes.CombinedOutput()
	metrics, err := GetMetrics(string(stdout))
	if err != nil {
		fmt.Println(err)
	}

	responses.TEXT(w, http.StatusOK, metrics)
}
