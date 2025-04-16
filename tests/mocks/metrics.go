package mocks

const PROMETHEUS_METRICS = `
# HELP rlm_license_usage Used rlm licenses
# TYPE rlm_license_usage gauge
rlm_license_info_nuke_i{computer="mocked-computer-01", product="nuke_i", version="v2024.0331", user="user1.lastname"} 1.0
rlm_license_info_nukex_i{computer="mocked-computer-01", product="nukex_i", version="v2024.0701", user="user1.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-02", product="nuke_i", version="v2024.0701", user="user2.lastname"} 1.0
rlm_license_info_nukex_i{computer="mocked-computer-02", product="nukex_i", version="v2024.0331", user="user2.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-03", product="nuke_i", version="v2024.0701", user="user3.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-04", product="nuke_i", version="v2024.0701", user="user4.lastname"} 1.0
rlm_license_info_nuke_r{computer="mocked-computer-05", product="nuke_r", version="v2024.0701", user="user1.lastname"} 1.0
rlm_license_info_nuke_r{computer="mocked-computer-06", product="nuke_r", version="v2024.0701", user="user1.lastname"} 1.0
# HELP rlm_license_total Total rlm licenses
# TYPE rlm_license_total gauge
rlm_licenses_total{product="hieroplayer_i", version="v2024.0701"} 23.0
rlm_licenses_total{product="nuke_i", version="v2024.0416"} 27.0
rlm_licenses_total{product="nuke_r", version="v2024.0701"} 25.0
rlm_licenses_total{product="nukex_i", version="v2024.0701"} 4.0
rlm_licenses_total{product="nukexassist_i", version="v2024.0701"} 8.0
`

const PROMETHEUS_METRICS_NO_LIC_NUKE_R = `
# HELP rlm_license_usage Used rlm licenses
# TYPE rlm_license_usage gauge
rlm_license_info_nuke_i{computer="mocked-computer-01", product="nuke_i", version="v2024.0331", user="user1.lastname"} 1.0
rlm_license_info_nukex_i{computer="mocked-computer-01", product="nukex_i", version="v2024.0701", user="user1.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-02", product="nuke_i", version="v2024.0701", user="user2.lastname"} 1.0
rlm_license_info_nukex_i{computer="mocked-computer-02", product="nukex_i", version="v2024.0331", user="user2.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-03", product="nuke_i", version="v2024.0701", user="user3.lastname"} 1.0
rlm_license_info_nuke_i{computer="mocked-computer-04", product="nuke_i", version="v2024.0701", user="user4.lastname"} 1.0
# HELP rlm_license_total Total rlm licenses
# TYPE rlm_license_total gauge
rlm_licenses_total{product="hieroplayer_i", version="v2024.0701"} 23.0
rlm_licenses_total{product="nuke_i", version="v2024.0416"} 27.0
rlm_licenses_total{product="nuke_r", version="v2024.0701"} 25.0
rlm_licenses_total{product="nukex_i", version="v2024.0701"} 4.0
rlm_licenses_total{product="nukexassist_i", version="v2024.0701"} 8.0
`
