package router

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include adds the options to the routers.
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init initializes the routers.
func Init(r *gin.Engine) {
	for _, opt := range options {
		opt(r)
	}
}
