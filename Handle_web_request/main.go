package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	databyte, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	} else {
		content := string(databyte)
		fmt.Println(content)
	}
	response.Body.Close()
}
