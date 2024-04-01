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
