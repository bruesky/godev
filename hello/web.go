package hello

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	msg := []byte("Hello!")
	_, err := writer.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func WebStart() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:9999", nil)
	log.Fatal(err)
}
