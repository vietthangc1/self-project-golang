package mysql

import (
	"log"
	"testing"

	"github.com/thangpham4/self-project/infra"
	"gorm.io/gorm"
)

var (
	conn      *gorm.DB
	userAdmin *UserAdminMysql
)

func TestMain(m *testing.M) {
	var err error

	conn, err = infra.NewMySQLConnectionTest()
	if err != nil {
		log.Fatal(err)
	}

	userAdmin = NewUserAdminMysql(conn)
}
