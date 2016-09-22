package algorithm

import (
	"k8s.io/kubernetes/pkg/api"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
	"github.com/gaocegege/simple-k8s-scheduler-extender/nodeStatus"
	"fmt"
)

// LeastHostedPriority is a priority function that favors nodes with less hosts.
func LeastHostedPriority(args *schedulerapi.ExtenderArgs) schedulerapi.HostPriorityList {
	result := schedulerapi.HostPriorityList{}

	pod := args.Pod
	nodes := args.Nodes
	pods := nodeStatus.GetAllPods()

	for _, node := range nodes.Items {
		ePods := nodeStatus.GetPodsByNodeName(pods, node.Name)
		result = append(result, calculateResourceScore(&pod, &node, ePods))
	}
	return result
}

func calculateResourceScore(pod *api.Pod, node *api.Node, epods []*api.Pod) schedulerapi.HostPriority {

	allocatableMilliCPU := node.Status.Allocatable.Cpu().MilliValue()
	allocatableMemory := node.Status.Allocatable.Memory().Value()
	milliCPURequested := int64(0)
	memoryRequested := int64(0)
	for _, epod := range epods {
		//fmt.Println("podName: ", pod.Name, "  NodeName: ", pod.Spec.NodeName)
		//fmt.Println("Resource: ")
		podRequest := getResourceRequest(epod)
		milliCPURequested += podRequest.milliCPU
		//fmt.Println("  cpu: ", podRequest.milliCPU)
		memoryRequested += podRequest.memory
		//fmt.Println("  memory: ", podRequest.memory)
	}

	canuseMilliCPU := allocatableMilliCPU - milliCPURequested
	canuseMemory := allocatableMemory - memoryRequested
	//fmt.Println("Node: ", node.Name)
	//fmt.Println("allocatable: ")
	//fmt.Println("    cpu: ", allocatableMilliCPU)
	//fmt.Println("    memory: ", allocatableMemory)
	//fmt.Println("canuse: ")
	//fmt.Println("    cpu: ", canuseMilliCPU)
	//fmt.Println("    memory: ", canuseMemory)

	//capacityMilliCPU := node.Status.Allocatable.Cpu().MilliValue()
	//capacityMemory := node.Status.Allocatable.Memory().Value()
	//fmt.Println(node.Name, "cpu: ", canuseMilliCPU, " memory: ", canuseMemory)
	//fmt.Println(node.Name, "score: ", 10 - (calculateScore(canuseMilliCPU, allocatableMilliCPU) +
	//	calculateScore(canuseMemory, allocatableMemory)) / 2)

	return schedulerapi.HostPriority{
		Host: node.Name,
		Score: 10 - ((calculateScore(canuseMilliCPU, allocatableMilliCPU) +
			calculateScore(canuseMemory, allocatableMemory)) / 2),
	}
}

type resourceRequest struct {
	milliCPU int64
	memory   int64
}

func getResourceRequest(pod *api.Pod) resourceRequest {
	result := resourceRequest{}
	for _, container := range pod.Spec.Containers {
		requests := container.Resources.Requests
		result.memory += requests.Memory().Value()
		result.milliCPU += requests.Cpu().MilliValue()
	}
	return result
}

// the unused capacity is calculated on a scale of 0-10
// 0 being the lowest priority and 10 being the highest
func calculateScore(allocatable int64, capacity int64) int {
	if capacity == 0 {
		return 0
	}
	return int(allocatable * 10 / capacity)
}

func GetPodScheduleStatus() {
	podList := nodeStatus.GetAllPods()
	fmt.Println("Number of pod scheduled : ", len(podList.Items))
	var totalScheduledCpu int64 = int64(0)
	var totalScheduleMemory int64 = int64(0)
	for index, pod := range podList.Items {
		fmt.Println("pod", index, " :")
		podRequest := getResourceRequest(&pod)
		totalScheduledCpu += podRequest.milliCPU
		totalScheduleMemory += podRequest.memory
		fmt.Println("  cpu: ", podRequest.milliCPU, "m")
		fmt.Println("  memory: ", podRequest.memory, "i")
		fmt.Println("scheduled to Node ", pod.Spec.NodeName)
	}
	fmt.Println("total scheduled workload:")
	fmt.Println("  cpu: ", totalScheduledCpu, "m")
	fmt.Println("  memory: ", totalScheduleMemory, "i")
}
