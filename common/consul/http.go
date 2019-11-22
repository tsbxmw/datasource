package consul

import (
    "github.com/hashicorp/consul/api"
    consulapi "github.com/hashicorp/consul/api"
    "github.com/sirupsen/logrus"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func (r *ConsulRegister) RegisterHTTP() {
    rand.Seed(time.Now().UTC().UnixNano())

    consulConfig := consulapi.DefaultConfig()
    consulConfig.Address = r.ConsulAddress + ":" + strconv.Itoa(r.ConsulPort)
    consulClient, err := consulapi.NewClient(consulConfig)
    if err != nil {
        logrus.Error("err", err)
        os.Exit(1)
    }

    portStr := strconv.Itoa(r.Port)

    check := consulapi.AgentServiceCheck{
        HTTP:     "http://" + r.Address + ":" + portStr + "/v1/health",
        Interval: "10s",
        Timeout:  "1s",
        Notes:    "Basic health checks",
    }

    asr := api.AgentServiceRegistration{
        ID:      r.Service,
        Name:    r.Service,
        Address: r.Address,
        Port:    r.Port,
        Tags:    r.Tag,
        Check:   &check,
    }
    err = consulClient.Agent().ServiceRegister(&asr)
    if err != nil {
        logrus.Error(err)
        panic(err)
    }

    return
}
