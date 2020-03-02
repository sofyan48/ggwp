package v1

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/ggwp/src/controller/v1/health"
	healthService "github.com/sofyan48/ggwp/src/service/v1/health"

	users "github.com/sofyan48/ggwp/src/controller/v1/user"
	userService "github.com/sofyan48/ggwp/src/service/v1/user"

	producer "github.com/sofyan48/ggwp/src/controller/v1/producer"
	producerService "github.com/sofyan48/ggwp/src/service/v1/producer"

	graphQL "github.com/sofyan48/ggwp/src/controller/v1/gql"

	"github.com/sofyan48/ggwp/src/util/middleware"
)

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
}

// V1Router Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Router(router *gin.Engine) {

	// Healthcheck Handler Routes
	healthHandler := &health.V1HealthController{
		HealthService: healthService.V1HealthCheckHandler(),
	}

	// Graphql Handler Routes
	graphQLHandler := &graphQL.GraphQlController{}

	// User Handler Routes
	userHandler := &users.V1UserController{
		UserService: userService.V1UserServiceHandler(),
	}

	// User Handler Routes
	producerHandler := &producer.V1ProducerController{
		Producer: producerService.ProducerServiceHAndler(),
	}

	//********* Calling Handler To Routers *********//
	rLoader.routerHealthCheck(router, healthHandler)
	rLoader.routerGraphQL(router, graphQLHandler)
	rLoader.routerUsers(router, userHandler)
	rLoader.routerProducer(router, producerHandler)

}

//********* Routing API *********//

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerHealthCheck(router *gin.Engine, handler *health.V1HealthController) {
	group := router.Group("v1/healthcheck")
	group.GET("", handler.HealthCheck)
}

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerGraphQL(router *gin.Engine, handler *graphQL.GraphQlController) {
	group := router.Group("v1/graphql")
	group.POST("", handler.GraphQLControll)
}

func (rLoader *V1RouterLoader) routerUsers(router *gin.Engine, handler *users.V1UserController) {
	group := router.Group("v1/user")
	group.POST("", handler.InsertUsers)
	group.PUT(":id", handler.UpdateUsersByID)
}

func (rLoader *V1RouterLoader) routerProducer(router *gin.Engine, handler *producer.V1ProducerController) {
	group := router.Group("v1/producer")
	group.POST("", handler.ProducerSend)
}
