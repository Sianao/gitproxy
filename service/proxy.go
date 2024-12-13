package service

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/dustin/go-humanize"
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
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	DefaultLogFormatter(
		LogFormatterParams{StatusCode: resp.StatusCode,
			ContentLength: humanize.Bytes(uint64(length)), ClientIP: v[0], Method: r.Method, Path: r.URL.Path})

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
