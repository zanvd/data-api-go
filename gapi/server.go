package gapi

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/GateHubNet/DataAPI/pb"
	"github.com/GateHubNet/DataAPI/util"
	"github.com/rs/zerolog/log"
)

// Server serves gRPC requests for our GateHub Data Api service.
type Server struct {
	pb.UnimplementedGateHubDataAPIServer
	config util.Config
	Store  *bigquery.Client
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config) (*Server, error) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, config.GCPBigQueryProjectId)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize BQ client")
	}

	server := &Server{
		config: config,
		Store:  client,
	}

	return server, nil
}
