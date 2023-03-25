package writer

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// TODO: name
type FileCreator struct {
	BasicDir  string
	ChunkSize int
}

func (f *FileCreator) CreateFile(name string) (*FileDescriptor, error) {
	path := filepath.Join(f.BasicDir, name)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return &FileDescriptor{
		file:        file,
		chunkStates: make([]bool, 0),
		chunkSize:   f.ChunkSize,
	}, nil
}

func NewFileCreator(BasicDir string, ChunkSize int) *FileCreator {
	return &FileCreator{
		BasicDir:  BasicDir,
		ChunkSize: ChunkSize,
	}
}

type FileDescriptor struct {
	file        OsFileWriter
	chunkStates []bool
	chunkSize   int
}

var _ ChunkWriter = &FileDescriptor{}

func (f *FileDescriptor) WriteChunk(chunkNumber int, chunk []byte) error {
	if len(chunk) != f.chunkSize {
		return errors.New("incorrect chunk size")
	}

	_, err := f.file.WriteAt(chunk, int64(f.chunkSize)*int64(chunkNumber))
	if err != nil {
		return errors.Wrap(err, "can't write to file")
	}

	if len(f.chunkStates) <= chunkNumber {
		f.chunkStates = append(f.chunkStates, make([]bool, chunkNumber-len(f.chunkStates)+1)...)
	}
	f.chunkStates[chunkNumber] = true

	return nil
}
