# sessions

A sessions middleware that works well (but not exclusively) with [rkusa/web](https://github.com/rkusa/web).

[![Build Status][travis]](https://travis-ci.org/rkusa/sessions)
[![GoDoc][godoc]](https://godoc.org/github.com/rkusa/sessions)

### Example

Use Middleware

```go
app := app.New()
app.Use(sessions.Middleware("testsid", NewCookieStore([]byte("key"))))
```

Read session

```go
sessions := sessions.FromContext(ctx)
fmt.Println(sessions["foo"])
```

[travis]: https://img.shields.io/travis/rkusa/sessions.svg
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg
