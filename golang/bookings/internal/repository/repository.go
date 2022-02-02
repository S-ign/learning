package repository

import (
	"time"

	"github.com/S-ign/bookings/internal/models"
)

// DatabaseRepo crud and other usful methods
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	IsReservationAvailable(start, end time.Time, roomID int) (bool, error)
	GetAllAvailableRooms(start, end time.Time) ([]string, error)
}
