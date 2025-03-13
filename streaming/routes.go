package streaming

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	api_main "github.com/pseudoelement/golang-utils/src/api"
)

func (m *StreamingModule) SetRoutes() {
	m.api.HandleFunc("/app", m._htmlPage).Methods(http.MethodGet)
	m.api.HandleFunc("/streaming/load-video", m._streamWs).Methods(http.MethodGet)
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

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		api_main.FailResponse(w, e.Error(), 400)
		return
	}

	for {
		chunk := <-chunkChan
		if chunk.Err != nil {
			log.Println("chunk.Err_Error ===>", e)
			conn.Close()
			return
		}
		if e := conn.WriteMessage(websocket.BinaryMessage, chunk.ChunkBytes); e != nil {
			log.Println("conn.WriteMessage_Error ===>", e)
			return
		}
		if chunk.IsEnd {
			conn.Close()
			return
		}
	}
}
