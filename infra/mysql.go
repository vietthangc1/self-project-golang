package infra

import (
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection() (*gorm.DB, error) {
	l := logger.Factory("Setup Mysql")
	dsn := envx.String("MYSQL_ADDR", "")
	return NewMySQL(dsn, l)
}

func NewMySQLConnectionTest() (*gorm.DB, error) {
	l := logger.Factory("Setup Mysql Dev")
	dsn := envx.String(
		"MYSQL_ADDR_TEST",
		"root:Chaugn@rs2@tcp(127.0.0.1:3306)/self_project_dev?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
	)
	return NewMySQL(dsn, l)
}

func NewMySQL(dsn string, l logger.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.V(logger.LogErrorLevel).Error(err, "failed to set up mysql", "mysql_uri", dsn)
		return nil, err
	}
	l.V(logger.LogInfoLevel).Info("successfully set up mysql", "mysql_uri", dsn)

	err = db.AutoMigrate(
		&entities.UserAdmin{},
		&entities.ProductInfo{},
		&entities.CustomerInfo{},
		&entities.ModelSource{},
		&entities.ModelInfo{},
		&entities.OrderInfo{},
	)

	if err != nil {
		l.V(logger.LogErrorLevel).Error(err, "error in migrating database")
	}

	return db, nil
}
