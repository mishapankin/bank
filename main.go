package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type client struct {
	id      int32
	name    string
	balance int32
}

func main() {
	fmt.Println(os.Getenv("DB_URL"))

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	router := gin.Default()
	router.GET("/client/:id", getClientInfo)
	router.PATCH("/balance/:id", updateClientBalance)
	router.Run(os.Getenv("SERVER_URL"))
}

func getClientInfo(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("Get", id)

	c.IndentedJSON(http.StatusOK, gin.H{"Hello": "World"})
}

func updateClientBalance(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Update", id)

	c.IndentedJSON(http.StatusOK, gin.H{})
}
