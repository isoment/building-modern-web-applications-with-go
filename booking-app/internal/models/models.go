package models

import "time"

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

/**********************
	Database Models
**********************/

// A model for the users table
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// A model for the rooms table
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// A model for the restrictions table
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// A model for the reservations table, notice we can add the Room Model
// for when we want to include the associated room.
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// A model for the room_restrictions table
type RoomRestrictions struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomId        int
	ReservationId int
	RestrictionId int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Reservations
	Restriction   Restrictions
}
