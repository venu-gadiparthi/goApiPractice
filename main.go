package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Configuration struct {
	Token  string `json:"token"`
	Domain string `json:"domain"`
	Users  string `json:"users"`
}

func main() {
	var config Configuration
	plan, _ := os.Open("configuration.json")
	byteValue, _ := io.ReadAll(plan)
	json.Unmarshal(byteValue, &config)
	usersUrl := config.Domain + config.Users

	resp, _ := http.NewRequest(http.MethodGet, usersUrl, nil)
	resp.Header.Set("Authorization", config.Token)
	outs, err := http.DefaultClient.Do(resp)

	if err != nil {
		fmt.Println(err)
	}
	body, err1 := io.ReadAll(outs.Body)
	if err1 != nil {
		fmt.Println(err1)
	}

	var v []map[string]interface{}
	err = json.Unmarshal(body, &v)

	// for k, val := range v[0] {
	// 	fmt.Println(k, " - ", val)
	// }
}
