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
	Version  string
	User     string
	Computer string
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

func GetUsages(value string) (map[string]Usage, error) {
	reg_usage_compiled := regexp.MustCompile(statics.REG_USAGE)
	usage_matchs := reg_usage_compiled.FindAllString(value, -1)

	usages := make(map[string]Usage)
	if len(usage_matchs) == 0 {
		return usages, fmt.Errorf("NO MATCH FOUND FOR USAGES")
	}

	for i := 0; i < len(usage_matchs); i++ {
		result := getNamedGroups(reg_usage_compiled, usage_matchs[i])

		unique_key := fmt.Sprintf("%s-%s", result["computer"], result["product"])

		usage_obj := Usage{
			Product:  result["product"],
			Version:  result["version"],
			User:     result["user"],
			Computer: result["computer"],
		}

		if _, ok := usages[unique_key]; ok {
			continue
		}

		// add to dictionary
		usages[unique_key] = usage_obj
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
		totals[total_obj.Product] = total_obj
	}
	return totals, nil
}

func GetMetrics(value string) (string, error) {
	usages, err := GetUsages(value)
	if err != nil {
		return "", err
	}

	tag := "rlm"
	if _tag := os.Getenv(statics.LIC_TAG); _tag != "" {
		tag = _tag
	}

	totals, err := GetTotals(value)
	if err != nil {
		return "", err
	}

	// sort the usages keys, required for the tests
	sorted_keys := make([]string, 0, len(usages))
	for k := range usages {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	str := fmt.Sprintf("\n# HELP %s_license_usage Used %s licenses\n# TYPE %s_license_usage gauge\n", tag, tag, tag)
	for _, key := range sorted_keys {
		usage := usages[key]

		str += fmt.Sprintf(
			"%s_license_info_%s{computer=\"%s\", product=\"%s\", version=\"%s\", user=\"%s\"} 1.0\n",
			tag,
			usage.Product,
			usage.Computer,
			usage.Product,
			usage.Version,
			usage.User,
		)
	}

	// sort the usages keys, required for the tests
	sorted_keys = make([]string, 0, len(totals))
	for k := range totals {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	str += fmt.Sprintf("# HELP %s_license_total Total %s licenses\n# TYPE %s_license_total gauge\n", tag, tag, tag)
	for _, product := range sorted_keys {
		total := totals[product]
		str += fmt.Sprintf(
			"%s_licenses_total{product=\"%s\", version=\"%s\"} %d.0\n",
			tag,
			product,
			total.Version,
			total.Count,
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
