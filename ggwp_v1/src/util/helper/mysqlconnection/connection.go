package db

import (
	"fmt"
	"os"
	"time"

	// "log"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// DB Global Connetction
var DB *gorm.DB

// TransactionDB ...
var TransactionDB *gorm.DB

// DBInit Initialization Connection
// return connection, error
func DBInit() (*gorm.DB, error) {
	godotenv.Load()
	mysqlCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_MYSQL_USERNAME"),
		os.Getenv("DB_MYSQL_PASSWORD"),
		os.Getenv("DB_MYSQL_HOST"),
		os.Getenv("DB_MYSQL_PORT"),
		os.Getenv("DB_MYSQL_DATABASE"),
	)
	var err error
	DB, err = gorm.Open("mysql", mysqlCon)

	if err != nil {
		// fmt.Println(fmt.Sprintf("Failed connected to database %s", mysqlCon))
		// return DB, err
		panic(fmt.Sprintf("Failed connected to database %s", mysqlCon))
	}

	fmt.Println(fmt.Sprintf("Successfully connected to database %s", mysqlCon))
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(200)
	DB.LogMode(true)

	log, err := zap.NewProduction()
	DB.SetLogger(CustomLogger(log))
	// DB.SetLogger(log.New(os.Stdout, "\r\n", 0))
	fmt.Println("Connection is created")

	return DB, err
}

// GetConnection function
// return connection
func GetConnection() *gorm.DB {
	if DB == nil {
		fmt.Println("No Active Connection Found")
		DB, _ = DBInit()
	}
	return DB
}

// GetTransactionConnection function
// return DB.Begin()
func GetTransactionConnection() *gorm.DB {
	if TransactionDB == nil {
		fmt.Println("No Active Connection Found")
		TransactionDB, _ = DBInit()
	}
	return TransactionDB
}

// CustomLogger params
// @zap: *zap.Logger
// return: *Logger
func CustomLogger(zap *zap.Logger) *Logger {
	return &Logger{
		zap: zap,
	}
}

// Logger derivated zap logger
type Logger struct {
	zap *zap.Logger
}

// Print logger view
// @values: interface
func (l *Logger) Print(values ...interface{}) {
	var additionalString = ""
	for _, item := range values {
		if _, ok := item.(string); ok {
			additionalString = additionalString + fmt.Sprintf("\n%v", item)
		}
		if err, ok := item.(*mysql.MySQLError); ok {
			err.Message = err.Message + additionalString
		}
	}
}
