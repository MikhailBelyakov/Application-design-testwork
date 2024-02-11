package order

import (
	"context"

	"application-design-test-master/internal/domain/room_availability"
	"application-design-test-master/internal/domain/types"
	"application-design-test-master/internal/infra/helper"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

var currencyAvailabilityRooms = []room_availability.RoomAvailability{
	{
		HotelUUID:        "48e862b2-bcb8-4a7e-9f70-ccee4f8297e4",
		RoomUUID:         "c2b7755d-4418-4d00-87dd-6f8f9b853549",
		AvailabilityDate: helper.Date(2024, 1, 1),
		Quota:            1,
	},
	{
		HotelUUID:        "48e862b2-bcb8-4a7e-9f70-ccee4f8297e4",
		RoomUUID:         "c2b7755d-4418-4d00-87dd-6f8f9b853549",
		AvailabilityDate: helper.Date(2024, 1, 2),
		Quota:            1,
	},
	{
		HotelUUID:        "48e862b2-bcb8-4a7e-9f70-ccee4f8297e4",
		RoomUUID:         "c2b7755d-4418-4d00-87dd-6f8f9b853549",
		AvailabilityDate: helper.Date(2024, 1, 3),
		Quota:            1,
	},
	{
		HotelUUID:        "48e862b2-bcb8-4a7e-9f70-ccee4f8297e4",
		RoomUUID:         "c2b7755d-4418-4d00-87dd-6f8f9b853549",
		AvailabilityDate: helper.Date(2024, 1, 4),
		Quota:            1,
	},
	{
		HotelUUID:        "48e862b2-bcb8-4a7e-9f70-ccee4f8297e4",
		RoomUUID:         "c2b7755d-4418-4d00-87dd-6f8f9b853549",
		AvailabilityDate: helper.Date(2024, 1, 5),
		Quota:            1,
	},
}

func (r *Repository) GetByHotelAndRoom(_ context.Context, hotelUUID types.UUID, roomUUID types.UUID) []room_availability.RoomAvailability {
	result := make([]room_availability.RoomAvailability, 0)

	for _, currencyAvailabilityRoom := range currencyAvailabilityRooms {
		if hotelUUID == currencyAvailabilityRoom.HotelUUID && roomUUID == currencyAvailabilityRoom.RoomUUID {
			result = append(result, currencyAvailabilityRoom)
		}
	}
	return result
}
