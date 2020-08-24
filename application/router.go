package application


func (app *Application) router() {
	app.r.HandleFunc("/health", app.HealthHandler)
	// users
	userRouter := app.r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/add", app.h.AddUser).Methods("POST")
	// chats
	chatRouter := app.r.PathPrefix("/chats").Subrouter()
	chatRouter.HandleFunc("/add", app.h.AddChat).Methods("POST")
	chatRouter.HandleFunc("/get", app.h.GetChatsByUser).Methods("POST")
	// messages
	msgRouter := app.r.PathPrefix("/messages").Subrouter()
	msgRouter.HandleFunc("/add", app.h.AddMessage).Methods("POST")
	msgRouter.HandleFunc("/get", app.h.GetMessages).Methods("POST")

}
