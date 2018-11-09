package main

import (
	"net/http"
	"time"

	"github.com/L-oris/yabb/mywire"

	"github.com/L-oris/yabb/foreign/env"
	"github.com/L-oris/yabb/inject"
	"github.com/L-oris/yabb/logger"
)

func main() {
	container := inject.CreateContainer()
	// router := container.Get(types.Router.String()).(http.Handler)
	router, err := mywire.ProvideRouter()
	// event, err := inject.InitializeEvent("custom message")
	// if err != nil {
	// 	fmt.Printf("failed to create event: %s\n", err)
	// 	os.Exit(2)
	// }
	server := &http.Server{
		Addr:         ":" + env.Vars.Port,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	logger.Log.Infof("server listening on port %s", env.Vars.Port)
	logger.Log.Fatal(server.ListenAndServe())
}

// TODO: move away post actions from postRepository (create separate engine)
