package utils

type Tests struct {
	VPNTest  VPNTest  `json:"vpn_test"`
	PingTest PingTest `json:"ping_test"`
}

type PingTest struct {
	URL         string  `json:"url,omitempty"`
	Transmitted int     `json:"transmitted_packets,omitempty"`
	Received    int     `json:"received_packets,omitempty"`
	Loss        float64 `json:"loss_packets,omitempty"`
}

type VPNTest struct {
	Status string `json:"status"`
}
