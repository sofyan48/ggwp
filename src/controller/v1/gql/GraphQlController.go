package gql

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gql "github.com/graphql-go/graphql"
	schema "github.com/sofyan48/ggwp/src/schema/v1"
	"github.com/sofyan48/ggwp/src/util/helper/rest"
)

// GraphQlController types
type GraphQlController struct{}

// QueryInterface ...
var QueryInterface map[string]interface{}

// GraphQLControll params
// @contex: gin Context
func (controller *GraphQlController) GraphQLControll(context *gin.Context) {
	context.BindJSON(&QueryInterface)
	querySchema, _ := schema.Schema()
	query, isOk := QueryInterface["query"]
	if !isOk {
		rest.ResponseMessages(context, http.StatusBadRequest, "Invlaid Query")
		return
	}
	result := gql.Do(gql.Params{
		Schema:        querySchema,
		RequestString: query.(string),
	})
	context.JSON(http.StatusOK, result)
}
