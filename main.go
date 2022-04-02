package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://cataas.com/api/cats?tags=cute")
	if err != nil {
		fmt.Println("Error fetching get request.")
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", body)
}
