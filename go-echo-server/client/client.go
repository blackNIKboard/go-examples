package main

import (
	"bytes"
	"github.com/blackNIKboard/go-examples/go-echo-server/shared"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://0.0.0.0:8080/test"
	method := "POST"

	request := shared.Person{
		Name:    "Test",
		Surname: "B",
		Age:     21,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
