package handler

import (
	"net/http"

	"github.com/Sergiios/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	var newOpening schemas.Opening

	if err := ctx.BindJSON(&newOpening); err != nil {
		logger.Errorf("error marshal opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := newOpening.CreateValidate(); err != nil {
		logger.Errorf("error validate opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&newOpening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "createOpening", newOpening)
}
