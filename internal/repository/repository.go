package repository

import (
	"time"

	"github.com/chuongchuong/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabitityByDateByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabitityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}
