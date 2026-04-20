package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func logRequest(r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}
	body, _ := io.ReadAll(r.Body)
	logger.Printf("%s %s %s %v %s", ip, r.Method, r.URL.Path, r.Header, string(body))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
}

func main() {
	hostname := flag.String("hostname", "127.0.0.1", "hostname this programm will bind to")
	port := flag.String("port", "8080", "port that will be used")
	logfile := flag.String("log-file", "", "logfile that will optionally be used")
	flag.Parse()

	if *logfile != "" {
		f, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		logger.SetOutput(io.MultiWriter(os.Stdout, f))
	}

	addr := *hostname + ":" + *port
	http.HandleFunc("/", getRoot)
	println("Server listening at ", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
