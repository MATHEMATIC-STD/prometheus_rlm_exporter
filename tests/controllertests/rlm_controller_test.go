package controllertests

import (
	"fmt"
	"testing"

	"github.com/OlivierArgentieri/NukeRlmPrometheusExporter/controllers"
	"github.com/OlivierArgentieri/NukeRlmPrometheusExporter/tests/mocks"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetUsages(t *testing.T) {

	result, err := controllers.GetUsages(mocks.STD_OUT_EMPTY_NUKE_R)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.Equal(t, result, map[string]controllers.Usage{
		"nuke_i": {
			Product:  "nuke_i",
			Versions: []string{"v2024.0331", "v2024.0701"},
			Users:    []string{"user1.lastname", "user2.lastname", "user3.lastname", "user4.lastname"},
			Workers:  []string{"mocked-computer-01", "mocked-computer-02", "mocked-computer-03", "mocked-computer-04"},
			Count:    4,
		},
		"nukex_i": {
			Product:  "nukex_i",
			Versions: []string{"v2024.0331", "v2024.0701"},
			Users:    []string{"user2.lastname", "user1.lastname"},
			Workers:  []string{"mocked-computer-02", "mocked-computer-01"},
			Count:    2,
		},
	})

	result, err = controllers.GetUsages(mocks.STD_OUT)
	if err != nil {
		assert.Equal(t, err, nil)
	}

	assert.Equal(t, result, map[string]controllers.Usage{
		"nuke_i": {
			Product:  "nuke_i",
			Versions: []string{"v2024.0331", "v2024.0701"},
			Users:    []string{"user1.lastname", "user2.lastname", "user3.lastname", "user4.lastname"},
			Workers:  []string{"mocked-computer-01", "mocked-computer-02", "mocked-computer-03", "mocked-computer-04"},
			Count:    4,
		},
		"nuke_r": {
			Product:  "nuke_r",
			Versions: []string{"v2024.0701"},
			Users:    []string{"user1.lastname"},
			Workers:  []string{"mocked-computer-05", "mocked-computer-06"},
			Count:    2,
		},
		"nukex_i": {
			Product:  "nukex_i",
			Versions: []string{"v2024.0331", "v2024.0701"},
			Users:    []string{"user2.lastname", "user1.lastname"},
			Workers:  []string{"mocked-computer-02", "mocked-computer-01"},
			Count:    2,
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
