package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendSystem(machineID string, data []byte) {
	url := fmt.Sprintf("http://192.168.0.106:3000/system/%s", machineID)
	fmt.Println("URL:>", url)

	var jsonStr = data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func SendPackages(machineID string, data []byte) {
	url := fmt.Sprintf("http://192.168.0.106:3000/system/%s/packages", machineID)
	fmt.Println("URL:>", url)

	var jsonStr = data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}