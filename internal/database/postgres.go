package database

import (
    "log"

    "github.com/gift-redemption/internal/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func NewPostgresConnection(cfg *config.Config) *gorm.DB {
    gormCfg := &gorm.Config{}

    if cfg.AppEnv == "development" {
        gormCfg.Logger = logger.Default.LogMode(logger.Info)
    } else {
        gormCfg.Logger = logger.Default.LogMode(logger.Error)
    }

    db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), gormCfg)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("failed to get sql.DB: %v", err)
    }

    sqlDB.SetMaxOpenConns(25)
    sqlDB.SetMaxIdleConns(10)

    log.Println("database connection established")
    return db
}