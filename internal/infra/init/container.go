package init

import (
	"github.com/go-chi/chi/v5"

	"application-design-test-master/internal/infra/logger"

	createOrderApi "application-design-test-master/internal/api/create_order"
	createOrderCommand "application-design-test-master/internal/app/order/command/create"
	availabilityRoomInfra "application-design-test-master/internal/infra/availability_room"
	orderInfra "application-design-test-master/internal/infra/order"
)

type DiContainer struct {
	Common struct {
		Logger *logger.Logger
		Router *chi.Mux
	}
	Infra struct {
		OrderRepository            *orderInfra.Repository
		AvailabilityRoomRepository *availabilityRoomInfra.Repository
	}
	App struct {
		OrderCreateCommand *createOrderCommand.Handler
	}
	Api struct {
		OrderCreateApi *createOrderApi.Handler
	}
}

func NewContainer() *DiContainer {
	di := &DiContainer{}
	provideCommon(di)
	provideInfra(di)
	provideApp(di)
	provideApi(di)

	return di
}

func provideCommon(di *DiContainer) {
	di.Common.Router = chi.NewRouter()
	di.Common.Logger = logger.New()
}

func provideInfra(di *DiContainer) {
	provideOrderRepository(di)
	provideAvailabilityRoomRepository(di)
}

func provideApp(di *DiContainer) {
	provideOrderCreateCommand(di)
}
func provideApi(di *DiContainer) {
	provideOrderCreateApi(di)
}

func provideAvailabilityRoomRepository(di *DiContainer) {
	di.Infra.AvailabilityRoomRepository = availabilityRoomInfra.NewRepository()
}

func provideOrderRepository(di *DiContainer) {
	di.Infra.OrderRepository = orderInfra.NewRepository()
}

func provideOrderCreateCommand(di *DiContainer) {
	di.App.OrderCreateCommand = createOrderCommand.NewHandler(
		di.Common.Logger,
		createOrderCommand.NewCommandValidator(),
		di.Infra.AvailabilityRoomRepository,
		di.Infra.OrderRepository,
	)
}

func provideOrderCreateApi(di *DiContainer) {
	di.Api.OrderCreateApi = createOrderApi.New(
		di.Common.Logger,
		createOrderApi.NewResponse(di.Common.Logger),
		di.App.OrderCreateCommand,
	)
}
