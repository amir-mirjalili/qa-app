package main

import (
	"net/http"
	"qa-app/controller"
	mysql "qa-app/pkg/gorm"
	"qa-app/repository"
	"qa-app/router"
	"qa-app/service"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//Database
	db := mysql.ConnectionDB()
	//Init Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	//Init Service
	userService := service.NewUserService(userRepository)

	//Init Controller
	userController := controller.NewUserController(userService)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(userController),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
