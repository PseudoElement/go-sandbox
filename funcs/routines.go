package funcs

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"sync/atomic"
	"time"
)

type User struct {
	Name string
	Id   string
}

func generateRandomString() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func fetchUser(idx int) User {
	time.Sleep(1 * time.Second)
	name := "User_" + strconv.Itoa(idx)
	return User{name, generateRandomString()}
}

func processUser(user User) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(user.Name, " is processing.")
}

func saveResult(user User) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(user.Name, " successfully saved.")
}

func handleRequests() <-chan User {
	outChan := make(chan User)
	var counter atomic.Uint64
	var defaultCounter int64
	for i := 0; i < 100000; i++ {
		go func() {
			user := fetchUser(i + 1)
			processUser(user)
			saveResult(user)
			outChan <- user
			counter.Add(1)
			defaultCounter++
			if counter.Load() == 100000 {
				log.Printf("Default counter - %v. Atomic - %v.\n", defaultCounter, counter.Load())
				close(outChan)
			}
		}()
	}

	return outChan
}

func processChan(resChan <-chan User) ([]User, error) {
	users := make([]User, 0, len(resChan))
	for {
		select {
		case <-time.After(5 * time.Second):
			return users, fmt.Errorf("Timeout in 5 seconds expired!\n")
		case user, ok := <-resChan:
			if !ok {
				return users, nil
			}
			users = append(users, user)
		}
	}
}

func MakeRequests() {
	start := time.Now()
	outChan := handleRequests()
	_, err := processChan(outChan)
	if err != nil {
		panic(err)
	}
	fmt.Println("Execution time: ", time.Since(start))
}
