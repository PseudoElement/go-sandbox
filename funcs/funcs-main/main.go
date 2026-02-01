package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func getRequest(urlStr string, params [][2]string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	p := url.Values{}
	for _, pair := range params {
		p.Add(pair[0], pair[1])
	}

	if len(params) > 0 {
		urlStr += "?" + p.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func main() {
	b1 := true
	b2 := false

	s1 := strconv.FormatBool(b1)
	s2 := strconv.FormatBool(b2)

	fmt.Printf("%v becomes %q\n", b1, s1)
	fmt.Printf("%v becomes %q\n", b2, s2)
}
