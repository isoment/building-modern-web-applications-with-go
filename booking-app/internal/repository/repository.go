package repository

import (
	"time"

	"github.com/isoment/booking-app/internal/models"
)

// Anytime we create a new function we can add it to this interface. We can access
// this in our handlers.
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomId int) (bool, error)
}
