package v1

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/ggwp/src/controller/v1/health"
	healthService "github.com/sofyan48/ggwp/src/service/v1/health"

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

	//********* Calling Handler To Routers *********//
	rLoader.routerHealthCheck(router, healthHandler)
	rLoader.routerGraphQL(router, graphQLHandler)

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
