package main

import (
	"bytes"
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

func TestShowSnippet(t *testing.T) {
	app := newTestWebApp(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old test snippet")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			checkStatus(t, tt.wantCode, code)

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}

}
