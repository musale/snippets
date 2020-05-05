package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/musale/snippets/pkg/models/mock"
)

func newTestWebApp(t *testing.T) *webApp {
	t.Helper()
	templateCache, err := newTemplateCache("./../../ui/html/")
	checkError(t, err)

	session := sessions.New([]byte("199d6fe6d83a43bb1599e42c9b4b2308"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	return &webApp{
		errorLog:      log.New(ioutil.Discard, "", 0),
		infoLog:       log.New(ioutil.Discard, "", 0),
		templateCache: templateCache,
		session:       session,
		snippets:      &mock.SnippetModel{},
		users:         &mock.UserModel{},
	}
}

type testServer struct {
	*httptest.Server
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	t.Helper()
	rs, err := ts.Client().Get(ts.URL + urlPath)
	checkError(t, err)
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	checkError(t, err)

	return rs.StatusCode, rs.Header, body
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	t.Helper()
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	checkError(t, err)

	ts.Client().Jar = jar
	ts.Client().CheckRedirect = func(
		req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &testServer{ts}
}

func assertEqual(t *testing.T, want, got interface{}) {
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}

func checkError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func checkStatus(t *testing.T, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
}

func checkBody(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
