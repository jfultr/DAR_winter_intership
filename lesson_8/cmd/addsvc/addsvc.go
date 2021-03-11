package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"

	"github.com/jfultr/DAR_winter_intership/lesson_8/pb"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addendpoint"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addservice"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addtransport"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"
)

const dbsource = "postgresql://postgres:6378@localhost:5432/postgres?sslmode=disable"

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv addservice.Service
	{
		repository := addservice.NewRepo(db, logger)

		srv = addservice.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	endpoints := addendpoint.MakeEndpoints(srv)
	grpcServer := addtransport.NewGRPCServer(ctx, endpoints)

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
