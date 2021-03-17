package test

import (
	"fmt"
	"openapi-go-demo/example"
	"testing"
)

func TestDescribeVdc(t *testing.T) {
	result := example.DescribeVdc("", "", "")
	fmt.Printf("%v", result)
}

func TestCreateVdc(t *testing.T) {
	param := example.PublicNetworkInfo{
		Name:          "",
		Type:          "",
		BillingMethod: "",
		Qos:           5,
		IPNum:         4,
	}
	result := example.CreateVdc("", "", param)
	fmt.Println(result)
}

func TestAddPublicIp(t *testing.T) {
	result := example.AddPublicIp("", "")
	fmt.Println(result)
}

func TestDeletePublicIp(t *testing.T) {
	result := example.DeletePublicIp("")
	fmt.Println(result)
}

func TestDescribeBandwidthTraffic(t *testing.T) {
	result := example.DescribeBandwidthTraffic("")
	fmt.Println(result)
}

func TestBiVdcList(t *testing.T) {
	result := example.BiVdcList("", "", "")
	fmt.Println(result)
}

func TestBiVdcFlow(t *testing.T) {
	result := example.BiVdcFlow("", "", "")
	fmt.Println(result)
}

func TestBiGpnBandwidth(t *testing.T) {
	result := example.BiGpnBandwidth("", "", "")
	fmt.Println(result)
}

func TestBiVdcPublicIp(t *testing.T) {
	result := example.BiVdcPublicIp("", "", "")
	fmt.Println(result)
}
