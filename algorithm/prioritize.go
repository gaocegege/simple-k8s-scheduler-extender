package algorithm

import (
	"k8s.io/kubernetes/pkg/api"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
)

// LeastHostedPriority is a priority function that favors nodes with less hosts.
func LeastHostedPriority(args *schedulerapi.ExtenderArgs) schedulerapi.HostPriorityList {
	result := schedulerapi.HostPriorityList{}

	pod := args.Pod
	nodes := args.Nodes

	for _, node := range nodes.Items {
		result = append(result, calculateResourceScore(&pod, &node))
	}
	return result
}

func calculateResourceScore(pod *api.Pod, node *api.Node) schedulerapi.HostPriority {
	allocatableMilliCPU := node.Status.Allocatable.Cpu().MilliValue()
	allocatableMemory := node.Status.Allocatable.Memory().Value()
	capacityMilliCPU := node.Status.Allocatable.Cpu().MilliValue()
	capacityMemory := node.Status.Allocatable.Memory().Value()

	return schedulerapi.HostPriority{
		Host: node.Name,
		Score: (calculateScore(allocatableMilliCPU, capacityMilliCPU) +
			calculateScore(allocatableMemory, capacityMemory)) / 2,
	}
}

// the unused capacity is calculated on a scale of 0-10
// 0 being the lowest priority and 10 being the highest
func calculateScore(allocatable int64, capacity int64) int {
	if capacity == 0 {
		return 0
	}
	return int(allocatable * 10 / capacity)
}
