package create

import (
	"context"
	"time"

	"github.com/pkg/errors"

	errorsDom "application-design-test-master/internal/domain/errors"
	orderDom "application-design-test-master/internal/domain/order"
	"application-design-test-master/internal/domain/types"
	"application-design-test-master/internal/infra/helper"
	"application-design-test-master/internal/infra/logger"
)

type Command struct {
	HotelUUID    types.UUID
	RoomUUID     types.UUID
	Email        string
	CheckInDate  time.Time
	CheckOutDate time.Time
}

type Handler struct {
	log                        Logger
	validator                  Validator
	availabilityRoomRepository AvailabilityRoomRepository
	orderRepository            OrderRepository
}

func NewHandler(
	log *logger.Logger,
	validator Validator,
	availabilityRoomRepository AvailabilityRoomRepository,
	orderRepository OrderRepository,
) *Handler {
	return &Handler{
		log:                        log,
		validator:                  validator,
		availabilityRoomRepository: availabilityRoomRepository,
		orderRepository:            orderRepository,
	}
}

func (h Handler) Handle(ctx context.Context, cmd Command) (*orderDom.Order, error) {
	if err := h.validator.Validate(ctx, cmd); err != nil {
		return nil, errors.WithStack(err)
	}

	newOrder := orderDom.Order{
		HotelUUID:    cmd.HotelUUID,
		RoomUUID:     cmd.RoomUUID,
		UserEmail:    cmd.Email,
		CheckInDate:  cmd.CheckInDate,
		CheckOutDate: cmd.CheckOutDate,
	}

	daysToBook := helper.DaysBetween(cmd.CheckInDate, cmd.CheckOutDate)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	availabilityRooms := h.availabilityRoomRepository.GetByHotelAndRoom(ctx, cmd.HotelUUID, cmd.RoomUUID)
	for _, dayToBook := range daysToBook {
		for i, availability := range availabilityRooms {
			if !availability.AvailabilityDate.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			availabilityRooms[i] = availability

			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		h.log.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return nil, errorsDom.New("Hotel room is not available for selected dates")
	}

	saveStatus := h.orderRepository.Create(ctx, newOrder)
	if saveStatus == false {
		return nil, errorsDom.New("New orderDom save error")
	}

	return &newOrder, nil
}
