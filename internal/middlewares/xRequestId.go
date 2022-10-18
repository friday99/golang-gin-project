package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware(ctx *gin.Context) {
	uuidV4 := uuid.New().String()
	ctx.Header("X-Request-Id", uuidV4)
	ctx.Next()
}
