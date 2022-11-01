package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"

	v1 "homework/internal/api/v1"
	"homework/internal/config"
	"homework/specs"
)

func main() {
	var (
		err         error
		ctx, cancel = signal.NotifyContext(
			context.Background(),
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
	)
	defer cancel()

	cfg, err := config.InitConfig(os.Args)
	if err != nil {
		log.Fatal("get config: ", err.Error())
		return
	}

	apiServer := v1.NewAPIServer()

	err = startHTTPServer(ctx, cfg, apiServer)
	if err != nil {
		log.Fatal("starting server: ", err.Error())
	}
}

func startHTTPServer(
	ctx context.Context,
	cfg *config.Config,
	apiServer specs.ServerInterface,
	middlewares ...specs.MiddlewareFunc,
) error {
	handler := specs.HandlerWithOptions(apiServer, specs.ChiServerOptions{
		BaseURL:     cfg.BasePath,
		Middlewares: middlewares,
	})

	router := chi.NewRouter()
	router.Handle("/*", handler)

	httpServer := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return httpServer.Shutdown(ctx)
	})

	return group.Wait()
}
