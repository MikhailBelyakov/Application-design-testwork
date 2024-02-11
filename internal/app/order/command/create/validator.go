package create

import (
	"context"

	"github.com/google/uuid"

	errorsDom "application-design-test-master/internal/domain/errors"
)

type CommandValidator struct {
}

func NewCommandValidator() *CommandValidator {
	return &CommandValidator{}
}

func (v CommandValidator) Validate(ctx context.Context, cmd Command) error {
	_, err := uuid.Parse(string(cmd.HotelUUID))
	if err != nil {
		return errorsDom.New(err.Error(), errorsDom.Validation)
	}

	_, err = uuid.Parse(string(cmd.RoomUUID))
	if err != nil {
		return errorsDom.New(err.Error(), errorsDom.Validation)
	}

	if len(cmd.Email) == 0 {
		return errorsDom.New("email не может быть пустым", errorsDom.Validation)
	}
	return nil
}
