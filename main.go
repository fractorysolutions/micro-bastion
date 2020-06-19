package main

import "fmt"
import "net/http"
import "net/url"
import "flag"
import "log"
import "io"
import "strings"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "Bastion is up")
		return
	}
	r.URL = calculateURL(r)
	r.Host = r.URL.Host
	log.Println("Requesting ", r.URL)
	resp, err := http.DefaultTransport.RoundTrip(r)

	if err != nil {
		log.Println("Could not fetch ", r.URL, ": ", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	copyHeader(resp.Header, w)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(from http.Header, to http.ResponseWriter) {
	toHeader := to.Header()
	for k, vs := range from {
		for _, v := range vs {
			toHeader.Add(k, v)
		}
	}
}

func calculateURL(r *http.Request) *url.URL {
	newURL := *r.URL
	oldPath := r.URL.Path
	oldPathParts := strings.Split(oldPath, "/")[1:]
	newURL.Host = oldPathParts[0] + ":" + oldPathParts[1]
	newURL.Path = "/" + strings.Join(oldPathParts[2:], "/")

	newURL.Scheme = "http"

	return &newURL
}

func main() {
	var port = flag.Int("port", 8888, "port that micro-bastion should listen on")
	flag.Parse()

	log.Println("Starting micro-bastion on port", *port)

	server := &http.Server{
		Addr:              fmt.Sprint(":", *port),
		Handler:           http.HandlerFunc(handleRequest),
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}

	// start the server
	log.Fatal(server.ListenAndServe())
}
