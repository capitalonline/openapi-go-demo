package test

import (
	"fmt"
	"openapi-go-demo/example"
	"testing"
)

func TestDescribeInstance(t *testing.T) {
	result := example.DescribeInstances("")
	fmt.Printf("%v", result)
}

func TestCreateInstance(t *testing.T) {
	t.Log("Start Test Create Instance")
	RegionId := ""
	VdcId := ""
	InstanceName := ""
	InstanceType := ""
	ImageId := ""
	Password := ""

	res := example.CreateInstance(RegionId, VdcId, InstanceName, InstanceType, ImageId, Password, 4, 4, 1)
	fmt.Println(res)
}

func TestDeleteInstance(t *testing.T) {
	instanceId := ""
	var param = []string{instanceId}
	res := example.DeleteInstance(param)
	fmt.Println(res)
}

func TestStopInstance(t *testing.T) {
	instanceId := ""
	res := example.StopInstance(instanceId)
	fmt.Println(res)
}

func TestStartInstance(t *testing.T) {
	instanceId := ""
	res := example.StartInstance(instanceId)
	fmt.Println(res)
}

func TestRebootInstance(t *testing.T) {
	instanceId := ""
	res := example.RebootInstance(instanceId)
	fmt.Println(res)
}

func TestModifyInstanceSpec(t *testing.T) {
	instanceId := ""
	res := example.ModifyInstanceSpec(instanceId, 4, 8)
	fmt.Println(res)
}

func TestCreateDisk(t *testing.T) {
	instanceId := ""
	res := example.CreateDisk(instanceId)
	fmt.Println(res)
}

func TestResizeDisk(t *testing.T) {
	instanceId := ""
	diskId := ""
	size := 200
	res := example.ResizeDisk(instanceId, diskId, size)
	fmt.Println(res)
}

func TestDeleteDisk(t *testing.T) {
	instanceId := ""
	diskIds := []string{""}
	res := example.DeleteDisk(instanceId, diskIds)
	fmt.Println(res)
}

func TestResetImageByPassword(t *testing.T) {
	instanceId := ""
	imageId := ""
	Password := ""
	res := example.ResetImage(instanceId, imageId, Password, "")
	fmt.Println(res)
}

func TestResetImageByPublicKey(t *testing.T) {
	instanceId := ""
	imageId := ""
	PublicKey := ""
	res := example.ResetImage(instanceId, imageId, "", PublicKey)
	fmt.Println(res)
}

func TestModifyIpAddress(t *testing.T) {
	instanceId := ""
	interfaceId := ""
	address := ""
	res := example.ModifyIpAddress(instanceId, interfaceId, address)
	fmt.Println(res)
}

func TestModifyIpAddressToNull(t *testing.T) {
	instanceId := ""
	interfaceId := ""
	address := ""
	res := example.ModifyIpAddress(instanceId, interfaceId, address)
	fmt.Println(res)
}

func TestModifyInstanceName(t *testing.T) {
	instanceId := ""
	instanceName := ""
	res := example.ModifyInstanceName(instanceId, instanceName)
	fmt.Println(res)
}

func TestGetHistoryInstance(t *testing.T) {
	res := example.GetHistoryInstance("2020-12-30", 10, 1)
	fmt.Println(res)
}
