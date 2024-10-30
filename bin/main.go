package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	userInput := os.Args[1]

	reponse, err := http.Get("https://www.dwds.de/api/wb/snippet/?q=" + userInput)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}

	reponseData, err1 := io.ReadAll(reponse.Body)
	if err1 != nil {
		fmt.Println("error1")
		os.Exit(1)
	}

	fmt.Println(string(reponseData))
}
