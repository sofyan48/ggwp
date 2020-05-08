package health

import (
	"time"

	"github.com/jinzhu/gorm"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	connection "github.com/sofyan48/ggwp/src/util/helper/mysqlconnection"
	redisConn "github.com/sofyan48/ggwp/src/util/helper/redis"
)

// V1HealthCheck | Derivated from UserRepository
type V1HealthCheck struct {
	DB gorm.DB
}

// V1HealthCheckHandler Handler
func V1HealthCheckHandler() *V1HealthCheck {
	return &V1HealthCheck{
		DB: *connection.GetConnection(),
	}
}

//V1HealthCheckInterface declare All Method
type V1HealthCheckInterface interface {
	HealthCheck() *httpEntity.HealthResponse
}

// HealthCheck Function
// return HealthResponse
func (service *V1HealthCheck) HealthCheck() *httpEntity.HealthResponse {
	result := &httpEntity.HealthResponse{}

	err := service.DB.DB().Ping()
	if err != nil {
		result.Mysql = "Fail : Check Your Connection"
	} else {
		result.Mysql = "OK"
	}
	_, errd := redisConn.InitRedis()
	if errd != nil {
		result.Redis = "Fail : Check Your Connection"
	} else {
		result.Redis = "OK"
	}

	result.Status = "Servers OK"

	nows := time.Now()
	result.CreatedAt = &nows
	return result
}
