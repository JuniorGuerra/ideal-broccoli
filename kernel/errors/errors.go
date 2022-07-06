package errors

import (
	"app/kernel/lang"

	"github.com/gin-gonic/gin"
)

type Extra map[string]interface{}

type HttpErrors struct {
	Ctx    gin.Context
	Status int
	Error  Error
}

type Error struct {
	Message lang.Message `json:"message,omitempty"`
	Extra   Extra        `json:"data,omitempty"`
}

func HttpError(httpError HttpErrors) {
	httpError.Ctx.JSON(httpError.Status, httpError.Error)
}
