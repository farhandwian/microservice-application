package eureka

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
)

/**
Below is the format required by Eureka to register and application instance
{
    "instance": {
        "hostName": "MY_HOSTNAME",
        "app": "org.github.hellosatish.microservicepattern.awesomeproject",
        "vipAddress": "org.github.hellosatish.microservicepattern.awesomeproject",
        "secureVipAddress": "org.github.hellosatish.microservicepattern.awesomeproject"
        "ipAddr": "10.0.0.10",
        "status": "STARTING",
        "port": {"$": "8080", "@enabled": "true"},
        "securePort": {"$": "8443", "@enabled": "true"},
        "healthCheckUrl": "http://WKS-SOF-L011:8080/healthcheck",
        "statusPageUrl": "http://WKS-SOF-L011:8080/status",
        "homePageUrl": "http://WKS-SOF-L011:8080",
        "dataCenterInfo": {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
            "name": "MyOwn"
        },
    }
}
*/

type AppRegistrationBody struct {
	Instance InstanceDetails `json:"instance"`
}

type InstanceDetails struct {
	HostName         string         `json:"hostName"`
	App              string         `json:"app"`
	VipAddress       string         `json:"vipAddress"`
	SecureVipAddress string         `json:"secureVipAddress"`
	IpAddr           string         `json:"ipAddr"`
	Status           string         `json:"status"`
	Port             Port           `json:"port"`
	SecurePort       Port           `json:"securePort"`
	HealthCheckUrl   string         `json:"healthCheckUrl"`
	StatusPageUrl    string         `json:"statusPageUrl"`
	HomePageUrl      string         `json:"homePageUrl"`
	DataCenterInfo   DataCenterInfo `json:"dataCenterInfo"`
	// MetaData         MetaData       `json:"metadata"`
}
type Port struct {
	Port    string `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

type MetaData struct {
	InstanceId string `json:"instanceId"`
}

// This struct shall be responsible for manager to manage registration with Eureka
type EurekaRegistrationManager struct {
}

func (erm EurekaRegistrationManager) RegisterWithSerivceRegistry(serviceRegistryURL string) {
	log.Print("Registering service with status : STARTING")
	body := erm.getBodyForEureka("STARTING")

	MakePostCall(serviceRegistryURL+"cart-service", body, nil)
	log.Print("Waiting for 10 seconds for application to start properly")
	time.Sleep(10 * time.Second)
	log.Print("Updating the status to : UP")
	bodyUP := erm.getBodyForEureka("UP")
	MakePostCall(serviceRegistryURL+"cart-service", bodyUP, nil)
}

func (erm EurekaRegistrationManager) SendHeartBeat(serviceRegistryURL string) {
	// instanceId := "asjdkajsdmaslkdnalsdn"
	log.Println("In SendHeartBeat!")

	// ipAddress, err := ExternalIP()
	hostname, err := os.Hostname()
	if err != nil {
		log.Print("Error while getting hostname which shall be used as APP ID")
	}
	job := func() {
		fmt.Println("sending heartbeat : ", time.Now().UTC())
		// error klo pake ipadrees
		// MakePutCall(serviceRegistryURL+"cart-service/"+ipAddress+":cart-service:"+instanceId, nil, nil)

		MakePutCall(serviceRegistryURL+"cart-service/"+hostname, nil, nil)
	}
	// Run every 25 seconds but not now.
	scheduler.Every(25).Seconds().Run(job)
	runtime.Goexit()

}
func (erm EurekaRegistrationManager) DeRegisterFromServiceRegistry(serviceRegistryURL string) {
	MakePostCall(serviceRegistryURL, nil, nil)
}

func (erm EurekaRegistrationManager) getBodyForEureka(status string) *AppRegistrationBody {
	httpport := "9090"
	hostname, err := os.Hostname()
	if err != nil {
		log.Print("Enable to find hostname form OS, sending appname as host name")
	}

	ipAddress, err := ExternalIP()
	if err != nil {
		log.Print("Enable to find IP address form OS")
	}

	port := Port{httpport, "true"}
	securePort := Port{"9090", "true"}
	dataCenterInfo := DataCenterInfo{"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", "MyOwn"}
	// metaData := MetaData{"cart-service:"}
	homePageUrl := "http://" + ipAddress + ":" + httpport + "/cart"
	statusPageUrl := "http://" + ipAddress + ":" + httpport + "/status"
	healthCheckUrl := "http://" + ipAddress + ":" + httpport + "/healthcheck"

	instance := InstanceDetails{hostname, "cart-service", "cart-service", "cart-service",
		ipAddress, status, port, securePort, healthCheckUrl, statusPageUrl, homePageUrl, dataCenterInfo}

	body := &AppRegistrationBody{instance}
	return body
}
