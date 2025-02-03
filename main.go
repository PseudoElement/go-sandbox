package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pseudoelement/go-sandbox/common/constants"
	"github.com/pseudoelement/go-sandbox/funcs"
)

// func runProfiling() {
// 	cpuProfileFile, err := os.Create("./prof/cpu.prof")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cpuProfileFile.Close()

// 	// Start CPU profiling
// 	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
// 		panic(err)
// 	}
// 	defer pprof.StopCPUProfile()

// 	fmt.Println("CPU profile written to /prof/cpu.prof")

// 	memProfileFile, err := os.Create("./prof/mem.prof")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer memProfileFile.Close()

// 	// Write memory profile to file
// 	if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Memory profile written to /prof/mem.prof")

// 	// Start tracing
// 	traceFile, err := os.Create("trace.out")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer traceFile.Close()

// 	if err := trace.Start(traceFile); err != nil {
// 		panic(err)
// 	}
// 	defer trace.Stop()

// 	fmt.Println("Tracing written to /prof/trace.out")
// }

// func main() {
// 	runProfiling()

// 	now := time.Now()
// 	funcs.IsHappy(1999959944342334234)
// 	fmt.Printf("IsHappy took %v ms!\n", time.Since(now))

// 	heavyCalc()

// 	fmt.Println("Profiling done!")
// }

// func heavyCalc() {
// 	for i := 0; i < 1_000; i++ {
// 		time.Sleep(10 * time.Millisecond)
// 		randNum := rand.Intn(100)
// 		square := math.Pow(float64(randNum), 2)
// 		if i%100 == 0 {
// 			fmt.Printf("%v ** 2 = %v\n", randNum, square)
// 		}
// 	}
// }

func main() {
	funcs.IsColRowKey("00_01_02_10_11_12_20_21_22", "1220")

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		msg := struct {
			Message string `json:"message"`
		}{Message: "Server is alive!"}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)
	}).Methods("GET")

	api.HandleFunc("/tonconnect/manifest/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		path, ok := constants.URL_TO_MANIFEST_PATH_MAP[id]
		if !ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Invalid environvent id!"))
			return
		}
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, path)
	})

	api.HandleFunc("/tonconnect/logo.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, "./static/rubic-logo.png")
	})

	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Listening port :8080")
	log.Fatal(http.ListenAndServe("8080", handlers.CORS(methods, ttl, origins)(api)))
}
