package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("p", "8080", "Listening port")

func main() {
	flag.Parse()
	//Check that all flags are correct
	if flag.NArg() > 0 {
		fmt.Println("Incorrect argument. Please see help below:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	http.HandleFunc("/", root)
	fmt.Printf("Server started and listen %s port\n", *port)
	if err := http.ListenAndServe("localhost:"+*port, nil); err != nil {
		log.Fatalf("Error start server: %v", err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	buf := &bytes.Buffer{}
	for k, v := range r.Header {
		for _, value := range v {
			buf.WriteString(value)
		}
		fmt.Fprintf(w, "%s: %s\n", k, buf.String())
		buf.Reset()
	}
}
