package main

import (
	"net/http"

	proxy "github.com/pulpfree/api-proxy/lib"
)

func main() {
	p := &proxy.Proxy{}
	p.New()

	server1 := http.NewServeMux()
	server1.HandleFunc("/", p.Handle)

	server2 := http.NewServeMux()
	server2.HandleFunc("/", p.Handle)

	server3 := http.NewServeMux()
	server3.HandleFunc("/", p.Handle)

	go func() {
		http.ListenAndServe(":8001", server1)
	}()

	go func() {
		http.ListenAndServe(":8002", server2)
	}()

	http.ListenAndServe(":8003", server3)
}
