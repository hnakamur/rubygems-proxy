package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var addr = flag.String("addr", ":8080", "http service address")
var root = flag.String("root", ".", "document root directory")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(handler))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	log.Printf("method:%s\turl:%s", req.Method, req.URL)
	filename := path.Join(*root, req.URL.Path)
	if exists(filename) {
		http.ServeFile(w, req, filename)
		log.Printf("served from local file: %s", filename)
	} else {
		proxy(w, req)
	}
}

func proxy(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}

	// Tweak the request as appropriate:
	//	RequestURI may not be sent to client
	//	URL.Scheme must be lower-case
	req.RequestURI = ""
	req.URL.Scheme = strings.ToLower(req.URL.Scheme)

	// And proxy
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if req.Method != "GET" || hasLocationHeader(resp) {
		log.Printf("forward proxy to remote: %s", req.URL)
		resp.Write(w)
		return
	}

	filename := path.Join(*root, req.URL.Path)
	dir := path.Dir(filename)
	if !exists(dir) {
		os.MkdirAll(dir, 0755)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	fw := bufio.NewWriter(file)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = fw.Flush()
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("saved to local file: %s", filename)

	http.ServeFile(w, req, filename)
}

func hasLocationHeader(resp *http.Response) bool {
	_, ok := resp.Header["Location"]
	return ok
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
