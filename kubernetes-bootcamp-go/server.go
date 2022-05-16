package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// var host string
// var startTime time.Time
// var version string
// var requests = 0

type server struct {
	auth struct {
		username, password string
	}
	info struct {
		host, version string
		requests      int
		startTime     time.Time
	}
}

// excellent article on how to implement this in go: https://www.alexedwards.net/blog/basic-authentication-in-go
func (srv *server) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// extra uysername & password from Auth request header.
		// if values are missing or header is invalid then `ok` ret val
		// will be false
		username, password, ok := req.BasicAuth()
		if ok {
			// calculate some shas to mitigate against timing attacks
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			// TOOD: Replace this with env vars or some other type of
			// config injection
			expectedUsernameHash := sha256.Sum256([]byte(srv.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(srv.auth.password))

			// use subtle.ConstantTimeCompare to further mitigate
			// against timing attacks. Furthermore, check both password
			// _AND_ username hashes before returning to avoid
			// accidentally leaking information.
			// TODO: if you are comparing hashes is `ConstantTimeCompare`
			// really necessary?
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			// if the username & password are correct, call the next
			// handler in the chain. Make sure to return so none of the
			// below code is run
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, req)
				return
			}
		}

		// if the auth didn't pass for some reason. Set a
		// [`WWW-Authenticate`](https://datatracker.ietf.org/doc/html/rfc7235#section-4.1)
		// header to inform the client basic auth is expected and return
		//  a 401
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (srv *server) healthHandler(w http.ResponseWriter, req *http.Request) {
	ok, _ := json.Marshal("OK")
	fmt.Printf("Health endpoint hit. | Log Time: %v\n", time.Now().Format(time.UnixDate))
	fmt.Fprintf(w, "%v\n", string(ok))
}

func (srv *server) handler(w http.ResponseWriter, req *http.Request) {

	now := time.Now()
	srv.info.requests++
	upTime := now.Sub(srv.info.startTime)
	fmt.Printf("Running On: %v | Requests: %v | Uptime: %v | Log Time: %v\n", srv.info.host, srv.info.requests, upTime.Truncate(time.Millisecond), now.Format(time.UnixDate))

	fmt.Fprintf(w, "Hello Kubernetes bootcamp go! | Running on: %v | v=%v\n", srv.info.host, srv.info.version)
}
