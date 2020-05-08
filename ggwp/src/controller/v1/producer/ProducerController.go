package producer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	services "github.com/sofyan48/ggwp/src/service/v1/producer"
	"github.com/sofyan48/ggwp/src/util/helper/rest"
)

// V1ProducerController ...
type V1ProducerController struct {
	Producer services.ProducerServiceInterface
}

// ProducerSend ...
func (controller *V1ProducerController) ProducerSend(context *gin.Context) {
	payload := &httpEntity.ProducerRequest{}
	if err := context.ShouldBindJSON(payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := controller.Producer.SendMessages(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
}
