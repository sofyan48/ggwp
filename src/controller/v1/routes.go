package v1

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/ggwp/src/controller/v1/health"
	healthService "github.com/sofyan48/ggwp/src/service/v1/health"

	users "github.com/sofyan48/ggwp/src/controller/v1/user"
	userService "github.com/sofyan48/ggwp/src/service/v1/user"

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

	// Health Handler Routes
	healthHandler := &health.V1HealthController{
		HealthService: healthService.V1HealthCheckHandler(),
	}

	graphQLHandler := &graphQL.GraphQlController{}

	// User Handler Routes
	userHandler := &users.V1UserController{
		UserService: userService.V1UserServiceHandler(),
	}

	//********* Calling Handler To Routers *********//
	rLoader.routerHealthCheck(router, healthHandler)
	rLoader.routerGraphQL(router, graphQLHandler)
	rLoader.routerUsers(router, userHandler)

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
	group := router.Group("v1/users")
	group.POST("", handler.InsertUsers)
	group.PUT(":id", handler.UpdateUsersByID)
}
