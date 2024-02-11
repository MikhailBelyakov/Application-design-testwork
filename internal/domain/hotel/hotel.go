package hotel

import (
	"time"

	"application-design-test-master/internal/domain/types"
)

type Hotel struct {
	UUID      types.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
