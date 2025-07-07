package initializer

import (
	"errors"
	"fmt"
	"log"

	"github.com/alirezasaharkhiz/FlagChain/config"
	"github.com/alirezasaharkhiz/FlagChain/controllers"
	"github.com/alirezasaharkhiz/FlagChain/middlewares"
	"github.com/alirezasaharkhiz/FlagChain/repositories"
	"github.com/alirezasaharkhiz/FlagChain/routes"
	"github.com/alirezasaharkhiz/FlagChain/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	cfg := config.LoadConfig()

	databaseURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	m, err := migrate.New("file://"+cfg.MigrationsDir, databaseURL)
	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration up failed: %v", err)
		}
	}

	gormDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	flagRepo := repositories.NewFlagRepository(db)
	depRepo := repositories.NewDependencyRepository(db)
	auditRepo := repositories.NewAuditRepository(db)

	service := services.NewFeatureFlagService(flagRepo, depRepo, auditRepo)
	controller := controllers.NewFeatureFlagController(service)

	router := gin.Default()
	router.Use(middlewares.ErrorHandler)
	routes.RegisterFlagRoutes(router, controller)

	return &App{Router: router}
}

func (a *App) Run() {
	cfg := config.LoadConfig()
	if err := a.Router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
