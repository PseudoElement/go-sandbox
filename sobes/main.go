package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func compareSlices[T any](s1, s2 []T) bool {
	buf1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	buf2, err := json.Marshal(s2)
	if err != nil {
		panic(err)
	}
	if len(buf1) != len(buf2) {
		return false
	}
	for i := range len(buf1) {
		equal := buf1[i] == buf2[i]
		if !equal {
			return false
		}
	}
	return true
}

func main() {
	s1 := []EmptyResp{{Data: []int{223, 22, 35}}}
	s2 := []EmptyResp{{Data: []int{223, 22, 35}}}
	log.Println(compareSlices(s1, s2))
	// checkEmptySlice()
	// fanOut()
	// customMutex()
}

type EmptyResp struct {
	Data []int `json:"data"`
}

func checkEmptySlice() {
	buf := []byte("{\"data\": null}")
	var emptyResp EmptyResp
	if err := json.Unmarshal(buf, &emptyResp); err != nil {
		panic(err)
	}

	fmt.Printf("resp: %+v\n", emptyResp)
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
