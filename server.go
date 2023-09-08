package main

import (
	"log"
	"net/http"
	"strconv"
)

// Server struct
type server struct {
	ipv4 string
	port uint16
	name string
}

// Create a new server struct
func NewServer(options ...serverOptions) *server {
	// Create server pointer struct with default configuration
	s := defaultServerConfiguration()

	// Change the intial configuration
	for _, option := range options {
		option(s)
	}

	// Return struct pointer
	return s
}

// Default configuration
func defaultServerConfiguration() *server {
	s := &server{name: "A", ipv4: "127.0.0.1", port: 8080}
	return s
}

// Function decoration
type serverOptions func(s *server)

// Set server IPv4
func WithIPv4(ipAddress string) serverOptions {
	return func(s *server) {
		s.ipv4 = ipAddress
	}
}

// Set service port
func WithPort(port uint16) serverOptions {
	return func(s *server) {
		s.port = port
	}
}

func withName(name string) serverOptions {
	return func(s *server) {
		s.name = name
	}
}

func (s server) Run(mux *http.ServeMux) {

	socketAddress := s.ipv4 + ":" + strconv.FormatUint(uint64(s.port), 10)
	log.Printf("The server %v is running on:%v\n", s.name, socketAddress)
	if err := http.ListenAndServe(socketAddress, mux); err != nil {
		log.Fatalln(err)
	}
}
