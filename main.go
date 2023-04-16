package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type multiValueFlag []string

func (m *multiValueFlag) String() string {
	return fmt.Sprintf("%v", *m)
}

func (m *multiValueFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	base, _ := url.Parse(baseurl)
	path, _ := url.Parse(r.URL.Path)
	resolvedURL := base.ResolveReference(path)
	req, _ := http.NewRequest(r.Method, resolvedURL.String(), r.Body)

	req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
	for _, header := range headers {
		kv := strings.Split(header, ":")
		req.Header.Set(kv[0], kv[1])
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Remote Error:%s", err), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	// レスポンスボディを読み取る
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	_, err = fmt.Fprintln(w, string(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
}

var headers multiValueFlag
var baseurl string

func main() {
	flag.Var(&headers, "H", "headers (can be used multiple times)")
	port := flag.Int("port", 8080, "listen port")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("url is required")
	} else {
		baseurl = args[0]
	}

	http.HandleFunc("/", handle)
	log.Printf("Listening on port %s", strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(*port)), nil))
}
