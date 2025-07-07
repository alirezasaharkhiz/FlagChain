package initializer

import (
	"github.com/alirezasaharkhiz/FlagChain/config"
	"github.com/alirezasaharkhiz/FlagChain/controllers"
	"github.com/alirezasaharkhiz/FlagChain/middlewares"
	"github.com/alirezasaharkhiz/FlagChain/repositories"
	"github.com/alirezasaharkhiz/FlagChain/routes"
	"github.com/alirezasaharkhiz/FlagChain/services"
	"log"

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

	// اجرای مهاجرت‌ها
	dsn := "mysql://" + cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName + "?multiStatements=true"
	m, err := migrate.New("file://"+cfg.MigrationsDir, dsn)
	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up failed: %v", err)
	}

	// اتصال GORM
	gormDSN := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// مقداردهی اولیه مخازن داده
	flagRepo := repositories.NewFlagRepository(db)
	depRepo := repositories.NewDependencyRepository(db)
	auditRepo := repositories.NewAuditRepository(db)

	// سرویس و کنترلر
	service := services.NewFeatureFlagService(flagRepo, depRepo, auditRepo)
	controller := controllers.NewFeatureFlagController(service)

	// راه‌اندازی Gin و روتر
	router := gin.Default()
	router.Use(middlewares.ErrorHandler)
	routes.RegisterFlagRoutes(router, controller)

	return &App{Router: router}
}

func (a *App) Run() {
	cfg := config.LoadConfig()
	a.Router.Run(cfg.ServerPort)
}
