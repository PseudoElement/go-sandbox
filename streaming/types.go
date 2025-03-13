package streaming

type ReadChunk struct {
	Err        error
	IsEnd      bool
	ChunkBytes []byte
}
