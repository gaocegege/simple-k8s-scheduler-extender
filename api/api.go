package api

import (
	"fmt"

	"github.com/simple-k8s-scheduler-extender/algorithm"
	"github.com/gin-gonic/gin"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
)

// PrioritizeHandler is the prioritize handler.
func PrioritizeHandler(c *gin.Context) {
	var args schedulerapi.ExtenderArgs
	c.BindJSON(&args)
	fmt.Println("args: ", args)
	fmt.Println(algorithm.LeastHostedPriority(&args))
	c.JSON(200, algorithm.LeastHostedPriority(&args))
}
