package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/sianao/gitproxy/cache"

	"github.com/gorilla/mux"
	"github.com/sianao/gitproxy/moudule"
	"github.com/sianao/gitproxy/service"
)

func ServeHTTP(w http.ResponseWriter, req *http.Request, route *mux.Router) {
	route.ServeHTTP(w, req)
}

// https://github.com/laurent22/joplin/releases/download/v3.1.24/Joplin-3.1.24-arm64.DMG
// https://github.com/laurent22/joplin/archive/refs/tags/v3.1.24.zip
// https://raw.githubusercontent.com/laurent22/joplin/dev/.gitignore
// https://github.com/laurent22/joplin/blob/dev/.gitignore
// https://raw.githubusercontent.com/laurent22/joplin/e652db05e1ba47725249a6ff543628aeeb32fad7/.gitignore
// https://raw.githubusercontent.com/laurent22/joplin/android-v3.2.2/.gitignore
// 建立新的router  这里先建立好 方便后续的router
func NewRouter(c *cache.Redis) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/git-upload-pack", func(w http.ResponseWriter, r *http.Request) {
		userBaisc := r.Context().Value(&moudule.B).([]string)
		var address = fmt.Sprintf("https://github.com/%s/%s/git-upload-pack", userBaisc[0], userBaisc[1])
		service.PacketProxy(w, r, address)
	})
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./out/")).ServeHTTP(w, r)

	})

	route.PathPrefix("/_next/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./out/")).ServeHTTP(w, r)

	})
	route.HandleFunc("/info/refs", func(w http.ResponseWriter, r *http.Request) {
		userBaisc := r.Context().Value(&moudule.B).([]string)
		var address = fmt.Sprintf("https://github.com/%s/%s/info/refs?service=git-upload-pack", userBaisc[0], userBaisc[1])
		service.PacketProxy(w, r, address)

	})

	route.HandleFunc("/releases/download/{version}/{file}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userBaisc := r.Context().Value(&moudule.B).([]string)
		var address = fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", userBaisc[0], userBaisc[1], vars["version"], vars["file"])
		path := fmt.Sprintf("github.com/%s/%s/releases/download/%s/%s", userBaisc[0], userBaisc[1], vars["version"], vars["file"])
		c.Incr(path)
		if c.Exists(path) {
			fmt.Println("this way ")
			r.URL.Path = path
			http.FileServer(http.Dir("./cache")).ServeHTTP(w, r)
			v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
			if !ok {
				v = []string{r.RemoteAddr}
			}
			length, _ := strconv.Atoi(w.Header().Get("Content-Length"))
			service.DefaultLogFormatter(
				service.LogFormatterParams{StatusCode: 200,
					ErrorMessage:  "cache",
					ContentLength: humanize.Bytes(uint64(length)), ClientIP: v[0], Method: r.Method, Path: r.URL.Path})
			return
		}
		service.PacketProxy(w, r, address)

	})
	route.HandleFunc("/archive/refs/tags/{tags}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userBaisc := r.Context().Value(&moudule.B).([]string)
		var address = fmt.Sprintf("https://github.com/%s/%s/archive/refs/tags/%s", userBaisc[0], userBaisc[1], vars["tags"])
		path := fmt.Sprintf("github.com/%s/%s/archive/refs/tags/%s",
			userBaisc[0], userBaisc[1], r.URL.Path)
		c.Incr(path)
		if c.Exists(path) {
			r.URL.Path = path
			http.FileServer(http.Dir("./cache")).ServeHTTP(w, r)
			v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
			if !ok {
				v = []string{r.RemoteAddr}
			}
			length, _ := strconv.Atoi(w.Header().Get("Content-Length"))
			service.DefaultLogFormatter(
				service.LogFormatterParams{StatusCode: 200,
					ErrorMessage:  "cache",
					ContentLength: humanize.Bytes(uint64(length)), ClientIP: v[0], Method: r.Method, Path: r.URL.Path})
			return
		}
		service.PacketProxy(w, r, address)

	})
	route.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userBaisc := r.Context().Value(&moudule.B).([]string)
		if len(userBaisc) != 3 {
			http.Error(w, "bad address", 512)
			return
		}
		if userBaisc[2] != "raw" && !strings.HasPrefix(r.RequestURI, "/blob") {
			http.Error(w, "bad address", 512)
			return
		}
		r.RequestURI = strings.TrimPrefix(r.RequestURI, "/blob")
		var address = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s%s",
			userBaisc[0], userBaisc[1], r.URL.Path)
		path := fmt.Sprintf("raw.githubusercontent.com/%s/%s%s",
			userBaisc[0], userBaisc[1], r.URL.Path)

		if c.Exists(path) {
			r.URL.Path = path
			http.FileServer(http.Dir("./cache/")).ServeHTTP(w, r)
			v, ok := r.Header[http.CanonicalHeaderKey("X-Real-IP")]
			if !ok {
				v = []string{r.RemoteAddr}
			}
			length, _ := strconv.Atoi(w.Header().Get("Content-Length"))
			service.DefaultLogFormatter(
				service.LogFormatterParams{StatusCode: 200,
					ErrorMessage:  "cache",
					ContentLength: humanize.Bytes(uint64(length)), ClientIP: v[0], Method: r.Method, Path: r.URL.Path})

			return
		}
		c.Incr(path)
		service.PacketProxy(w, r, address)

	})
	return route

}
