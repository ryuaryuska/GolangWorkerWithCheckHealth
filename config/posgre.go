
package config

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"
	"WorkerWithCheckHealth/exception"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgreConnection() *gorm.DB {
	ctx, cancel := NewPostgreContext()
	defer cancel()

	sqlDB, err := sql.Open("postgres", os.Getenv("POSTGRE_HOST"))
	exception.PanicIfNeeded(err)

	err = sqlDB.PingContext(ctx)
	exception.PanicIfNeeded(err)

	postgrePoolMax, err := strconv.Atoi(os.Getenv("POSTGRE_POOL_MAX"))
	exception.PanicIfNeeded(err)

	postgreIdleMax, err := strconv.Atoi(os.Getenv("POSTGRE_IDLE_MAX"))
	exception.PanicIfNeeded(err)

	postgreMaxLifeTime, err := strconv.Atoi(os.Getenv("POSTGRE_MAX_LIFE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

	// mysqlMaxIdleTime, err := strconv.Atoi(os.Getenv("POSTGRE_MAX_IDLE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(postgrePoolMax)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(postgreIdleMax)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(postgreMaxLifeTime) * time.Minute)

	//sqlDB.SetConnMaxIdleTime(time.Duration(mysqlMaxIdleTime) * time.Minute)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	exception.PanicIfNeeded(err)
	return gormDB
}
func NewPostgreContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
