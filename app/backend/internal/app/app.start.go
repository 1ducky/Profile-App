package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"profile-template/infra/storage/s3"
	errApp "profile-template/internal/app/err"
	"profile-template/internal/app/registry"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/transport"
	"syscall"
	"time"
)

func StartApp() {
	startAt := time.Now()
	log.Print("Starting App")
	cfg, _ := config.LoadConfig()
	log.Print("Loading Config")
	ctx := context.Background()
	log.Print("Creating Context")
	db, err := db.NewMySQLDatabase(&cfg.DBConfig)
	log.Print("Creating DB")
	if err != nil {
		panic(err)
	}
	bucket, err := s3.NewS3Storage(ctx, cfg.StorageConfig)
	log.Print("initalize Storage")
	if err != nil {
		panic(err)
	}

	// Registry
	log.Print("Creating Err Registry")
	errs := registry.NewErrRegistry(registry.RegisterErr{AuthKey: errApp.Auth, AssetKey: errApp.Assets, ProfileKey: errApp.Profile, ArticleKey: errApp.Article})
	log.Print("Creating Adapter")
	adapter := transport.NewAdapter(errs)
	log.Print("Creating Repo")
	persistent := registry.NewRegisterPersistent(db, bucket)
	log.Print("Creating UC")
	repos := registry.NewRegisterRepo(persistent, cfg)
	log.Print("Creating UC")
	ucs := registry.NewRegisterUC(repos, cfg) //
	log.Print("Creating UoW")
	uows := registry.NewUoWCase(db, cfg)
	log.Print("Creating App")
	apps := registry.NewRegisterApp(ucs, uows)
	log.Print("Creating Handler")
	handlers := registry.NewRegisterHandler(apps)
	log.Print("Creating Routes")
	routes := registry.NewRegisterRoutes(handlers, ucs, adapter)
	log.Print("Creating Transport")
	httpServer := transport.NewTransport(cfg.HTTPConfig, routes)

	go func(ctx context.Context) {
		// Graceful Shutdown
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}

		log.Println("Server stopped gracefully")
	}(ctx)

	log.Print("Starting Transport")
	donetime := time.Since(startAt)
	log.Printf("App started in %s", donetime.String())
	log.Printf("Server address: %s", httpServer.Address())
	httpServer.Start()

	defer func() {
		db.Close()
		log.Print("Closing DB")
	}()

}
