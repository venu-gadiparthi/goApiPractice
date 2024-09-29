package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Data struct {
	Token string `json:"token"`
}

func main() {
	plan, _ := os.Open("configuration.json")
	byteValue, _ := io.ReadAll(plan)
	var dats Data
	json.Unmarshal(byteValue, &dats)
	fmt.Println(dats.Token)
	usersUrl := "https://trial-3926351.okta.com/api/v1/users"
	resp, _ := http.NewRequest(http.MethodGet, usersUrl, nil)
	resp.Header.Set("Authorization", dats.Token)
	outs, err := http.DefaultClient.Do(resp)
	if err != nil {
		fmt.Println(err, "er one")
	}
	body, err1 := io.ReadAll(outs.Body)
	if err1 != nil {
		fmt.Println(err1, "er two")
	}
	var v []map[string]interface{}
	err3 := json.Unmarshal(body, &v)
	fmt.Println(err3)
	for k, val := range v[0] {
		fmt.Println(k, " - ", val)
	}
}
