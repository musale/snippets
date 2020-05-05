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
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()
	defer rs.Body.Close()

	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %s but got %s", "deny", frameOptions)
	}
	xssProtection := rs.Header.Get("X-XSS-Protection")
	xssWant := "1;mode=block"
	if xssProtection != xssWant {
		t.Errorf("want %s but got %s", xssWant, xssProtection)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want code %d but got %d", http.StatusOK, rs.StatusCode)
	}

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want %s but got %s", "OK", string(body))
	}
}
