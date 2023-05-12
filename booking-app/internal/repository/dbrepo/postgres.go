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

	stmt := `INSERT INTO reservations (
				first_name, last_name, email, phone, start_date, end_date,
				room_id, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

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

	stmt := `INSERT INTO room_restrictions (
				start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id
			) VALUES ($1, $2, $3, $4, $5, $6, $7)`

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

/*
Determine if a room is available for booking on the given dates, true is returned if the
room is available, otherwise false is returned.
*/
func (m *postgresDBRepo) SearchRoomAvailabilityByDates(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	// $2 is start data and $3 is end date
	query := `
		SELECT 
			count(id)
		FROM room_restrictions rr 
		WHERE
			room_id = $1
		AND
			$2 < rr.end_date
		AND 
			$3 > rr.start_date`

	row := m.DB.QueryRowContext(ctx, query, roomId, start, end)

	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

/*
Searches room availability for the given dates and returns a collection
of available rooms if there are any.
*/
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		SELECT
			r.id, r.room_name
		FROM rooms r
		WHERE r.id NOT IN (
			SELECT rr.room_id 
			FROM room_restrictions rr
			WHERE $1 < rr.end_date
			AND $2 > rr.start_date
		)`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

/*
Get a room by id
*/
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `
		SELECT 
			id, room_name, created_at, updated_at 
		FROM rooms
		WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}
