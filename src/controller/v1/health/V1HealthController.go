package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com:sofyan48/ggwp/src/service/v1/health"
)

// V1HealthController types
type V1HealthController struct {
	HealthService service.V1HealthCheckInterface
}

// HealthCheck params
// @contex: gin Context
func (service *V1HealthController) HealthCheck(context *gin.Context) {
	result := service.HealthService.HealthCheck()
	context.JSON(http.StatusOK, result)
	return
}