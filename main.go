package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/GateHubNet/DataAPI/gapi"
	"github.com/GateHubNet/DataAPI/pb"
	"github.com/GateHubNet/DataAPI/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	go runGatewayServer(config)
	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	// lets here defer BQ connection to be closed
	// this means that we have one connection per server...not really 100% sure that this is ok
	defer server.Store.Close()

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterGateHubDataAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create grpc listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	// lets here defer BQ connection to be closed
	// this means that we have one connection per server...not really 100% sure that this is ok
	// we createt different connection for gRPC and different for HTTP
	defer server.Store.Close()

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterGateHubDataAPIHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create HTTP listener")
	}

	log.Info().Msgf("start HTTP server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP server")
	}
}
