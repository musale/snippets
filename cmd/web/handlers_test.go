package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("Unit test Ping", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest("Get", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		ping(rr, r)

		rs := rr.Result()
		checkStatus(t, http.StatusOK, rs.StatusCode)

		defer rs.Body.Close()

		body, err := ioutil.ReadAll(rs.Body)
		checkError(t, err)

		checkBody(t, "OK", string(body))
	})

	t.Run("Test Ping in the server", func(t *testing.T) {
		app := newTestWebApp(t)

		ts := newTestServer(t, app.routes())
		defer ts.Close()
		statusCode, _, body := ts.get(t, "/ping")

		checkStatus(t, http.StatusOK, statusCode)
		checkBody(t, "OK", string(body))
	})
}
