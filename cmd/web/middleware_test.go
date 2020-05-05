package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	checkError(t, err)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()
	defer rs.Body.Close()

	frameOptions := rs.Header.Get("X-Frame-Options")
	assertEqual(t, "deny", frameOptions)

	xssProtection := rs.Header.Get("X-XSS-Protection")
	xssWant := "1;mode=block"
	assertEqual(t, xssWant, xssProtection)

	checkStatus(t, http.StatusOK, rs.StatusCode)
	body, err := ioutil.ReadAll(rs.Body)
	checkError(t, err)
	checkBody(t, "OK", string(body))
}
