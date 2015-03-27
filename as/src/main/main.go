package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"
)

var (
	host string
	port int
)

func init() {
	// Flags
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.IntVar(&port, "p", 8001, "web port")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", dummyHandler)
	http.HandleFunc("/health", healthHandler)
	fmt.Printf("Server started on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func dummyHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s - - [%s] \"%s %s %s\" \"%s\" \"%s\"\n", req.RemoteAddr,
		time.Now().Format("02/Jan/2006:15:04:05 -0700"), req.Method, req.RequestURI, req.Proto, req.Referer(), req.UserAgent())

	host, _ := os.Hostname()
	fmt.Fprintf(rw, "Hello, %q. I am server hosted on: %s", html.EscapeString(req.URL.Path), host)
}

func healthHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s - - [%s] \"%s %s %s\" \"%s\" \"%s\"\n", req.RemoteAddr,
		time.Now().Format("02/Jan/2006:15:04:05 -0700"), req.Method, req.RequestURI, req.Proto, req.Referer(), req.UserAgent())

	rw.Write([]byte("ok"))
}
