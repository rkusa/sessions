package sessions

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rkgo/web"
)

const testCookie = "testsid=MTQ0MjczNTc0MnxEdi1CQkFFQ180SUFBUkFCRUFBQUlQLUNBQUVHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5fD8SeqoIye4A9ZQPuT0MIe2SUV-UYT1li2Uj8SlRS9Ka; Path=/; Expires=Tue, 20 Oct 2015 07:55:42 GMT; Max-Age=2592000; HttpOnly"

func TestMiddleware(t *testing.T) {
	app := web.New()
	app.Use(Middleware("testsid", NewCookieStore([]byte("key"))))

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, r)

	if !strings.HasPrefix(rec.Header().Get("Set-Cookie"), "testsid=") {
		t.Error("cookie not set")
	}
}

func TestDecode(t *testing.T) {
	app := web.New()
	app.Use(Middleware("testsid", NewCookieStore([]byte("key"))))
	app.Use(func(ctx web.Context, next web.Next) {
		session := FromContext(ctx)
		foo, ok := session["foo"]

		if !ok || foo != "bar" {
			t.Error("session read error")
		}

		ctx.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	r.Header.Add("Cookie", testCookie)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, r)
}
