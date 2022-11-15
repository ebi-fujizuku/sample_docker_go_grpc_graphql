package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	// log.Println("start http server :",os.Getenv("BACK_PORT"))
	// port := ":" + os.Getenv("BACK_PORT")
	// log.Fatal(http.ListenAndServe(port, nil))
	log.Println("start http server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
