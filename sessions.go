// A sessions middleware for [rkgo/web](https://github.com/rkgo/web)
//
//  app := app.New()
//  app.Use(sessions.Middleware("testsid", NewCookieStore([]byte("key"))))
//
// Read session
//
//  sessions := sessions.FromContext(ctx)
//  fmt.Println(sessions["foo"])
//
package sessions

import (
	"fmt"
	"net/http"

	gorilla "github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/rkgo/web"
	"golang.org/x/net/context"
)

type key int

const sessionKey key = 1

type Store sessions.Store

type CookieStore interface {
	Store
}

type cookieStore struct {
	*sessions.CookieStore
}

// NewCookieStore creates a new CookieStore with the given key pairs.
func NewCookieStore(keyPairs ...[]byte) CookieStore {
	store := &cookieStore{sessions.NewCookieStore(keyPairs...)}
	store.Options.HttpOnly = true
	return store
}

// Middleware returns a [rkgo/web](https://github.com/rkgo/web) compatible
// sessions middleware.
func Middleware(name string, store Store) web.Middleware {
	return func(ctx web.Context, next web.Next) {
		session, _ := store.Get(ctx.Req(), name)
		defer gorilla.Clear(ctx.Req())

		ctx.Before(func(rw http.ResponseWriter) {
			session.Save(ctx.Req(), rw)
		})

		next(ctx.WithValue(sessionKey, session))
	}
}

// FromContext reads the session from the given context.
func FromContext(ctx context.Context) map[interface{}]interface{} {
	session, ok := ctx.Value(sessionKey).(*sessions.Session)

	if !ok {
		panic(fmt.Errorf("Sessions Middleware not in use"))
	}

	return session.Values
}
