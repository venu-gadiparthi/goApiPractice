package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	// Unmarshal the byte array into the 'person' variable.
	json.Unmarshal(byteValue, &dats)
	fmt.Println(dats.Token)

}
