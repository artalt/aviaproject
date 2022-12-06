package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"

	v1Pkg "homework/internal/api/v1"
	"homework/internal/config"
	fServicePkg "homework/internal/service/flight"
	fOrderPkg "homework/internal/service/order"
	fStoragePkg "homework/internal/storage/postgresql/flight"
	oStoragePkg "homework/internal/storage/postgresql/order"
	tStoragePkg "homework/internal/storage/postgresql/ticket"
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

	// инициализация пакета/драйвера БД
	db, err := pgxpool.Connect(ctx, cfg.Db.Postgresql)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()
	// инициализация хранилищ
	flightStorage := fStoragePkg.NewFlightStorage(db)
	orderStorage := oStoragePkg.NewOrderStorage(db)
	ticketStorage := tStoragePkg.NewTicketStorage(db)
	// инициализация сервисов
	flightService := fServicePkg.NewFlightService(flightStorage)
	orderService := fOrderPkg.NewOrderService(orderStorage, ticketStorage)

	apiServer := v1Pkg.NewAPIServer(flightService, orderService)

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
