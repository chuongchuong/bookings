package dbrepo

import (
	"errors"
	"time"

	"github.com/chuongchuong/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room into database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// SearchAvailabitityByDateByRoomID returns true if availability exists for roomID, and returns false if no availability exists
func (m *testDBRepo) SearchAvailabitityByDateByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabitityForAllRooms returns a slice of available rooms, if any,for given date range
func (m *testDBRepo) SearchAvailabitityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {

	return 1, "", nil
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// AllNewReservations returns a slice of all reservation
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// GetRerservationByID returns one reservation by ID
func (m *testDBRepo) GetRerservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

// UpdateReservation updates a reservation in database
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// Delete Reservations by ID
func (m *testDBRepo) DeleteReservation (id int) error{
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by ID
func (m *testDBRepo) UpdateProcessedForReservation(id,processed int)error{

	return nil
}