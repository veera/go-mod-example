package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	port := kingpin.Flag("port", "Port Number to listen").Default("8000").Short('p').String()
	kingpin.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:"+*port, r)
	fmt.Println("Server started 0.0.0.0:%v", port)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "GoLang")
	logrus.Info(request.RemoteAddr)
}
