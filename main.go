package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ecadlabs/tezos-streamer/config"
	"github.com/ecadlabs/tezos-streamer/service"
	"github.com/ecadlabs/tezos-streamer/streamer"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		config     config.Config
		configFile string
	)

	flag.StringVar(&configFile, "c", "", "Config file.")
	flag.StringVar(&config.HTTPAddress, "address", ":8002", "HTTP address to listen on.")
	flag.StringVar(&config.RPCUrl, "rpc", "https://mainnet-node.tzscan.io", "RPC Url")

	flag.Parse()

	if configFile != "" {
		if err := config.Load(configFile); err != nil {
			log.Fatal(err)
		}
		// Override from command line
		flag.Parse()
	}

	str, err := streamer.NewStreamer(&config)

	go str.Start()

	if err != nil {
		log.Fatal(err)
	}

	svc, err := service.NewService(str)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    config.HTTPAddress,
		Handler: svc.NewAPIHandler(),
	}

	log.Printf("HTTP server listening on %s", config.HTTPAddress)

	errChan := make(chan error)
	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		httpServer.Shutdown(ctx)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}

		case s := <-signalChan:
			log.Printf("Captured %v. Exiting...\n", s)
			return
		}
	}
}
