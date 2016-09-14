package algorithm

import (
	"fmt"

	"github.com/gin-gonic/gin"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
)

func PrioritizeHandler(c *gin.Context) {
	var args schedulerapi.ExtenderArgs
	c.BindJSON(&args)
	fmt.Println(args)
}
