package handler

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"

	"github.com/gorilla/mux"
	"github.com/sianao/gitproxy/moudule"
	"github.com/sianao/gitproxy/router"
	"github.com/sianao/gitproxy/service"
)

const (
	Github = "git"
	RawGit = "raw"
)

func urlProcess(w http.ResponseWriter, r *http.Request) string {
	switch {
	case strings.HasPrefix(r.RequestURI, "/https:/raw.githubusercontent.com/"):
		///https:/raw.githubusercontent.com/
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/https:/raw.githubusercontent.com/"))
		r.RequestURI = r.URL.String()
		return RawGit
	case strings.HasPrefix(r.RequestURI, "/https://raw.githubusercontent.com/"):
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/https://raw.githubusercontent.com/"))
		r.RequestURI = r.URL.String()
		return RawGit
	case strings.HasPrefix(r.RequestURI, "/https:/github.com/"):
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/https:/github.com/"))
		r.RequestURI = r.URL.String()
		return Github
	case strings.HasPrefix(r.RequestURI, "/https://github.com/"):
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/https://github.com/"))
		r.RequestURI = r.URL.String()
		return Github
	case strings.HasPrefix(r.RequestURI, "/github.com/"):
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/github.com/"))
		r.RequestURI = r.URL.String()
		return Github

	case strings.HasPrefix(r.RequestURI, "/raw.githubusercontent.com/"):
		r.URL, _ = url.Parse(strings.TrimPrefix(r.RequestURI, "/raw.githubusercontent.com/"))
		r.RequestURI = r.URL.String()
		return RawGit

	default:
		// 不支持的文件形式
		http.NotFound(w, r)
		return ""
	}
}
func NewHandler(route *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/" || strings.HasPrefix(r.RequestURI, "/_next/") {
			router.ServeHTTP(w, r, route)
			return
		}
		var types = urlProcess(w, r)
		if types == "" {
			v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
			if !ok {
				v = []string{r.RemoteAddr}
			}
			length, _ := strconv.Atoi(w.Header().Get("Content-Length"))

			service.DefaultLogFormatter(
				service.LogFormatterParams{StatusCode: 404,
					ContentLength: humanize.Bytes(uint64(length)), ClientIP: v[0], Method: r.Method, Path: r.URL.Path})
			return
		}
		//去除掉host方便进入路由匹配
		sub := strings.Split(strings.TrimPrefix(r.URL.String(), "/"), "/")
		if len(sub) <= 2 {
			v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
			if !ok {
				v = []string{r.RemoteAddr}
			}
			service.DefaultLogFormatter(
				service.LogFormatterParams{StatusCode: 404, ClientIP: v[0], Method: r.Method, Path: r.URL.Path})
			http.NotFound(w, r)
			return
		}
		//去除掉用户以及仓库的具体信息 重新建立 url 方便路由
		r = r.WithContext(
			context.WithValue(r.Context(), &moudule.B, append(append([]string{}, sub[:2]...), types)),
		)
		r.URL, _ = url.Parse("/" + strings.Join(sub[2:], "/"))
		r.RequestURI = r.URL.String()
		router.ServeHTTP(w, r, route)
	}
}
