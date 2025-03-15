package streaming

import (
	"fmt"
	"io"
	"net/http"
	"os"

	api_main "github.com/pseudoelement/golang-utils/src/api"
)

func (m *StreamingModule) SetRoutes() {
	m.api.HandleFunc("/app", m._htmlPage).Methods(http.MethodGet)
	m.api.HandleFunc("/js/client.js", m._jsFile).Methods(http.MethodGet)
	m.api.HandleFunc("/streaming/load-video", m._streamWs).Methods(http.MethodGet)
	m.api.HandleFunc("/streaming/video", m._videoFile).Methods(http.MethodGet)
}

func (m *StreamingModule) _jsFile(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	path := pwd + "/streaming/public/client.js"

	http.ServeFile(w, r, path)
}

func (m *StreamingModule) _videoFile(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	path := pwd + "/streaming/video/1080p/sea.mp4"

	file, err := os.Open(path)
	// Handle Range request
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}

	fileSize := int(fileInfo.Size())
	start := 0
	end := fileSize - 1 // Default end of file

	chunkSize := (end - start) + 1
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", chunkSize))
	w.Header().Set("Content-Type", "video/mp4")
	w.WriteHeader(http.StatusPartialContent) // 206 Partial Content

	// Seek to the requested position and stream the file
	file.Seek(int64(start), 0)
	io.CopyN(w, file, int64(chunkSize))
}

func (m *StreamingModule) _htmlPage(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	path := pwd + "/streaming/public/client.html"

	http.ServeFile(w, r, path)
}

func (m *StreamingModule) _streamWs(w http.ResponseWriter, r *http.Request) {
	params, err := api_main.MapQueryParams(r, "fileName", "quality")
	if err != nil {
		api_main.FailResponse(w, err.Error(), 400)
		return
	}

	if err := checkValidQuality(params["quality"]); err != nil {
		api_main.FailResponse(w, err.Error(), 400)
		return
	}
	if err := checkFileNameExists(params["quality"], params["fileName"]); err != nil {
		api_main.FailResponse(w, err.Error(), 400)
		return
	}

	chunkChan, e := m.ReadChunks(params["quality"], params["fileName"], HANDRED_KB)
	if e != nil {
		api_main.FailResponse(w, e.Error(), 400)
		return
	}

	socket := NewStreamSocket(w, r)
	socket.Connect()
	socket.SendChunks(chunkChan)
}
