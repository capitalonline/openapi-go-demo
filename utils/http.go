package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CommonReturn struct {
	Code       string
	Data       interface{}
	Message    string
	PageCount  int
	PageNumber int
	TaskId     string
}

func DoHttpPost(action, url, method string, reqBody []byte) CommonReturn {
	fmt.Println("ReqBody: ", string(reqBody))

	signedURL := getSignedURL(action, url, method, reqBody)
	req, _ := http.NewRequest(method, signedURL, bytes.NewBuffer(reqBody))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("request Url", signedURL)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	commonReturn := CommonReturn{}
	err = json.Unmarshal(body, &commonReturn)
	if err != nil {
		panic(err)
	}
	return commonReturn
}
