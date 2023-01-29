package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

type proxy struct {
	client *http.Client
}

var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = ""
	delHopHeaders(r.Header)

	resp, err := p.client.Do(r)
	if err != nil {
		var urlErr *url.Error
		errors.As(err, &urlErr)
		if urlErr.Timeout() {
			http.Error(w, "request timed out", http.StatusRequestTimeout)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	defer resp.Body.Close()

	env.infoLog.Println(r.RemoteAddr, r.Method, r.URL, resp.Status)
	delHopHeaders(resp.Header)
	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func copyHeaders(dst, src http.Header) {
	for hkey, hvals := range src {
		for _, hval := range hvals {
			dst.Add(hkey, hval)
		}
	}
}
