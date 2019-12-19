package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
)

func RegisterConsul() {
	config := consulapi.DefaultConfig()

	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	port, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}

	register := new(consulapi.AgentServiceRegistration)
	register.ID = "transaction-microservice"
	register.Name = "transaction-microservice"
	register.Address = hostname()
	register.Port = port
	register.Check = new(consulapi.AgentServiceCheck)
	register.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", hostname(), port)
	register.Check.Interval = "5s"
	register.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(register)
}

func port() string {
	p := os.Getenv("PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8080"
	}
	return fmt.Sprintf(":%s", p)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}

func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server in good condition",
	})
}
