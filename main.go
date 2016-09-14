package main

import (
	"fmt"

	"github.com/gaocegege/simple-k8s-scheduler-extender/algorithm"
	"github.com/gin-gonic/gin"
)

// PORT is the port to listen.
const PORT = 12345

func main() {
	r := gin.Default()
	r.POST("/prioritize", algorithm.PrioritizeHandler)
	// Listen and server on 0.0.0.0:12345.
	r.Run(fmt.Sprintf(":%d", PORT))
}
