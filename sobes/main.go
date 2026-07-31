package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
)

func main() {
	// fanOut()
	// customMutex()
}

func doRequest(ctx context.Context, addr string, request []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", addr, bytes.NewBuffer(request))
	if err != nil {
		return []byte(""), err
	}

	resp, err := client.Do(req)
	buf, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return []byte(""), err
	}

	return buf, nil
}

type ApiResponse struct {
	value string
}

func batchReqs(addrs []string, req []byte) ApiResponse {
	wg := &sync.WaitGroup{}
	var resp ApiResponse
	wg.Add(len(addrs))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, addr := range addrs {
		go func() {
			defer wg.Done()
			respBuf, err := doRequest(ctx, addr, req)
			if err == nil {
				resp = ApiResponse{value: string(respBuf)}
				cancel()
			}
		}()
	}

	wg.Wait()

	return resp
}
