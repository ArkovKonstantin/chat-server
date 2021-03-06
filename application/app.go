package application

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"chat-server/handlers"
)

type Application struct {
	servicePort int
	r           *mux.Router
	h           handlers.Handler
}

func New(handler handlers.Handler) Application {
	return Application{servicePort: 9000, r: mux.NewRouter(), h: handler}
}

func (app *Application) Start() {
	app.router()
	fmt.Println("start listening on :9000 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, app.servicePort), app.r))
}

func (app *Application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


