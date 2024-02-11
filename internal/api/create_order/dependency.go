package create_order

import (
	"context"
	"net/http"

	createOrderCommand "application-design-test-master/internal/app/order/command/create"
	orderDom "application-design-test-master/internal/domain/order"
)

type CommandHandler interface {
	Handle(ctx context.Context, cmd createOrderCommand.Command) (*orderDom.Order, error)
}

type Logger interface {
	LogInfo(format string, v ...any)
	LogErrorf(format string, v ...any)
}

type Resp interface {
	OK(out http.ResponseWriter, responseData responseOk)
	Error(out http.ResponseWriter, err error)
}
