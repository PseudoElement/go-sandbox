package streaming

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type StreamSocket struct {
	w    http.ResponseWriter
	r    *http.Request
	conn *websocket.Conn
}

func NewStreamSocket(w http.ResponseWriter, r *http.Request) *StreamSocket {
	return &StreamSocket{w: w, r: r}
}

func (this *StreamSocket) Connect() error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(this.w, this.r, nil)
	if err != nil {
		return err
	}

	this.conn = conn

	return nil
}

func (this *StreamSocket) SendChunks(chunkChan <-chan ReadChunk) {
	for {
		chunk := <-chunkChan
		if chunk.Err != nil {
			log.Println("chunk.Err_Error ===>", chunk.Err)
			this.conn.Close()
			return
		}

		if e := this.conn.WriteMessage(websocket.BinaryMessage, chunk.ChunkBytes); e != nil {
			log.Println("conn.WriteMessage_Error ===>", e)
			return
		}
		if chunk.IsEnd {
			this.conn.Close()
			return
		}
	}
}
