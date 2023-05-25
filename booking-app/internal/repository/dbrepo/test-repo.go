package dbrepo

import (
	"errors"
	"time"

	"github.com/isoment/booking-app/internal/models"
)

/*
Get an index of users from the database
*/
func (m *testDBRepo) AllUsers() bool {
	return true
}

/*
Insert a reservation into the database
*/
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

/*
Insert a room restriction into the database
*/
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

/*
Determine if a room is available for booking on the given dates, true is returned if the
room is available, otherwise false is returned.
*/
func (m *testDBRepo) SearchRoomAvailabilityByDates(start, end time.Time, roomId int) (bool, error) {
	return false, nil
}

/*
Searches room availability for the given dates and returns a collection
of available rooms if there are any.
*/
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

/*
Get a room by id
*/
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	// Simulate the case where we get a non existent room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}
