package util

import (
	"encoding/json"
	"errors"
	"flag"
	"strconv"

	cApi "github.com/hashicorp/consul/api"
	"kittest.com/def"
)

var client *cApi.Client
var serviceCfg cApi.AgentServiceRegistration

var ServiceType = flag.Int("s", 1, "Input Service Type: 1-http 2-rpc others are wrong")
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

func RegistService(t string) {
	serviceCfg = cApi.AgentServiceRegistration{
		ID:      *serviceId,
		Name:    *serviceName,
		Address: *serviceAddress,
		Port:    *ServicePort,
	}
	if t == "http" {

		serviceCfg.Tags = []string{"http"}
		serviceCfg.Check = &cApi.AgentServiceCheck{
			HTTP:     "http://" + *serviceAddress + ":" + strconv.Itoa(*ServicePort) + "/activity/doublegift/health",
			Interval: "3s",
			Timeout:  "1s",
		}
	} else if t == "rpc" {
		serviceCfg.Tags = []string{"rpc"}
		serviceCfg.Check = &cApi.AgentServiceCheck{
			TCP:      *serviceAddress + ":" + strconv.Itoa(*ServicePort),
			Timeout:  "1s",
			Interval: "5s",
		}
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

func Response(err bool, s string, data interface{}) (interface{}, error) {
	r := def.FormatResponse{Err: err, Msg: s, Data: data}
	if err {
		j, _ := json.Marshal(r)
		return data, errors.New(string(j))
	}
	return r, nil
}
