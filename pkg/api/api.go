package api

import (
	"github.com/gorilla/mux"
	_ "github.com/poncheska/iot-mousetrap/docs"
	"github.com/poncheska/iot-mousetrap/pkg/store/fake"
	hp "github.com/poncheska/iot-mousetrap/pkg/transport/http"
	"github.com/poncheska/iot-mousetrap/pkg/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := hp.Handler{
		Store: fake.NewFakeStore(),
		Logs:  utils.NewStringLogger(),
	}

	log.SetOutput(io.MultiWriter(os.Stdout, h.Logs))

	r := mux.NewRouter()
	r.HandleFunc("/log", h.GetLog).Methods(http.MethodGet)
	r.HandleFunc("/log/clear", h.ClearLog).Methods(http.MethodGet)
	r.HandleFunc("/org/sign-in", h.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/org/sign-up", h.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/mousetraps", h.AuthChecker(h.GetMousetraps)).Methods(http.MethodGet)
	r.HandleFunc("/trigger/{org}/{name}/{status:[01]}", h.Trigger).Methods(http.MethodGet)
	r.PathPrefix("/swagger/").HandlerFunc(httpSwagger.Handler()).Methods(http.MethodGet)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
