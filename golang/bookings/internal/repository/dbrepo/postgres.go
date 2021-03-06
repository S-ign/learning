package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/S-ign/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(
	res models.Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `insert into reservations (first_name, last_name, email, phone,
					start_date, end_date, room_id, created_at, updated_at)
					values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions(start_date, end_date, room_id, 
				reservation_id, created_at, updated_at, restriction_id)
				values ($1,$2,$3,$4,$5,$6,$7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)
	if err != nil {
		return err
	}

	return nil
}

// IsReservationAvailable Checks Database for Reservation Availability for
// roomID
func (m *postgresDBRepo) IsReservationAvailable(
	start, end time.Time, roomID int) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `select count(id)
						from room_restrictions 
						where room_id = $1
						$2 < end_date and $3 > start_date;`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// GetAllAvailableRooms returns a list of available rooms
func (m *postgresDBRepo) GetAllAvailableRooms(
	start, end time.Time) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var roomName string
	var rooms []string

	query := `select room_name from rooms
						left join room_restrictions on rooms.id = room_restrictions.room_id
						where $1 < end_date and $2 > start_date;
	`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	for rows.Next() {
		err = rows.Scan(&roomName)
		if err != nil {
			return []string{}, err
		}
		rooms = append(rooms, roomName)
	}

	if err != nil {
		return []string{}, err
	}

	log.Println(rooms)
	return rooms, nil
}
