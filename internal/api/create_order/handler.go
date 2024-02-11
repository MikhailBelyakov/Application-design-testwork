package create_order

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	createCommand "application-design-test-master/internal/app/order/command/create"
	"application-design-test-master/internal/domain/types"
	"application-design-test-master/internal/infra/logger"
)

type Handler struct {
	log            Logger
	commandHandler CommandHandler
	resp           Resp
}

func New(
	log *logger.Logger,
	resp Resp,
	commandHandler CommandHandler,
) *Handler {
	return &Handler{
		log:            log,
		resp:           resp,
		commandHandler: commandHandler,
	}
}

func (h *Handler) Handle(out http.ResponseWriter, in *http.Request) {
	var inDTO requestDTO
	ctx := context.Background()

	err := json.NewDecoder(in.Body).Decode(&inDTO)
	if err != nil {
		h.resp.Error(out, err)
		return
	}

	newEntity, err := h.commandHandler.Handle(ctx, createCommand.Command{
		HotelUUID:    types.UUID(inDTO.HotelUUID),
		RoomUUID:     types.UUID(inDTO.RoomUUID),
		Email:        inDTO.Email,
		CheckInDate:  inDTO.CheckInDate,
		CheckOutDate: inDTO.CheckOutDate,
	})
	if err != nil {
		h.resp.Error(out, err)
		return
	}

	h.resp.OK(out, responseOk{
		Result: true,
		Order: OrderItem{
			HotelUUID:    string(newEntity.HotelUUID),
			RoomUUID:     string(newEntity.RoomUUID),
			Email:        newEntity.UserEmail,
			CheckInDate:  newEntity.CheckInDate.Format(time.DateTime),
			CheckOutDate: newEntity.CheckOutDate.Format(time.DateTime),
		},
	})
	return
}
