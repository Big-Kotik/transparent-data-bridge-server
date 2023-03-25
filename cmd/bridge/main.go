package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"
	"syscall"

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

func signalHandler(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	for signal := range signals {
		if signal == syscall.SIGINT {
			log.Debug().Msg("recive sigint")
			cancel()
		} else {
			log.Debug().Str("signal", signal.String()).Msg("recive unexpected signal")
		}
	}
}

// TODO: tests, log, docker
func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	cfg, err := config.ConfigFromFile(cfgFile)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("can't parse config")
	}

	w := writer.NewFileCreator(cfg.Writer.BasicDir)
	conn, err := grpc.Dial(cfg.ProxyEndpoint)
	defer conn.Close()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("can't create socket")
	}

	proxyClient := v1.NewTransparentDataRelayServiceClient(conn)

	server := bridgeservice.NewBridgeService(
		proxyClient,
		w,
		cfg.Id,
		cfg.Workers,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go signalHandler(cancel)

	log.Trace().Msg("starting server")
	if err := server.Run(ctx); err != nil {
		log.Fatal().
			Err(err).
			Msg("error while serving tcp conn")
	}
}
