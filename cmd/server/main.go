package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/KennyChenFight/gin-starter/pkg/dao"
	"github.com/KennyChenFight/gin-starter/pkg/route"
	"github.com/KennyChenFight/gin-starter/pkg/service"
	"github.com/KennyChenFight/gin-starter/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/jessevdk/go-flags"
)

type PostgresConfig struct {
	URL              string `long:"url" description:"database url" env:"URL" required:"true"`
	MigrationFileDir string `long:"migration-file-dir" description:"migration file dir" env:"MIGRATION_FILE_DIR" required:"true"`
}

type GinConfig struct {
	Port string `long:"port" description:"port" env:"PORT" default:"8080"`
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

	migrationConfig := util.MigrationConfig{
		Driver:      util.PGDriver,
		DatabaseURL: env.PostgresConfig.URL,
		FileDir:     env.PostgresConfig.MigrationFileDir,
	}
	if err := util.RunMigrations(migrationConfig); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("run database migration fail:%v", err)
	}

	opts, err := pg.ParseURL(env.PostgresConfig.URL)
	if err != nil {
		log.Fatalf("fail to parse pg url:%v", err)
	}
	db := pg.Connect(opts)
	memberDAO := dao.NewPGMemberDAO(db)

	bindingValidator, _ := binding.Validator.Engine().(*validator.Validate)
	CustomValidator, err := util.NewValidationTranslator(bindingValidator, "en")
	if err != nil {
		log.Fatalf("fail to init validation translator:%v", err)
	}

	svc := service.NewService(memberDAO, CustomValidator)

	gin.SetMode(env.GinConfig.Mode)
	server := gin.Default()
	route.InitRoutingRule(server, svc).
		Run(fmt.Sprintf(":%s", env.GinConfig.Port))
}
