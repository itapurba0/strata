package organization

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PermissionRepository struct {
	db DBTX
}

func NewPermissionRepository(db *pgxpool.Pool) *PermissionRepository {
	return &PermissionRepository{
		db: db,
	}
}
