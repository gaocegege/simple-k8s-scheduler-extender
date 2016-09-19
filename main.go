package main

import (
	"fmt"

	"github.com/gaocegege/simple-k8s-scheduler-extender/api"
	"github.com/gin-gonic/gin"
)

// PORT is the port to listen.
const PORT = 12345

func main() {
	r := gin.Default()
	r.POST("v1/prioritize", api.PrioritizeHandler)
	r.Run(fmt.Sprintf(":%d", PORT))
}
