package repository

import (
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/logger"
	"XDSEC2022-Backend/src/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

var modelsInit []interface{}

func Initialize() error {
	logger.Info("Initializing database...")
	l := logger.DBLoggerNew(zap.L())
	l.SetAsDefault()
	sqlConfigSrc := config.DatabaseConfig
	sqlConfig := postgres.Config{
		DSN: sqlConfigSrc.Dsn(),
	}
	var err error
	Database, err = gorm.Open(postgres.New(sqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   l,
	})
	if err != nil {
		return err
	}
	logger.Info("Database initialized, recovering state...")
	err = initTables()
	if err != nil {
		return err
	}
	logger.Info("Database state recovered.")
	return nil
}

func initTables() error {
	if err := Database.AutoMigrate(modelsInit...); err != nil {
		return err
	}
	Database.Config.DisableForeignKeyConstraintWhenMigrating = false
	if err := Database.AutoMigrate(modelsInit...); err != nil {
		return err
	}

	// test user here, if no user has registered, set a flag that make the first registration to be an admin user
	var userCount int64
	err := Database.Model(&model.User{}).Count(&userCount).Error
	if err != nil {
		return err
	}
	if userCount == 0 {
		UserInitFlag = true
	}
	return nil
}

func RegisterModel(model interface{}) {
	modelsInit = append(modelsInit, model)
}
