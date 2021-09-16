package main

import (
	"Test-Mogodb/ipi-mogodb/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	gr := g.Group("v1/ipi")
	gr.POST("/create", gin.HandlerFunc(handlers.CreateUser))
	gr.PUT("/update", gin.HandlerFunc(handlers.UpdateUser))
	gr.GET("/Find/:id", gin.HandlerFunc(handlers.FindUser))
	gr.GET("/List", gin.HandlerFunc(handlers.LisAllUser))
	gr.GET("/ListName/:name", gin.HandlerFunc(handlers.LisName))
	gr.PUT("/delete/:id", gin.HandlerFunc(handlers.DeleteUser))
	http.ListenAndServe(":8080", g)

}
