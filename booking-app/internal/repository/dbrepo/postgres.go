package dbrepo

import (
	"context"
	"time"

	"github.com/isoment/booking-app/internal/models"
)

/*
Get an index of users from the database
*/
func (m *postgresDBRepo) AllUsers() bool {
	return true
}

/*
Insert a reservation into the database
*/
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// If the database operation does not complete within 3 seconds cancel it.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int

	stmt := `insert into reservations (
				first_name, last_name, email, phone, start_date, end_date,
				room_id, created_at, updated_at
			) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	// We can pass in the context when executing the query
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil
}

/*
Insert a room restriction into the database
*/
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (
				start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id
			) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomId,
		r.ReservationId,
		time.Now(),
		time.Now(),
		r.RestrictionId,
	)

	if err != nil {
		return err
	}

	return nil
}
