package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the users model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is the rooms model
type Rooms struct {
	Id        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the restrictions model
type Restrictions struct {
	ID              int
	ReservationName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the reservations model
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RoomRestrictions is the RoomRestrictions model
type RoomRestrictions struct {
	ID            int
	RoomID        int
	ReservationID int
	RestrictionID int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Reservations
	Restriction   Restrictions
}
