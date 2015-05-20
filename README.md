# sessions

A sessions middleware for [rkgo/web](https://github.com/rkgo/web)

[![Build Status][drone]](https://ci.rkusa.st/github.com/rkgo/sessions)
[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/sessions)

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

[drone]: http://ci.rkusa.st/api/badge/github.com/rkgo/sessions/status.svg?branch=master&style=flat-square
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square