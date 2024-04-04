package postgresql

import (
	"database/sql"

	"github.com/SawitProRecruitment/UserService/repository"
)

type postgresqlRepository struct {
	DB *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) repository.PostgresqlRepositoryInterface {
	return &postgresqlRepository{
		DB: db,
	}
}
