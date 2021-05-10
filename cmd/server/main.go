package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KennyChenFight/gin-starter/pkg/server"

	"go.uber.org/zap"

	"github.com/KennyChenFight/gin-starter/pkg/middleware"
	"github.com/KennyChenFight/gin-starter/pkg/validation"
	"github.com/KennyChenFight/golib/loglib"
	"github.com/KennyChenFight/golib/migrationlib"
	"github.com/KennyChenFight/golib/pglib"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"

	"github.com/KennyChenFight/gin-starter/pkg/dao"
	"github.com/KennyChenFight/gin-starter/pkg/service"
	"github.com/gin-gonic/gin"

	"github.com/jessevdk/go-flags"
)

type PostgresConfig struct {
	URL              string `long:"url" description:"database url" env:"URL" required:"true"`
	PoolSize         int    `long:"pool-size" description:"database pool size" env:"POOL_SIZE" default:"10"`
	MigrationFileDir string `long:"migration-file-dir" description:"migration file dir" env:"MIGRATION_FILE_DIR" default:"file://migrations"`
}

type GinConfig struct {
	Port string `long:"port" description:"port" env:"PORT" default:":8080"`
	Mode string `long:"mode" description:"mode" env:"MODE" default:"debug"`
}

type Environment struct {
	GinConfig      GinConfig      `group:"gin" namespace:"Gin" env-namespace:"GIN"`
	PostgresConfig PostgresConfig `group:"postgres" namespace:"postgres" env-namespace:"POSTGRES"`
}

func main() {
	var env Environment
	parser := flags.NewParser(&env, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	migrationLib := migrationlib.NewMigrateLib(migrationlib.Config{
		DatabaseDriver: migrationlib.PostgresDriver,
		DatabaseURL:    env.PostgresConfig.URL,
		SourceDriver:   migrationlib.FileDriver,
		SourceURL:      env.PostgresConfig.MigrationFileDir,
		TableName:      "migrate_version",
	})

	if err := migrationLib.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("run database migration fail:%v", err)
	}

	pgClient, err := pglib.NewDefaultGOPGClient(pglib.GOPGConfig{
		URL:       env.PostgresConfig.URL,
		DebugMode: false,
		PoolSize:  env.PostgresConfig.PoolSize,
	})

	logger, err := loglib.NewProductionLogger()
	if err != nil {
		log.Fatalf("fail to init logger:%v", err)
	}

	memberDAO := dao.NewPGMemberDAO(logger, pgClient)

	bindingValidator, _ := binding.Validator.Engine().(*validator.Validate)
	CustomValidator, err := validation.NewValidationTranslator(bindingValidator, "en")
	if err != nil {
		log.Fatalf("fail to init validation translator:%v", err)
	}

	svc := service.NewService(memberDAO)

	mwe := middleware.NewMiddleware(logger, CustomValidator)

	gin.SetMode(env.GinConfig.Mode)
	GracefulRun(logger, StartFunc(logger, server.NewHTTPServer(gin.Default(), env.GinConfig.Port, mwe, svc)))
}

func StartFunc(logger *loglib.Logger, server *http.Server) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("http server listen error", zap.Error(err))
			}
		}()

		<-ctx.Done()

		ctx1, cancel1 := context.WithCancel(context.Background())
		go func() {
			logger.Info("shutdown http server...")
			server.Shutdown(ctx1)
			cancel1()
		}()
		<-ctx1.Done()
		logger.Info("http server existing")
		return nil
	}
}

func GracefulRun(logger *loglib.Logger, fn func(ctx context.Context) error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)

	go func() {
		done <- fn(ctx)
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
		return
	case <-shutdown:
		cancel()
		timeoutCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		select {
		case <-done:
			return
		case <-timeoutCtx.Done():
			logger.Error("shutdown timeout", zap.Error(timeoutCtx.Err()))
			return
		}
	}
}
