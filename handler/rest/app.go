package rest

import (
	"botkiosgamercodm/helpers"
	"botkiosgamercodm/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

var port = ":" + helpers.GetEnv("PORT")

func StartApp() {
	route := gin.Default()

	// db, err := sql.Open("sqlite3", "database.db")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// 	sts := `
	// DROP TABLE IF EXISTS redeem;
	// CREATE TABLE redeem(id INTEGER PRIMARY KEY AUTOINCREMENT, code TEXT, status INT);
	// `
	// 	_, err = db.Exec(sts)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	userssRepository := users.NewRepository()
	usersService := users.NewService(userssRepository)
	usersHandler := NewUser(usersService)

	route.POST("/topup", usersHandler.Redeem)

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
