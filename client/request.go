package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kenael/simpledeviceinventory/env"
)

func SendSystem(machineID string, data []byte) {
	url := fmt.Sprintf("%s://%s:%s/system/%s",
		env.GetEnv("SERVER_PROTO", "http"),
		env.GetEnv("SERVER_URL", "localhost"),
		env.GetEnv("SERVER_PORT", "3000"),
		machineID,
	)
	fmt.Println("URL:>", url)

	var jsonStr = data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("Create Request failed")
	}
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
	url := fmt.Sprintf("%s://%s:%s/system/%s/packages",
		env.GetEnv("SERVER_PROTO", "http"),
		env.GetEnv("SERVER_URL", "localhost"),
		env.GetEnv("SERVER_PORT", "3000"),
		machineID,
	)
	fmt.Println("URL:>", url)

	var jsonStr = data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("Create Request failed")
	}
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
