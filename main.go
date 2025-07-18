package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Link struct{
	Id string
	Url string
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var linkMap = map[string]*Link{"example": { Id: "example", Url: "https://example.com"}}

func main() {
	e:=echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.GET("/:id", RedirectHandler)
	// e.GET("/", IndexHandler)
	// e.POST("/submit", SubmitHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func RedirectHandler(c echo.Context) error {
	id:= c.Param("id")
	link, found := linkMap[id]

	if !found{
		return c.String(http.StatusNotFound, "Link not found")
	}

	return c.Redirect(http.StatusMovedPermanently, link.Url)
}

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	
	for i:=0; i< length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}