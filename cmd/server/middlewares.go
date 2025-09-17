package server

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"sync"
	"time"
)

//LOGGER MIDDLEWARE START

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	return lrw.ResponseWriter.Write(b)
}

func (s *Server) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(lrw, r)

		s.logger.Printf("Method: %s, URL: %s, Status: %d, IP:%s \n",
			r.Method, r.URL.String(), lrw.statusCode, r.RemoteAddr)
	})
}

//LOGGER MIDDLEWARE END

//SESSIONID MIDDLEWARE START

type Session struct {
	user_id      string
	expiringTime time.Time
}

var sessionPack sync.Map

func (s *Server) getSessionId(r *http.Request) (string, error) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return "", err
	}

	_, ok := sessionPack.Load(cookie.Value)

	if !ok {
		return "", errors.New("session id is invalid")
	}
	return cookie.Value, nil
}

func (s *Server) periodicSessionClear() {
	for {
		sessionPack.Range(func(key, value any) bool {
			s := value.(Session)
			if time.Now().After(s.expiringTime) {
				sessionPack.Delete(key)
			}
			return true
		})
		time.Sleep(10 * time.Minute)
	}
}

func sessionIdCreator() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *Server) sessionIdFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			if err == http.ErrNoCookie {
				sid := sessionIdCreator()
				sessionPack.Store(sid, Session{
					user_id:      "",
					expiringTime: time.Now().Add(60 * time.Minute),
				})
				http.SetCookie(w, &http.Cookie{
					Name:     "sessionID",
					Value:    sid,
					Path:     "/",
					HttpOnly: true,
					Secure:   false,
					MaxAge:   3600,
				})
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "sessionID",
				Value:    cookie.Value,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				MaxAge:   3600,
			})
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) sessionIdApi(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, ok := sessionPack.Load(cookie.Value)
		if !ok {
			http.Error(w, "there is no session id like that", http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	})
}

//SESSIONID MIDDLEWARE END

//TODO add a function that creates session_id as a cookie

//TODO add a funtcion that checks session_id from a cookie

//TODO add a function that create CSRF_TOKEN for a login page

//TODO add a function that checks CSRF_TOKEN from a login page
