package internal

import (
	"log"

	"github.com/Aritra640/docker-compose-eg/Database/db"
	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
  ctx := c.Request().Context()
  var input_user db.AddUserParams 
  err := c.Bind(&input_user); if err != nil {
    log.Println("Error in getting user data")
    return c.JSON(404 , "could not get user data")
  }
  
  _,err = App.QueryObj.AddUser(ctx , input_user)
  if err != nil {
    log.Println("Error ! could not handle query add user")
    return c.JSON(404 , "could not add user")
  }

  return c.JSON(200 , "user added")
}


func FindUser(c echo.Context) error {
  ctx := c.Request().Context()
  var Username string 
  err := c.Bind(&Username); if err != nil {
    log.Println("Error : could not get username")
    panic(err)
  }

  _,err = App.QueryObj.GetUserByUsername(ctx , Username); if err != nil {
    log.Println("error cannot get username")
    return c.JSON(404 , "could not find user")
  }

  return c.JSON(200 , "user found")
}
