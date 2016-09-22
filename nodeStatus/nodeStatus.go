package nodeStatus

import (
	"k8s.io/kubernetes/pkg/api"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

const (
	userName string = "admin"
	password string = "FlbY3CD6mcFUfZvb"
	destinationServer_Caicloud string = "https://sjtu.caicloudapp.com"
	allpod_caicloud string = "api/v1/pods"
	destinationServer_Test string = "http://202.120.40.177:16380"
	destinationServer_Test2 string = "http://192.168.1.163:8080"
)

func setBasicAuthOfCaicloud(r *http.Request) {
	r.SetBasicAuth(userName, password)
}

func InvokeRequest_Caicloud(method string, url string, body []byte) *http.Response {
	client := http.Client{}
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	setBasicAuthOfCaicloud(req)
	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	//fmt.Println(resp.Header)
	//fmt.Println(resp.Status)
	//fmt.Println(resp.StatusCode)
	if err != nil {
		fmt.Print(err)
	}
	return resp
}
func InvokeGetReuqest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return resp
}

func GetAllPods() api.PodList {
	//resp := InvokeRequest_Caicloud("GET", destinationServer_Caicloud + allpod_caicloud, nil)
	resp := InvokeGetReuqest(destinationServer_Test2 + allpod_caicloud)
	var v api.PodList

	if (resp != nil) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err != nil) {
			fmt.Print(err)
			return v
		}
		if err = json.Unmarshal(body, &v); err != nil {
		}
		return v
	}
	return v
}
func GetPodsByNodeName(podList api.PodList, nodeName string) []*api.Pod {
	var pods []*api.Pod
	for _, pod := range podList.Items {
		if pod.Spec.NodeName == nodeName {
			pods = append(pods, &pod)
		}
	}
	return pods
}