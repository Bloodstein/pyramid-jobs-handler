package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Bloodstein/pyramid-jobs-handler/domain"
	_ "github.com/lib/pq"
)

type PostgreSQLRepository struct {
	db *sql.DB
}

func (postgres PostgreSQLRepository) createConnection(config domain.PostgreSQLDBConfiguration) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Database))

	if err != nil {
		log.Fatalf("Fail to open PostgreSQL connection: %s", err.Error())
	}

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	postgres.db = db
}

func NewPostgreSQLDB(config domain.PostgreSQLDBConfiguration) PostgreSQLRepository {
	postgres := PostgreSQLRepository{}
	postgres.createConnection(config)

	return postgres
}

func (postgres PostgreSQLRepository) PopJob() {

}

func (postgres PostgreSQLRepository) StoreJob() {

}

func (postgres PostgreSQLRepository) SaveResult() {

}
