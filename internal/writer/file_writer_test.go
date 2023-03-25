package writer

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWriteChunk(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("ok, base file write", func(t *testing.T) {
		fw := NewMockOsFileWriter(ctrl)

		fw.EXPECT().WriteAt(gomock.Any(), gomock.Any()).Return(4, nil).Times(2)
		fd := &FileDescriptor{
			file:        fw,
			chunkStates: make([]bool, 0),
			chunkSize:   4,
		}

		err := fd.WriteChunk(0, []byte("abcd"))
		assert.NoError(t, err)
		err = fd.WriteChunk(1, []byte("efgh"))
		assert.NoError(t, err)
	})

	t.Run("err, bad chunk", func(t *testing.T) {
		fw := NewMockOsFileWriter(ctrl)

		fw.EXPECT().WriteAt(gomock.Any(), gomock.Any()).Return(4, nil).Times(1)
		fd := &FileDescriptor{
			file:        fw,
			chunkStates: make([]bool, 0),
			chunkSize:   4,
		}

		err := fd.WriteChunk(0, []byte("abcd"))
		assert.NoError(t, err)
		err = fd.WriteChunk(1, []byte("efg"))
		assert.Error(t, err)
	})
}
