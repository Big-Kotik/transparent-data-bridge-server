package bridgeservice

import (
	"io"

	v1 "github.com/Big-Kotik/transparent-data-bridge-api/bridge/api/v1"
	"github.com/Big-Kotik/transparent-data-bridge-server/internal/writer"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type bridgeService struct {
	v1.UnimplementedTransparentDataBridgeServiceServer
	writer *writer.FileCreator
}

func (b *bridgeService) SendChunks(stream v1.TransparentDataBridgeService_SendChunksServer) error {
	log.Trace().Msg("recive SendChunks request")

	metaInfo, err := stream.Recv()
	if err != nil {
		log.Debug().Err(err).Msg("can't extract first value from stream")
		return errors.Wrap(err, "can't extract first value from stream")
	}

	req := metaInfo.GetRequest()
	if req == nil {
		log.Debug().Err(err).Msg("first message in the stream must be of request type")
		return errors.New("first message in the stream must be of request type")
	}

	if err := validateRequest(req); err != nil {
		log.Debug().Err(err).Msg("validation error")
		return err
	}

	d, err := b.writer.CreateFile(req.FileName)
	if err != nil {
		log.Debug().Err(err).Msg("can't create file")
		return errors.Wrap(err, "can't create file")
	}

	chunks := 0
	for {
		reqWithChunk, err := stream.Recv()
		if err == io.EOF {
			log.Trace().Msg("end of the stream")
			return stream.SendAndClose(&v1.FileStatus{LastChunkOffset: int32(chunks)})
		}
		if err != nil {
			log.Debug().Err(err).Msg("stream recv return err")
			return err
		}

		reqWithChunk.GetChunk()
		c := reqWithChunk.GetChunk()
		if c == nil {
			log.Debug().Any("recived value", reqWithChunk).Msg("wrong type of request")
			return errors.New("expected chunk")
		}

		err = d.WriteChunk(chunks, c.Chunk)
		if err != nil {
			log.Debug().Err(err).Msg("can't write chunk to file")
			return errors.Wrap(err, "can't write chunk to file")
		}
	}
}

func NewBridgeService(writer *writer.FileCreator) v1.TransparentDataBridgeServiceServer {
	return &bridgeService{
		writer: writer,
	}
}
