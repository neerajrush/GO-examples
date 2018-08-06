package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

/**
* EndpointNodeMap: keeps mapping for Endpoint and NodeIp:Port in IP:Port format
 */

type EndpointNodeMap struct {
	nodeMap   map[string]string // k:EP->v:NodeIpPort (format: IP:Port)
	probesMap *ProbesMap
}

const (
	ProbeRespTimeout time.Duration = 3
	ProbeInterval    time.Duration = 3
)

func NewEpNodeMap() (*EndpointNodeMap, error) {
	epNodeMap := EndpointNodeMap{
		nodeMap:   make(map[string]string),
		probesMap: NewProbesMap(),
	}
	return &epNodeMap, nil
}

func (epNodeMap *EndpointNodeMap) AddEpNodeMap(ep string, nodeIp string) {
	runes := []rune(ep)
	portIndex := strings.Index(ep, ":")
	nodeIpPort := nodeIp + string(runes[portIndex:])
	epNodeMap.nodeMap[ep] = nodeIpPort
	fmt.Println("Added mapping: ", ep, " ==> ", nodeIpPort)
	//epNodeMap.probesMap.AddEndPoint(nodeIpPort)
	epNodeMap.probesMap.AddEndPoint(nodeIp)
}

func (epNodeMap *EndpointNodeMap) DeleteEpNodeMap(ep string) {
	nodeIpPort, ok := epNodeMap.nodeMap[ep]
	if ok {
		delete(epNodeMap.nodeMap, ep)
		fmt.Println("Deleted mapping for EP: ", ep)
		epNodeMap.probesMap.DeleteEndpoint(nodeIpPort)
	}
}

func (epNodeMap *EndpointNodeMap) ListEpNodeMap() {
	fmt.Println("**list:EndpointNodeMap**")
	for k, v := range epNodeMap.nodeMap {
		fmt.Println(k, " ==> ", v)
	}
	fmt.Println("**done**")
	epNodeMap.probesMap.ListEndpoints()
}

func (epNodeMap *EndpointNodeMap) GetEpNodeMap(ep string) (string, bool) {
	nodeIpPort, ok := epNodeMap.nodeMap[ep]
	if ok {
		fmt.Println("Got mapping: ", ep, " ==> ", nodeIpPort)
		epNodeMap.probesMap.mutex.Lock()
		status, ok := epNodeMap.probesMap.status[nodeIpPort]
		epNodeMap.probesMap.mutex.Unlock()
		if !ok || !status {
			return nodeIpPort, false
		}
	}
	return nodeIpPort, ok
}

/**
* ProbesMap: keeps mapping for NodeIp:Port and probe status (bool)
 */

type ProbesMap struct {
	mutex  *sync.Mutex
	probes map[string]string // k:nodeIpPort->v:req(http://request/healthz)
	status map[string]bool   // k:nodeIpPort->v:ProbeResp(bool)
}

func NewProbesMap() *ProbesMap {
	probesMap := ProbesMap{
		mutex:  &sync.Mutex{},
		probes: make(map[string]string),
		status: make(map[string]bool),
	}
	return &probesMap
}

func (probesMap *ProbesMap) AddEndPoint(addr string) {
	aProbe := fmt.Sprintf("http://%s/healthz", addr) //ip:port
	probesMap.mutex.Lock()
	probesMap.probes[addr] = aProbe
	probesMap.status[addr] = false
	probesMap.mutex.Unlock()
}

func (probesMap *ProbesMap) UpdateEndpoint(addr string) {
	aProbe := fmt.Sprintf("http://%s/healthz", addr) //ip:port
	probesMap.mutex.Lock()
	probesMap.probes[addr] = aProbe
	probesMap.status[addr] = false
	probesMap.mutex.Unlock()
}

func (probesMap *ProbesMap) DeleteEndpoint(addr string) {
	probesMap.mutex.Lock()
	delete(probesMap.probes, addr)
	delete(probesMap.status, addr)
	probesMap.mutex.Unlock()
}

func (probesMap *ProbesMap) ListEndpoints() {
	fmt.Println("**list:Endpoints(ProbesStatus)**")
	probesMap.mutex.Lock()
	for k, v := range probesMap.status {
		fmt.Println(k, " ==> ", v)
	}
	probesMap.mutex.Unlock()
	fmt.Println("**done**")
}

func (probesMap *ProbesMap) GetProbeStatus(ep string) (bool, bool) {
	probesMap.mutex.Lock()
	probeStatus, ok := probesMap.status[ep]
	probesMap.mutex.Unlock()
	if ok {
		fmt.Println("Probe status: ", ep, " ==> ", probesMap.status[ep])
	}
	return probeStatus, ok
}

func (probesMap *ProbesMap) doProbes() {
	probesStatusMap := make(map[string]bool)
	theProbes := make(map[string]string)
	for {
		probesMap.mutex.Lock()
		for k, v := range probesMap.probes {
			theProbes[k] = v
		}
		probesMap.mutex.Unlock()
		for k, v := range theProbes {
			probesStatusMap[k] = probeEP(v)
		}
		probesMap.mutex.Lock()
		for k, v := range probesStatusMap {
			probesMap.status[k] = v
		}
		probesMap.mutex.Unlock()
		for k := range theProbes {
			delete(theProbes, k)
		}
		for k := range probesStatusMap {
			delete(probesStatusMap, k)
		}
		time.Sleep(ProbeInterval * time.Second)
	}
}

func probeEP(probe string) bool {
	var c = &http.Client{
		Timeout: time.Second * ProbeRespTimeout,
	}
	rsp, err := c.Head(probe)
	if err == nil && rsp.StatusCode == http.StatusOK {
		fmt.Printf("Probe Status: %s -> %s\n", probe, rsp.Status)
		return true
	} else {
		if rsp != nil {
			fmt.Printf("ERROR: Probe Status: %s -> %s\n", probe, rsp.Status)
		} else {
			fmt.Printf("ERROR: Probe Status(err): %s -> %s\n", probe, string(err.Error()))
		}
	}
	return false
}

var epNodeMap *EndpointNodeMap
var err error

func main() {
	epNodeMap, err = NewEpNodeMap()
	if err != nil {
		fmt.Println("Error ..", string(err.Error()))
		return
	}
	eplist := []string{"10.20.30.01:1001", "10.20.30.02:1002", "10.20.30.03:1003", "10.20.30.04:1004"}
	nodelist := []string{"172.22.122.65:9031", "172.22.122.65:9032", "172.22.122.65:9033", "172.22.122.65:9034"}
	idx := 0
	for _, ep := range eplist {
		epNodeMap.AddEpNodeMap(ep, nodelist[idx])
		idx += 1
	}
	epNodeMap.ListEpNodeMap()
	go epNodeMap.probesMap.doProbes()
	for {
		epNodeMap.probesMap.ListEndpoints()
		epNodeMap.probesMap.mutex.Lock()
		for k, _ := range epNodeMap.probesMap.status {
			epNodeMap.probesMap.status[k] = false
		}
		epNodeMap.probesMap.mutex.Unlock()
		epNodeMap.probesMap.ListEndpoints()
		time.Sleep(3 * time.Second)
	}
}
