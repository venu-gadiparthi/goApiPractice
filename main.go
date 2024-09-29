package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Data struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Token  string `json:"token"`
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
	bodyString := string(body)
	fmt.Println(bodyString)

	// Unmarshal the byte array into the 'person' variable.

}
