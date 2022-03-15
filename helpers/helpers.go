package helpers

import (
	"io"
	"log"
	"net/http"
)

func GetData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get response.")
	}
	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error getting the response.")
	}
	return responseByte
}
