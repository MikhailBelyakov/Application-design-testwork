package room_availability

import (
	"time"

	"application-design-test-master/internal/domain/types"
)

type RoomAvailability struct {
	HotelUUID        types.UUID
	RoomUUID         types.UUID
	AvailabilityDate time.Time
	Quota            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}
