package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/bmstu-itstech/contest-auth/config"
	"github.com/bmstu-itstech/contest-auth/pkg/closer"
	pb "github.com/bmstu-itstech/contest-auth/pkg/user_v1"
)

type App struct {
	cfg             *config.Config
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates a new instance of App.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run runs the App.
func (a *App) Run(ctx context.Context, cancel context.CancelFunc) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Starting gRPC server
	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Panic(err)
		}
	}()

	// Handle graceful shutdown
	gracefulShutdown(ctx, cancel, wg)

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	log.Info("Config loaded")

	a.cfg = cfg

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.cfg)

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	pb.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserAPI(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	listener, err := net.Listen(
		"tcp",
		a.cfg.GRPC.Address(),
	)
	if err != nil {
		return errors.Wrapf(err, "Error starting listener")
	}

	log.Infof("Starting gRPC server on %s", a.cfg.GRPC.Address())

	if err := a.grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "Error starting gRPC server")
	}

	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Info("terminating: context cancelled")
	case <-waitSignal():
		log.Infof("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
