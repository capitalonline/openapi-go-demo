package test

import (
	"openapi-go-demo/example"
	"fmt"
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
