package repository

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"xdsec-join/src/config"
	logger2 "xdsec-join/src/logger"
	"xdsec-join/src/model"
)

var Database *gorm.DB

var modelsInit []interface{}

func Initialize() error {
	logger2.Info("Initializing database...")
	l := logger2.DBLoggerNew(zap.L())
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
	logger2.Info("Database initialized, recovering state...")
	err = initTables()
	if err != nil {
		return err
	}
	logger2.Info("Database state recovered.")
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
