package util

import (
	"flag"
	"strconv"

	cApi "github.com/hashicorp/consul/api"
)

var client *cApi.Client
var serviceCfg cApi.AgentServiceRegistration

var serviceId = flag.String("id", "", "Input Service Id")
var serviceName = flag.String("name", "doubleGiftService", "Input Service Name")
var serviceAddress = flag.String("address", "127.0.0.1", "Input Service Address")
var ServicePort = flag.Int("port", 8080, "Input Service Port")

func init() {
	cfg := cApi.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	c, err := cApi.NewClient(cfg)
	client = c
	if err != nil {
		panic(err)
	}
	flag.Parse()
}

func RegistService() {

	serviceCfg = cApi.AgentServiceRegistration{
		ID:      *serviceId,
		Name:    *serviceName,
		Address: *serviceAddress,
		Port:    *ServicePort,
		Tags:    []string{"default"},
		Check: &cApi.AgentServiceCheck{
			HTTP:     "http://" + *serviceAddress + ":" + strconv.Itoa(*ServicePort) + "/activity/doublegift/health",
			Interval: "3s",
		},
	}

	if serviceCfg.ID == "" {
		panic("service id not input")
	}
	err := client.Agent().ServiceRegister(&serviceCfg)
	if err != nil {
		panic(err)
	}
}

func DeregistService() {
	client.Agent().ServiceDeregister(*serviceId)
}
