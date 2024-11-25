package service

import (
	"fmt"
	"io"
	"net/http"
)

func PacketProxy(w http.ResponseWriter, r *http.Request, address string) {
	req, _ := http.NewRequest(r.Method, address, r.Body)
	req.Header = r.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("download err  %s  \n", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
	if !ok {
		v = []string{r.RemoteAddr}
	}
	DefaultLogFormatter(LogFormatterParams{StatusCode: resp.StatusCode, ClientIP: v[0], Method: r.Method, Path: r.URL.Path})
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// // StatusCode is HTTP response code.
// StatusCode int
// // Latency is how much time the server cost to process a certain request.
// Latency time.Duration
// // ClientIP equals Context's ClientIP method.
// ClientIP string
// // Method is the HTTP method given to the request.
// Method string
// // Path is a path the client requests.
// Path string
// // ErrorMessage is set if error has occurred in processing the request.
// ErrorMessage string
