package main

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/musale/snippets/pkg/models/mock"
)

var csrfTokenRX = regexp.MustCompile(`<input type="hidden" name="csrf_token" value='(.+)'>`)

func extractCSRFToken(t *testing.T, body []byte) string {
	// Use the FindSubmatch method to extract the token from the HTML body.
	// Note that this returns an array with the entire matched pattern in the
	// first position, and the values of any captured data in the subsequent
	// positions.
	matches := csrfTokenRX.FindSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}

	return html.UnescapeString(string(matches[1]))
}

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

func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, []byte) {
	rs, err := ts.Client().PostForm(ts.URL+urlPath, form)
	checkError(t, err)

	// Read the response body.
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	checkError(t, err)

	// Return the response status, headers and body.
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
		t.Errorf("want status code %d but got %d", want, got)
	}
}

func checkBody(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want body to be %s but got %s", want, got)
	}
}
