package writer

//go:generate mockgen -source model.go -destination model_mock.go -package writer

type OsFileWriter interface {
	WriteAt([]byte, int64) (int, error)
}

type ChunkWriter interface {
	WriteChunk(int64, []byte) error
}
