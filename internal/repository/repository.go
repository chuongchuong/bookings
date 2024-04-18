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

	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	GetRerservationByID(id int) (models.Reservation, error)
	UpdateReservation(u models.Reservation) error
	DeleteReservation(id int) error
	UpdateProcessedForReservation(id,proccessed int)error
	AllRooms()([]models.Room,error)
	GetRestrictionForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
	InsertBlockForRoom(id int, startDate time.Time) error
	DeleteBlockByID(id int) error
}
