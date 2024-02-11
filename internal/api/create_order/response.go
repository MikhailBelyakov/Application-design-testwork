package create_order

import (
	"encoding/json"
	"errors"
	"net/http"

	errorsDom "application-design-test-master/internal/domain/errors"
	infra "application-design-test-master/internal/infra/logger"
)

type Response struct {
	log Logger
}

func NewResponse(
	log *infra.Logger,
) *Response {
	return &Response{
		log: log,
	}
}

type responseOk struct {
	Result bool      `json:"result"`
	Order  OrderItem `json:"order"`
}

type OrderItem struct {
	HotelUUID    string `json:"hotelUUID"`
	RoomUUID     string `json:"roomUUID"`
	Email        string `json:"email"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
}

func (r Response) OK(out http.ResponseWriter, responseData responseOk) {
	out.Header().Set("Content-Type", "application/json")
	out.WriteHeader(http.StatusCreated)
	marshal, err := json.Marshal(responseData)
	if err != nil {
		r.log.LogInfo(err.Error())
		return
	}

	_, err = out.Write(marshal)
	if err != nil {
		r.log.LogInfo(err.Error())
		return
	}
	return
}

type responseError struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func (r Response) Error(out http.ResponseWriter, err error) {
	out.Header().Set("Content-Type", "application/json")
	marshal, err := json.Marshal(responseError{
		Result:  false,
		Message: err.Error(),
	})
	if err != nil {
		r.log.LogInfo(err.Error())
		return
	}

	var v errorsDom.Error
	if errors.As(err, &v) {
		code := http.StatusInternalServerError
		switch v.Kind {
		case errorsDom.KindInternal:
			code = http.StatusInternalServerError
		case errorsDom.KindNotFound:
			code = http.StatusNotFound
		case errorsDom.KindValidation:
			code = http.StatusUnprocessableEntity
		case errorsDom.KindConflict:
			code = http.StatusConflict
		case errorsDom.KindForbidden:
			code = http.StatusForbidden
		}
		out.WriteHeader(code)
	} else {
		out.WriteHeader(http.StatusBadRequest)
	}

	_, err = out.Write(marshal)
	if err != nil {
		r.log.LogInfo(err.Error())
		return
	}
}
