package server

import "net/http"

//LOGGER MIDDLEWARE START

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body = append(lrw.body, b...)
	return lrw.ResponseWriter.Write(b)
}

func (s *Server) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(lrw, r)

		s.logger.Printf("Method: %s, URL: %s, Status: %d, ResponseBody: %s, IP:%s \n",
			r.Method, r.URL.String(), lrw.statusCode, string(lrw.body), r.RemoteAddr)
	})
}

//LOGGER MIDDLEWARE END

//TODO add a function that creates session_id as a cookie

//TODO add a funtcion that checks session_id from a cookie

//TODO add a function that create CSRF_TOKEN for a login page

//TODO add a function that checks CSRF_TOKEN from a login page
