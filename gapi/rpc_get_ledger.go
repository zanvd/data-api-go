package gapi

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/GateHubNet/DataAPI/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LedgerParseData struct {
	field    string
	operator string
	value    any
}

func (server *Server) parseLedgerInformation(req *pb.GetLedgerRequest) LedgerParseData {
	layout := "2006-01-02T15:04:05Z07:00"
	// lets parse identifier aka what to show

	var ledgerData *LedgerParseData

	if _, err := strconv.Atoi(req.Identifier); err == nil {
		parsed, _ := strconv.Atoi(req.Identifier)
		ledgerData = &LedgerParseData{
			field:    "ledger_index",
			operator: "=",
			value:    parsed,
		}
	} else if _, err := time.Parse(layout, req.Identifier); err == nil {
		dt, _ := time.Parse(layout, req.Identifier)
		ledgerData = &LedgerParseData{
			field:    "close_time",
			operator: ">",
			value:    dt.Unix(),
		}
	} else {
		ledgerData = &LedgerParseData{
			field:    "hash",
			operator: "=",
			value:    req.Identifier,
		}
	}

	return *ledgerData
}

func (server *Server) getLedgerFromBigQuery(ctx context.Context, q *bigquery.Query) (*pb.GetLedgerResponse, error) {
	response := &pb.GetLedgerResponse{}

	// Run the query and process the returned row iterator.
	it, err := q.Read(ctx)
	if err != nil {
		msg := "internal error when reading query"
		log.Fatal().Err(err).Msg(msg)

		errorResponse := status.New(codes.Internal, msg)
		return nil, status.ErrorProto(errorResponse.Proto())
	}

	var ledger pb.Ledger
	for {
		err = it.Next(&ledger)
		if err == iterator.Done {
			if len(ledger.LedgerHash) > 0 {
				response = &pb.GetLedgerResponse{
					Result: "success",
					Ledger: &ledger,
				}
				return response, nil
			}

			errorResponse := status.New(codes.NotFound, "ledger not found")
			return nil, status.ErrorProto(errorResponse.Proto())
		}

		if err != nil {
			msg := "internal error when checking iterator"
			log.Fatal().Err(err).Msg(msg)

			errorResponse := status.New(codes.Internal, msg)
			return nil, status.ErrorProto(errorResponse.Proto())
		}
	}
}

func (server *Server) GetLastValidatedLedger(ctx context.Context, req *pb.GetLastValidatedLedgerRequest) (*pb.GetLedgerResponse, error) {
	q := server.Store.Query(fmt.Sprintf(
		"SELECT * FROM `%s.%s.%s` ORDER BY close_time DESC LIMIT 1;",
		server.config.GCPBigQueryProjectId,
		server.config.GCPBigQueryDataSet,
		server.config.GCPBigQueryLedgersTable,
	))

	return server.getLedgerFromBigQuery(ctx, q)
}

func (server *Server) GetLedger(ctx context.Context, req *pb.GetLedgerRequest) (*pb.GetLedgerResponse, error) {
	ledgerParseData := server.parseLedgerInformation(req)
	q := server.Store.Query(fmt.Sprintf(
		"SELECT * FROM `%s.%s.%s` WHERE %s %s @value;",
		server.config.GCPBigQueryProjectId,
		server.config.GCPBigQueryDataSet,
		server.config.GCPBigQueryLedgersTable,
		ledgerParseData.field,
		ledgerParseData.operator,
	))

	q.Parameters = []bigquery.QueryParameter{
		{Name: "value", Value: ledgerParseData.value},
	}

	return server.getLedgerFromBigQuery(ctx, q)
}
