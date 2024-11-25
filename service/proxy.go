package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PacketProxy(w http.ResponseWriter, r *http.Request, address string, log *logrus.Logger) {
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
	log.Println(address, resp.StatusCode)
	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, fmt.Sprintf("copying response body err: %s", err), http.StatusInternalServerError)
	}
}
