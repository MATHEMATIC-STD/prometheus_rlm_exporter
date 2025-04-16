package controllertests

import (
	"fmt"
	"testing"

	"github.com/OlivierArgentieri/PrometheusExporter/controllers"
	"github.com/OlivierArgentieri/PrometheusExporter/tests/mocks"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetUsages(t *testing.T) {
	result, err := controllers.GetUsages(mocks.STD_OUT_EMPTY_NUKE_R)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.Equal(t, result, map[string]controllers.Usage{
		"mocked-computer-01-nuke_i": {
			Computer: "mocked-computer-01",
			Product:  "nuke_i",
			Version:  "v2024.0331",
			User:     "user1.lastname",
		},
		"mocked-computer-01-nukex_i": {
			Computer: "mocked-computer-01",
			Product:  "nukex_i",
			Version:  "v2024.0701",
			User:     "user1.lastname",
		},
		"mocked-computer-02-nuke_i": {
			Computer: "mocked-computer-02",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user2.lastname",
		},
		"mocked-computer-02-nukex_i": {
			Computer: "mocked-computer-02",
			Product:  "nukex_i",
			Version:  "v2024.0331",
			User:     "user2.lastname",
		},
		"mocked-computer-03-nuke_i": {
			Computer: "mocked-computer-03",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user3.lastname",
		},
		"mocked-computer-04-nuke_i": {
			Computer: "mocked-computer-04",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user4.lastname",
		},
	})

	result, err = controllers.GetUsages(mocks.STD_OUT)
	if err != nil {
		assert.Equal(t, err, nil)
	}

	assert.Equal(t, result, map[string]controllers.Usage{
		"mocked-computer-01-nuke_i": {
			Computer: "mocked-computer-01",
			Product:  "nuke_i",
			Version:  "v2024.0331",
			User:     "user1.lastname",
		},
		"mocked-computer-01-nukex_i": {
			Computer: "mocked-computer-01",
			Product:  "nukex_i",
			Version:  "v2024.0701",
			User:     "user1.lastname",
		},
		"mocked-computer-02-nuke_i": {
			Computer: "mocked-computer-02",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user2.lastname",
		},
		"mocked-computer-02-nukex_i": {
			Computer: "mocked-computer-02",
			Product:  "nukex_i",
			Version:  "v2024.0331",
			User:     "user2.lastname",
		},
		"mocked-computer-03-nuke_i": {
			Computer: "mocked-computer-03",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user3.lastname",
		},
		"mocked-computer-04-nuke_i": {
			Computer: "mocked-computer-04",
			Product:  "nuke_i",
			Version:  "v2024.0701",
			User:     "user4.lastname",
		},
		"mocked-computer-05-nuke_r": {
			Computer: "mocked-computer-05",
			Product:  "nuke_r",
			Version:  "v2024.0701",
			User:     "user1.lastname",
		},
		"mocked-computer-06-nuke_r": {
			Computer: "mocked-computer-06",
			Product:  "nuke_r",
			Version:  "v2024.0701",
			User:     "user1.lastname",
		},
	})

	_, err = controllers.GetUsages(mocks.STD_OUT_WRONG)
	assert.Equal(t, err, fmt.Errorf("NO MATCH FOUND FOR USAGES"))
}

func TestGetTotals(t *testing.T) {
	result, err := controllers.GetTotals(mocks.STD_OUT)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.Equal(t, result, map[string]controllers.Total{
		"hieroplayer_i": {
			Product: "hieroplayer_i",
			Version: "v2024.0701",
			Count:   23,
		},
		"nuke_i": {
			Product: "nuke_i",
			Version: "v2024.0416",
			Count:   27,
		},
		"nuke_r": {
			Product: "nuke_r",
			Version: "v2024.0701",
			Count:   25,
		},
		"nukex_i": {
			Product: "nukex_i",
			Version: "v2024.0701",
			Count:   4,
		},
		"nukexassist_i": {
			Product: "nukexassist_i",
			Version: "v2024.0701",
			Count:   8,
		},
	})

	_, err = controllers.GetTotals(mocks.STD_OUT_WRONG)
	assert.Equal(t, err, fmt.Errorf("NO MATCH FOUND FOR TOTAL"))
}

func TestGetMetrics(t *testing.T) {
	metrics_result, err := controllers.GetMetrics(mocks.STD_OUT)
	if err != nil {
		assert.Equal(t, err, nil)
	}

	assert.Equal(t, metrics_result, mocks.PROMETHEUS_METRICS)

	metrics_result, err = controllers.GetMetrics(mocks.STD_OUT_EMPTY_NUKE_R)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.Equal(t, metrics_result, mocks.PROMETHEUS_METRICS_NO_LIC_NUKE_R)
}
