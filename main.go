package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Aritra640/docker-compose-eg/Database/db"
	"github.com/Aritra640/docker-compose-eg/internal"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {

  err := godotenv.Load(); if err != nil {
    log.Println("Error could not load .env file")
    panic(err)
  }

  postgresURL := os.Getenv("DATABASE_URL")
  if postgresURL == "" {
    log.Println("Database url is found empty!")
    panic("Database url cannot be empty!")
  }
  
  pg,err := sql.Open("postgres" , postgresURL)
  if err != nil {
    panic(err)
  }
  defer pg.Close()

  //Run migrations
  if err := goose.Up(pg , "./Database/migrations"); err != nil {
    log.Println("migrations failed!")
    panic(err)
  }
  log.Println("Migrations successfull!")
  internal.App = new(internal.Config)
  internal.App.DB = pg
  internal.App.QueryObj = db.New(pg)
  ctx := context.Background()
  internal.App.CTX = ctx

  e := echo.New()
  
  e.GET("/" , func(c echo.Context) error {
    return c.JSON(200 , "server active!")
  })
  e.POST("/add_user" , internal.AddUser)
  e.GET("/find_user" , internal.FindUser)

  e.GET("/" , func(c echo.Context) error {
    return c.JSON(200 , "server working")
  })

  e.Logger.Fatal(e.Start(":8080"))
}


