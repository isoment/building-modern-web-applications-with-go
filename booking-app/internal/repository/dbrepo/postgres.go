package dbrepo

import (
	"context"
	"time"

	"github.com/isoment/booking-app/internal/models"
)

// Get an index of users from the database
func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// Insert a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	// If the database operation does not complete within 3 seconds cancel it.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations (
				first_name, last_name, email, phone, start_date, end_date,
				room_id, created_at, updated_at
			) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	// We can pass in the context when executing the query
	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
