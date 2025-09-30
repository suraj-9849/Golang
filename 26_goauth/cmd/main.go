package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suraj-9849/Golang/internal/http/handlers"
	"github.com/suraj-9849/Golang/internal/http/routes"
	"github.com/suraj-9849/Golang/prisma/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	handler := handlers.NewHandler(client)
	r := gin.Default()
	routes.SetupRoutes(r, handler)
	r.Run()
}
