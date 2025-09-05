package repositories

import (
	"RiderApi/internal/domain"
	"database/sql"
)

type RiderRepository interface {
	GetAll() ([]domain.KamenRider, error)
	GetById(id int) (domain.KamenRider, error)
	Post(rider domain.KamenRider) (domain.KamenRider, error)
	Patch(rider domain.KamenRider, id int) (domain.KamenRider, error)
	Delete(id int) (int, error)
}

type riderRepository struct {
	db *sql.DB
}

func NewRiderRepository(d *sql.DB) RiderRepository {

	return &riderRepository{
		db: d,
	}
}

func (r *riderRepository) GetAll() ([]domain.KamenRider, error) {

	query := "SELECT * FROM Riders"

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

func (r *riderRepository) GetById(id int) (domain.KamenRider, error) {
	query := "SELECT * FROM Riders WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var rider domain.KamenRider

	if err := row.Scan(&rider.ID, &rider.Name, &rider.Henshin, &rider.Kick); err != nil {
		return domain.KamenRider{}, err
	}

	return rider, nil

}

func (r *riderRepository) Post(rider domain.KamenRider) (domain.KamenRider, error) {

	query := "INSERT INTO RIDER (name,henshin,kick) VALUES($1,$2,$3) RETURNING id"

	err := r.db.QueryRow(query, rider.Name, rider.Henshin, rider.Kick).Scan(&rider.ID)

	if err != nil {
		return domain.KamenRider{}, err
	}

	return rider, nil

}

func (r *riderRepository) Patch(rider domain.KamenRider, id int) (domain.KamenRider, error) {

	query := "UPDATE Riders set name=$1 , henshin=$2 ,kick=$3"

	_, err := r.db.Exec(query)

	if err != nil {
		return domain.KamenRider{}, err
	}

	return rider, nil

}

func (r *riderRepository) Delete(id int) (int, error) {

	query := "DELETE FROM Riders WHERE id = $1"

	_, err := r.db.Exec(query, id)

	if err != nil {
		return 0, err
	}

	return id, nil

}
