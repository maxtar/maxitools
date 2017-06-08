package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	port        = flag.String("p", "8080", "Listening port")
	postFileDir = flag.String("pfd", "none", "Directory for saving files from POST multi-part requests. If 'none' - files will not be saved.")
	logdir      = flag.String("logdir", "none", "Directory for saving requests history. If 'none' - requests will not be saved.")
	stdout      = flag.Bool("stdout", true, "Enable print requests to standart output.")
	stdlogger   = log.New(os.Stdout, "", log.LstdFlags)
	filelogger  *log.Logger
)

func main() {
	flag.Parse()
	//Check that all flags are correct
	if flag.NArg() > 0 {
		fmt.Println("Incorrect argument. Please see help below:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	var logfile *os.File
	if *logdir != "none" {
		fi, err := os.Stat(*logdir)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if !fi.IsDir() {
			log.Fatalf("%q is not directory.", *logdir)
		}
		logfile, err = os.Create(filepath.Join(*logdir, "stubserver.log")) //todo change hardcorded name to current date
		if err != nil {
			log.Fatalf("Error creating log file: %v", err)
		}
		filelogger = log.New(logfile, "", log.LstdFlags)
		defer logfile.Close()
	}

	http.HandleFunc("/", root)
	fmt.Printf("Server started and listen %s port\n", *port)
	if err := http.ListenAndServe("localhost:"+*port, nil); err != nil {
		log.Fatalf("Error start server: %v", err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	//Print method, url and protocol
	fmt.Fprintf(buf, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	//Print headers
	for k, v := range r.Header {
		fmt.Fprintf(buf, "%s: ", k)
		for _, value := range v {
			buf.WriteString(value)
		}
		fmt.Fprintln(buf)
	}

	//Print body if exists
	body := r.Body
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		buf.WriteString("Error getting body\n")
	}
	if len(content) != 0 {
		//Split headers from body
		fmt.Fprintf(buf, "\nBody:\n")
		buf.Write(content)
	}
	//Send response
	fmt.Fprint(w, buf)
	printOut(buf)
}

func printOut(srcbuf *bytes.Buffer) {
	if filelogger == nil && *stdout == false {
		return
	}
	buf := new(bytes.Buffer)
	buf.WriteString("---------- Start request ----------\n")
	srcbuf.WriteTo(buf)
	buf.WriteString("---------- End request ----------\n")
	if *stdout == true {
		stdlogger.Println(buf)
	}
	if filelogger != nil {
		filelogger.Println(buf)
	}
}
