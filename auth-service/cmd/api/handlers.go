package main

import "net/http"

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate"))
}
