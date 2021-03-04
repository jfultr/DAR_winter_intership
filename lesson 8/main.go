package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"

	"./user"
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
	var srv user.Service
	{
		repository := user.NewRepo(db, logger)

		srv = user.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	endpoints := user.MakeEndpoints(srv)
	grpcServer := user.NewGRPCServer(ctx, endpoints)

	go func() {
		baseServer := grpc.NewServer()
		user.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
