package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type appContext struct {
	dbpool *pgxpool.Pool
}

func main() {
	ctx := appContext{}
	var err error

	ctx.dbpool, err = pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer ctx.dbpool.Close()

	router := gin.Default()
	router.GET("/client/:id", wrappContext(&ctx, routeGetClientInfo))
	router.POST("/operation/:id", wrappContext(&ctx, routeCreateOperation))
	router.POST("/client", wrappContext(&ctx, routeCreateNewClient))

	err = router.Run()
	if err != nil {
		log.Fatalf("Unable to start http server: %v\n", err)
	}
}
