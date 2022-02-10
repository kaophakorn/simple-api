package provider

import (
	"database/sql"
	"log"
	"net/url"
	"sync"
	"time"

	gormDriverMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var cfg *Config
var gormConnections map[string]*gorm.DB
var dbConnections map[string]*sql.DB
var m1 sync.Mutex

type Config struct {
	Port  string
	MySQL map[string]ConfigMySQL
}

type ConfigMySQL struct {
	Name     string
	Host     string
	Port     string
	User     string
	Pass     string
	Timezone string
	Database string
}

func init() {
	gormConnections = make(map[string]*gorm.DB)
	dbConnections = make(map[string]*sql.DB)
	cfg = &Config{
		Port: "5000",
		MySQL: map[string]ConfigMySQL{
			"main": {
				Host:     "host.docker.internal",
				Port:     "3306",
				User:     "root",
				Pass:     "root",
				Timezone: "Asia/Bangkok",
				Database: "main",
			},
		},
	}
	initDB()
	initGorm()
}

func GetConfig() *Config {
	return cfg
}

func initGorm() {
	// prepare gorm
	var err error
	gormConnections["main"], err = gorm.Open(gormDriverMySQL.New(gormDriverMySQL.Config{
		Conn: dbConnections["main"],
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic(err)
	}
}
func initDB() {
	dbConfigs := GetConfig().MySQL
	for key, val := range dbConfigs {
		connectionString := val.User + ":" + val.Pass + "@tcp(" + val.Host + ":" + val.Port + ")/" + val.Database + "?timeout=1m&readTimeout=1m&writeTimeout=1m&charset=utf8mb4&parseTime=True&loc=" + url.QueryEscape(val.Timezone)
		var err error
		sqlDB, err := sql.Open("mysql", connectionString)
		if err != nil {
			log.Panic(err)
		}
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(100)
		dbConnections[key] = sqlDB
	}
}

func GetGormMain() *gorm.DB {
	return gormConnections["main"]
}
