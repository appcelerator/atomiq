package core

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

const baseURL = "/api/v1"

// Start API server
func (a *Agent) initAPI() {
	log.Infoln("Start API server on port " + conf.apiPort)
	go func() {
		http.HandleFunc(baseURL+"/health", a.agentHealth)
		err := http.ListenAndServe(":"+conf.apiPort, nil)
		if err != nil {
			log.Fatalln("Unable to start server: ", err)
		}
	}()
}

// or HEALTHCHECK Dockerfile instruction
func (a *Agent) agentHealth(resp http.ResponseWriter, req *http.Request) {
	if a.eventStreamReading {
		resp.WriteHeader(200)
	} else {
		log.Errorln("Error: health check failed")
		resp.WriteHeader(400)
	}
}
