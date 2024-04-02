package mocks

const PROMETHEUS_METRICS = `# HELP nuke_rlm_hieroplayer_i_license_usage nuke_rlm_hieroplayer_i_license_usage
# TYPE nuke_rlm_hieroplayer_i_license_usage gauge
rlm_license_info_hieroplayer_i{id="hieroplayer_i_usage", product="hieroplayer_i", versions="none", users="none", workers="none", count="0"} 0
# HELP nuke_rlm_nuke_i_license_usage nuke_rlm_nuke_i_license_usage
# TYPE nuke_rlm_nuke_i_license_usage gauge
rlm_license_info_nuke_i{id="nuke_i_usage", product="nuke_i", versions="v2024.0331,v2024.0701", users="user1.lastname,user2.lastname,user3.lastname,user4.lastname", workers="mocked-computer-01,mocked-computer-02,mocked-computer-03,mocked-computer-04", count="4"} 4
# HELP nuke_rlm_nuke_r_license_usage nuke_rlm_nuke_r_license_usage
# TYPE nuke_rlm_nuke_r_license_usage gauge
rlm_license_info_nuke_r{id="nuke_r_usage", product="nuke_r", versions="v2024.0701", users="user1.lastname", workers="mocked-computer-05,mocked-computer-06", count="2"} 2
# HELP nuke_rlm_nukex_i_license_usage nuke_rlm_nukex_i_license_usage
# TYPE nuke_rlm_nukex_i_license_usage gauge
rlm_license_info_nukex_i{id="nukex_i_usage", product="nukex_i", versions="v2024.0331,v2024.0701", users="user2.lastname,user1.lastname", workers="mocked-computer-02,mocked-computer-01", count="2"} 2
# HELP nuke_rlm_nukexassist_i_license_usage nuke_rlm_nukexassist_i_license_usage
# TYPE nuke_rlm_nukexassist_i_license_usage gauge
rlm_license_info_nukexassist_i{id="nukexassist_i_usage", product="nukexassist_i", versions="none", users="none", workers="none", count="0"} 0
# HELP nuke_rlm_hieroplayer_i_total_licenses nuke_rlm_hieroplayer_i_total_licenses
# TYPE nuke_rlm_hieroplayer_i_total_licenses gauge
rlm_license_info_hieroplayer_i{id="hieroplayer_i_total", product="hieroplayer_i", version="v2024.0701", count="23"} 23
# HELP nuke_rlm_nuke_i_total_licenses nuke_rlm_nuke_i_total_licenses
# TYPE nuke_rlm_nuke_i_total_licenses gauge
rlm_license_info_nuke_i{id="nuke_i_total", product="nuke_i", version="v2024.0416", count="27"} 27
# HELP nuke_rlm_nuke_r_total_licenses nuke_rlm_nuke_r_total_licenses
# TYPE nuke_rlm_nuke_r_total_licenses gauge
rlm_license_info_nuke_r{id="nuke_r_total", product="nuke_r", version="v2024.0701", count="25"} 25
# HELP nuke_rlm_nukex_i_total_licenses nuke_rlm_nukex_i_total_licenses
# TYPE nuke_rlm_nukex_i_total_licenses gauge
rlm_license_info_nukex_i{id="nukex_i_total", product="nukex_i", version="v2024.0701", count="4"} 4
# HELP nuke_rlm_nukexassist_i_total_licenses nuke_rlm_nukexassist_i_total_licenses
# TYPE nuke_rlm_nukexassist_i_total_licenses gauge
rlm_license_info_nukexassist_i{id="nukexassist_i_total", product="nukexassist_i", version="v2024.0701", count="8"} 8
`

const PROMETHEUS_METRICS_NO_LIC_NUKE_R = `# HELP nuke_rlm_hieroplayer_i_license_usage nuke_rlm_hieroplayer_i_license_usage
# TYPE nuke_rlm_hieroplayer_i_license_usage gauge
rlm_license_info_hieroplayer_i{id="hieroplayer_i_usage", product="hieroplayer_i", versions="none", users="none", workers="none", count="0"} 0
# HELP nuke_rlm_nuke_i_license_usage nuke_rlm_nuke_i_license_usage
# TYPE nuke_rlm_nuke_i_license_usage gauge
rlm_license_info_nuke_i{id="nuke_i_usage", product="nuke_i", versions="v2024.0331,v2024.0701", users="user1.lastname,user2.lastname,user3.lastname,user4.lastname", workers="mocked-computer-01,mocked-computer-02,mocked-computer-03,mocked-computer-04", count="4"} 4
# HELP nuke_rlm_nuke_r_license_usage nuke_rlm_nuke_r_license_usage
# TYPE nuke_rlm_nuke_r_license_usage gauge
rlm_license_info_nuke_r{id="nuke_r_usage", product="nuke_r", versions="none", users="none", workers="none", count="0"} 0
# HELP nuke_rlm_nukex_i_license_usage nuke_rlm_nukex_i_license_usage
# TYPE nuke_rlm_nukex_i_license_usage gauge
rlm_license_info_nukex_i{id="nukex_i_usage", product="nukex_i", versions="v2024.0331,v2024.0701", users="user2.lastname,user1.lastname", workers="mocked-computer-02,mocked-computer-01", count="2"} 2
# HELP nuke_rlm_nukexassist_i_license_usage nuke_rlm_nukexassist_i_license_usage
# TYPE nuke_rlm_nukexassist_i_license_usage gauge
rlm_license_info_nukexassist_i{id="nukexassist_i_usage", product="nukexassist_i", versions="none", users="none", workers="none", count="0"} 0
# HELP nuke_rlm_hieroplayer_i_total_licenses nuke_rlm_hieroplayer_i_total_licenses
# TYPE nuke_rlm_hieroplayer_i_total_licenses gauge
rlm_license_info_hieroplayer_i{id="hieroplayer_i_total", product="hieroplayer_i", version="v2024.0701", count="23"} 23
# HELP nuke_rlm_nuke_i_total_licenses nuke_rlm_nuke_i_total_licenses
# TYPE nuke_rlm_nuke_i_total_licenses gauge
rlm_license_info_nuke_i{id="nuke_i_total", product="nuke_i", version="v2024.0416", count="27"} 27
# HELP nuke_rlm_nuke_r_total_licenses nuke_rlm_nuke_r_total_licenses
# TYPE nuke_rlm_nuke_r_total_licenses gauge
rlm_license_info_nuke_r{id="nuke_r_total", product="nuke_r", version="v2024.0701", count="25"} 25
# HELP nuke_rlm_nukex_i_total_licenses nuke_rlm_nukex_i_total_licenses
# TYPE nuke_rlm_nukex_i_total_licenses gauge
rlm_license_info_nukex_i{id="nukex_i_total", product="nukex_i", version="v2024.0701", count="4"} 4
# HELP nuke_rlm_nukexassist_i_total_licenses nuke_rlm_nukexassist_i_total_licenses
# TYPE nuke_rlm_nukexassist_i_total_licenses gauge
rlm_license_info_nukexassist_i{id="nukexassist_i_total", product="nukexassist_i", version="v2024.0701", count="8"} 8
`
