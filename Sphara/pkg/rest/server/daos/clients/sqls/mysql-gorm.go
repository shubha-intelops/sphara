package sqls

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	user     = os.Getenv("MYSQL_DB_USER")
	password = os.Getenv("MYSQL_DB_PASSWORD")
	host     = os.Getenv("MYSQL_DB_HOST")
	port     = os.Getenv("MYSQL_DB_PORT")
	database = os.Getenv("MYSQL_DB_DATABASE")
)

type MySQLClient struct {
	DB *gorm.DB
}

var mySQLClient *MySQLClient

func InitGormMySQLDB() (*MySQLClient, error) {
	var o sync.Once
	var err error
	var db *gorm.DB

	o.Do(func() {
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

		// Please modify your configs based on the your requirements as  below comments
		mySQLConfigs := mysql.Config{
			DSN:                       dataSource, // data source name
			DefaultStringSize:         256,        // default size for string fields
			DisableDatetimePrecision:  true,       // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false,      // auto configure based on currently MySQL version
		}

		db, err = gorm.Open(mysql.New(mySQLConfigs), &gorm.Config{})

		if err != nil {
			log.Debugf("database connection error, %v", err)
			os.Exit(1)
		}

		mySQLClient = &MySQLClient{
			DB: db,
		}
	})

	return mySQLClient, nil
}
