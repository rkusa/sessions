// Package sessions provides a sessions middleware that works well (but not
// exclusively) with [rkusa/web](https://github.com/rkusa/web).
//
//  app := app.New()
//  app.Use(sessions.Middleware("testsid", sessions.NewCookieStore([]byte("your-secret-key"))))
//
//  add this
//
//
// Read session
//
//  sessions := sessions.FromContext(r.Context())
//  fmt.Println(sessions["foo"])
//
package sessions

import (
	"context"
	"fmt"
	"net/http"

	gorilla "github.com/gorilla/context"
	"github.com/gorilla/sessions"
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

// Middleware returns a middleware.
func Middleware(name string, store Store) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session, _ := store.Get(r, name)
		defer gorilla.Clear(r)

		next(rw, r.WithContext(context.WithValue(r.Context(), sessionKey, session)))

		session.Save(r, rw)
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
