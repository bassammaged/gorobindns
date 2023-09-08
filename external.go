package main

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	u "net/url"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func NewExternalServer() {
	internalServ := NewServer(withName("public service"), WithIPv4("192.168.1.235"), WithPort(8080))
	mux := http.NewServeMux()
	mux.HandleFunc("/check-response", requesterHandler)
	internalServ.Run(mux)
}

func requesterHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	url := params.Get("url")

	// Check the existence of url parameter
	if len(url) == 0 {
		io.WriteString(w, "url Parameter is missing"+"\n")
		return
	}

	// Validate the URL format
	host, err := urlValidator(url)
	if err != nil {
		io.WriteString(w, err.Error()+"\n")
		return
	}

	// Check SSRF attacks
	if err := againstSSRF(host); err != nil {
		io.WriteString(w, err.Error()+"\n")
		return
	}

	// Make a HTTP request
	responseBody, err := requester(url)
	if err != nil {
		io.WriteString(w, err.Error()+"\n")
		return
	}

	io.WriteString(w, responseBody+"\n")
}

func urlValidator(url string) (string, error) {
	URL, err := u.ParseRequestURI(url)
	if err != nil {
		return "", err
	}
	return URL.Host, nil
}

func againstSSRF(host string) error {
	blacklist := []string{
		"127.0.0.1",
	}

	// Resolve the host
	result := strings.Split(host, ":")

	lookupResult, err := net.LookupIP(result[0])
	if err != nil {
		return err
	}
	for _, ip := range lookupResult {
		if ip.IsLoopback() || ip.IsPrivate() {
			return errors.New("private IPs are not allowed")
		}
		// Debugging purpose
		log.Println(ip)

		if slices.Contains(blacklist, ip.String()) {
			return errors.New("private IPs are not allowed [hardcoded list]")
		}

	}

	return nil
}

func requester(url string) (string, error) {
	time.Sleep(time.Second * 5)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	byteReader, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(byteReader), nil
}
