package main

import (
	"io"
	"net/http"
)

func NewInternalServer() {
	internalServ := NewServer(withName("Internal service"), WithIPv4("127.0.0.1"), WithPort(8080))
	mux := http.NewServeMux()
	mux.HandleFunc("/ctf", ctf)
	internalServ.Run(mux)
}

func ctf(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "CTF{R0undR0b!nC4u5eTr0uble}!\n")
}
