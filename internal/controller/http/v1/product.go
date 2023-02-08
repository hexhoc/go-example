package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"hexhoc/go-examples/internal/entity"
	"hexhoc/go-examples/internal/usecase"
	"net/http"
)

type productRoutes struct {
	p usecase.Product
	l *zerolog.Logger
}

func newProductRoutes(handler *gin.RouterGroup, p usecase.Product, l *zerolog.Logger) {
	r := &productRoutes{p, l}
	h := handler.Group("/product")
	{
		h.GET("/findAll", r.findAll)
		h.GET("/findById", r.findById)
	}
}

type productReponse struct {
	Products []entity.Product `json:"products"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *productRoutes) findAll(c *gin.Context) {
	products, err := r.p.GetAllProduct(c.Request.Context())
	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, productReponse{Products: products})
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *productRoutes) findById(c *gin.Context) {
	products, err := r.p.GetAllProduct(c.Request.Context())
	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, productReponse{Products: products})
}
