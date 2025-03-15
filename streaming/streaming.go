package streaming

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type StreamingModule struct {
	api *mux.Router
}

func NewStreamingModule(api *mux.Router) *StreamingModule {
	return &StreamingModule{api}
}

func (m *StreamingModule) TestReader() {
	ch, err := m.ReadChunks("360p", "sea.mp4", ONE_MB)
	if err != nil {
		panic(err)
	}

	for {
		obj := <-ch
		log.Println("read chunk ==> ", obj.ChunkBytes)

		if obj.Err != nil {
			panic(obj.Err)
		}
		if obj.IsEnd {
			break
		}
	}
}

func (m *StreamingModule) ReadChunks(quality, fileName string, size int) (<-chan ReadChunk, error) {
	pwd, _ := os.Getwd()
	path := pwd + "/streaming/video/" + quality + "/" + fileName

	dataChan := make(chan ReadChunk)
	buf := make([]byte, size, size)

	file, err := os.Open(path)
	if err != nil {
		file.Close()
		return nil, err
	}

	go func() {
		for {
			_, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					dataChan <- ReadChunk{
						Err:        nil,
						IsEnd:      true,
						ChunkBytes: buf,
					}
					file.Close()
					break
				} else {
					dataChan <- ReadChunk{
						Err:        err,
						IsEnd:      true,
						ChunkBytes: nil,
					}
					file.Close()
					break
				}
			}

			dataChan <- ReadChunk{
				Err:        nil,
				IsEnd:      false,
				ChunkBytes: buf,
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	return dataChan, nil
}
