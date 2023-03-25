package bridgeservice

import (
	"context"
	"io"

	v1 "github.com/Big-Kotik/transparent-data-bridge-api/bridge/api/v1"
	"github.com/Big-Kotik/transparent-data-bridge-server/internal/writer"
	"github.com/rs/zerolog/log"
)

type bridgeService struct {
	proxyClient v1.TransparentDataRelayServiceClient
	writer      *writer.FileCreator
	requests    chan *v1.SendFileRequest
	id          int32
	workers     int
}

func (b *bridgeService) Run(ctx context.Context) error {
	stream, err := b.proxyClient.RegisterServer(ctx, &v1.Auth{
		Id: b.id,
	})
	if err != nil {
		return err
	}

	// Start workers
	for i := 0; i < b.workers; i++ {
		go func() {
			for work := range b.requests {
				b.saveFile(ctx, work)
			}
			log.Trace().Msg("exit from worker")
		}()
	}

	for {
		req, err := stream.Recv()
		if err != nil {
			// TODO: validate code, add reconnects
			return err
		}

		select {
		case b.requests <- req:
			log.Trace().Msg("send job to pool")
		case <-ctx.Done():
			close(b.requests)
			log.Trace().Msg("receive ctx cancel, stoping service")
			return nil
		}
	}
}

func (b *bridgeService) saveFile(ctx context.Context, req *v1.SendFileRequest) {
	logger := log.With().Str("req", req.String()).Logger()
	logger.Trace().Msg("saving file")

	if err := validateRequest(req); err != nil {
		logger.Debug().Err(err).Msg("error in validation")
	}

	d, err := b.writer.CreateFile(req.FileName)
	if err != nil {
		logger.Debug().Err(err).Msg("can't create file")
	}

	stream, err := b.proxyClient.ReceiveChunks(ctx, &v1.SendFileRequest{
		FileName:    req.FileName,
		Offset:      req.Offset,
		Destination: req.Destination,
	})
	if err != nil {
		logger.Debug().Err(err).Msg("can't save file skiping")
	}

	chunks := 0
	offset := int64(0)

	for {
		reqWithChunk, err := stream.Recv()
		if err == io.EOF {
			logger.Trace().Msg("end of the stream")
			return
		}
		if err != nil {
			logger.Debug().Err(err).Msg("stream recv return err")
			return
		}

		c := reqWithChunk.GetChunk()

		err = d.WriteChunk(offset, c)
		if err != nil {
			logger.Debug().Err(err).Msg("can't write chunk to file")
		}
		chunks++
		offset += int64(len(c))
	}
}

func NewBridgeService(
	client v1.TransparentDataRelayServiceClient,
	writer *writer.FileCreator,
	id int32,
	workers int,
) *bridgeService {
	return &bridgeService{
		proxyClient: client,
		writer:      writer,
		requests:    make(chan *v1.SendFileRequest, workers),
		id:          id,
		workers:     workers,
	}
}
