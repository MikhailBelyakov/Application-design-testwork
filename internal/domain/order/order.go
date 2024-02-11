package order

import (
	"time"

	"application-design-test-master/internal/domain/types"
)

type Order struct {
	HotelUUID    types.UUID
	RoomUUID     types.UUID
	UserEmail    string
	CheckInDate  time.Time
	CheckOutDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
