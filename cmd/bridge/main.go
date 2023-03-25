package main

import (
	_ "embed"
	"net"

	"github.com/rs/zerolog/log"

	v1 "github.com/Big-Kotik/transparent-data-bridge-api/bridge/api/v1"
	bridgeservice "github.com/Big-Kotik/transparent-data-bridge-server/internal/bridge_service"
	"github.com/Big-Kotik/transparent-data-bridge-server/internal/config"
	"github.com/Big-Kotik/transparent-data-bridge-server/internal/writer"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

//go:embed config.yaml
var cfgFile string

// TODO: tests, log, docker
func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	cfg, err := config.ConfigFromFile(cfgFile)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("can't parse config")
	}

	w := writer.NewFileCreator(cfg.Writer.BasicDir, int(cfg.Writer.ChunkSize))
	lis, err := net.Listen("tcp", cfg.Endpoint)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("can't create socket")
	}

	bs := bridgeservice.NewBridgeService(w)

	grpcServer := grpc.NewServer()
	v1.RegisterTransparentDataBridgeServiceServer(grpcServer, bs)

	log.Trace().Msg("starting grpc server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().
			Err(err).
			Msg("error while serving tcp conn")
	}
}
