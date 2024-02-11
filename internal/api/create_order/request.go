package create_order

import (
	"time"
)

type requestDTO struct {
	HotelUUID    string    `json:"hotel_uuid"`
	RoomUUID     string    `json:"room_uuid"`
	Email        string    `json:"email"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
}
