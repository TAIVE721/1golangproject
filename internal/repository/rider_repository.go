package repository

import (
	"RiderApi/internal/domain"
	"database/sql"
)

type RiderRepository interface {
	FindALL() ([]domain.KamenRider, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewRiderRepository(db *sql.DB) RiderRepository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) FindALL() ([]domain.KamenRider, error) {

	query := "SELECT id, name, henshin , kick FROM riders"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var riders []domain.KamenRider

	for rows.Next() {
		var rider domain.KamenRider

		if err := rows.Scan(&rider.ID, &rider.Name, &rider.Henshin, &rider.Kick); err != nil {
			return nil, err
		}

		riders = append(riders, rider)
	}

	return riders, nil

}
