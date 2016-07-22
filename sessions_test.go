package sessions

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rkusa/web"
)

const sessionName = "testsid"

var keyPair = []byte("key")
var testCookie string

func TestEncode(t *testing.T) {
	app := web.New()
	app.Use(Middleware(sessionName, NewCookieStore(keyPair)))
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session := FromContext(r.Context())
		session["foo"] = "bar"

		rw.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, r)

	testCookie = rec.Header().Get("Set-Cookie")
	if !strings.HasPrefix(testCookie, sessionName+"=") {
		t.Fatal("Cookie not set")
	}
}

func TestDecode(t *testing.T) {
	app := web.New()
	app.Use(Middleware("testsid", NewCookieStore(keyPair)))
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session := FromContext(r.Context())
		foo, ok := session["foo"]

		if !ok || foo != "bar" {
			t.Errorf(`expected foo="bar", got "%v"`, foo)
		}

		rw.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	r.Header.Add("Cookie", testCookie)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, r)
}
