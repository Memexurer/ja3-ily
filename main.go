package main

import (
	"fmt"
	"io"
	"net/url"

	http "github.com/Danny-Dasilva/fhttp"
)

func hello(w http.ResponseWriter, req *http.Request) {
	var err error

	originalRawQuery := req.URL.RawQuery
	if req.URL, err = url.ParseRequestURI(req.Header.Get("TargetUrl")); err != nil {
		panic(err)
	}
	req.URL.RawQuery = originalRawQuery // url param proxy

	httpclient, err := NewClient(Browser{JA3: req.Header.Get("JA3"), UserAgent: req.Header.Get("UA"), Cookies: nil},
		60, false, req.Header.Get("UA"), req.Header.Get("Proxy")) // idk??? i stole this from dmdgo
	if err != nil {
		w.WriteHeader(502)
		fmt.Fprintf(w, "%v", err)
		return
	}

	req.Header.Del("TargetUrl")
	req.Header.Del("JA3")
	req.Header.Del("UA")
	req.Header.Del("Proxy")
	// remove headers so they wont get passed to the upstream

	resp, err := httpclient.Do(&http.Request{
		Method:     req.Method,
		URL:        req.URL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     req.Header,
		Body:       req.Body,
		Host:       req.URL.Host,
	})
	if err != nil {
		w.WriteHeader(502)
		fmt.Fprintf(w, "%v", err)
		return
	}

	for name, headers := range resp.Header {
		for _, value := range headers {
			w.Header().Add(name, value)
		}
	}

	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Running on 127.0.0.1:8090")
	err := http.ListenAndServe("127.0.0.1:8090", nil)
	if err != nil {
		panic(err)
	}
}
