package create

import (
	"context"

	orderDom "application-design-test-master/internal/domain/order"
	roomAvailabilityDom "application-design-test-master/internal/domain/room_availability"
	"application-design-test-master/internal/domain/types"
)

type AvailabilityRoomRepository interface {
	GetByHotelAndRoom(ctx context.Context, hotelUUID types.UUID, roomUUID types.UUID) []roomAvailabilityDom.RoomAvailability
}

type OrderRepository interface {
	GetAll(ctx context.Context) []orderDom.Order
	Create(context.Context, orderDom.Order) bool
}

type Validator interface {
	Validate(ctx context.Context, command Command) error
}

type Logger interface {
	LogErrorf(format string, v ...any)
}
