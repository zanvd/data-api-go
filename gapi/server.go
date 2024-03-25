package gapi

import (
	"github.com/GateHubNet/DataAPI/pb"
	"github.com/GateHubNet/DataAPI/util"
	"github.com/go-redis/redis"
	"github.com/gocql/gocql"
	"github.com/rs/zerolog/log"
	"github.com/scylladb/gocqlx/v2"
	"strconv"
	"strings"
)

// Server serves gRPC requests for our GateHub Data Api service.
type Server struct {
	pb.UnimplementedGateHubDataAPIServer
	config util.Config
	Store  gocqlx.Session
	Cache  *redis.Client

	ledgerCacheKey string
}

// NewServer creates a new gRPC server
func NewServer(config util.Config) (*Server, error) {
	parsedScyllaPort, err := strconv.Atoi(config.ScyllaPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot convert scylla port to number")
	}

	scyllaHosts := strings.Split(config.ScyllaHosts, ",")

	cluster := gocql.NewCluster(scyllaHosts...)
	cluster.Keyspace = config.ScyllaKeyspace
	cluster.Port = parsedScyllaPort

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create ScyllaDB session")
	}

	redisDatabase, _ := strconv.Atoi(config.RedisDatabase)

	cache := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       redisDatabase,
	})

	server := &Server{
		config: config,
		Store:  session,
		Cache:  cache,
	}

	server.ledgerCacheKey = "ledgers:%s"

	return server, nil
}
