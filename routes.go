package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func routeGetClientInfo(c *gin.Context, ctx *appContext) {
	id_s := c.Param("id")

	id, err := strconv.Atoi(id_s)
	if err != nil {
		returnRequestError(c, "id is not an integer", err)
		return
	}

	if err := computeOperations(ctx, id); err != nil {
		returnInternalError(c, "Error updating balance", err)
		return
	}

	name, balance, _, err := getClientInfo(ctx, id)

	if err != nil {
		returnRequestError(c, "client doesn't exist", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"name": name, "balance": balance, "id": id})
}

func routeCreateNewClient(c *gin.Context, ctx *appContext) {
	log.Printf("Create new client")

	name := c.Query("name")
	if name == "" {
		returnRequestError(c, "Empty client name", nil)
		return
	}

	var id, balance int
	err := ctx.dbpool.QueryRow(context.Background(),
		"insert into clients(name) values($1) returning id, balance", name).Scan(&id, &balance)

	if err != nil {
		returnInternalError(c, "can't create a new client", err)
		return
	}

	log.Printf("Added client with id=%d; name=%s", id, name)
	c.IndentedJSON(http.StatusOK, gin.H{"id": id, "balance": balance, name: "name"})
}

// Не используется. В последней версии баланс считается лениво, все изменения баланса сохраняются в базу данных.
func updateClientBalance(c *gin.Context, ctx *appContext) {
	log.Printf("Update client balance")

	id := c.Param("id")
	id_num, err := strconv.Atoi(id)
	if err != nil {
		returnRequestError(c, "id is not valid", err)
		return
	}

	change := c.Query("change")
	change_num, err := strconv.Atoi(change)
	if err != nil {
		returnRequestError(c, "change value is not valid", err)
		return
	}

	log.Printf("Updating balance id=%d change=%d", id_num, change_num)

	tx, err := ctx.dbpool.Begin(context.Background())
	if err != nil {
		returnInternalError(c, "Can't do a transaction. Transaction added to waiting list", err)
		return
	}
	defer tx.Rollback(context.Background())

	var balance int
	err = ctx.dbpool.QueryRow(context.Background(),
		"select balance from clients where id=$1", id_num).Scan(&balance)
	if err != nil {
		returnInternalError(c, "Can't get client balance. Transaction added to waiting list", err)
		return
	}

	if balance+change_num < 0 {
		returnRequestError(c, "Not enough money on balance. Reverting transaction", err)
		return
	}

	var new_balance int
	err = ctx.dbpool.QueryRow(context.Background(),
		"update clients set balance=balance+$2 where id=$1 returning balance",
		id_num, change_num).Scan(&new_balance)
	if err != nil {
		returnInternalError(c, "Can't perform operation. Transaction added to waiting list", err)
		return
	}
	err = tx.Commit(context.Background())
	if err != nil {
		returnInternalError(c, "Error while perfoming transactoin. Transaction added to waiting list", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id_num, "balance": new_balance})
}

func routeCreateOperation(c *gin.Context, ctx *appContext) {
	log.Printf("Create operation")

	id_s := c.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		returnRequestError(c, "id is not valid", err)
		return
	}

	change_s := c.Query("change")
	change, err := strconv.Atoi(change_s)
	if err != nil {
		returnRequestError(c, "change value is not valid", err)
		return
	}

	log.Printf("Adding operation with id=%d change=%d", id, change)

	var op_id int
	err = ctx.dbpool.QueryRow(context.Background(),
		"insert into operations(client_id, change) values($1, $2) returning id", id, change).Scan(&op_id)

	if err != nil {
		returnInternalError(c, "can't add operation", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id, "change": change_s, "op_id": op_id})
	log.Printf("Added operation with id=%d change=%d op_id=%d", id, change, op_id)
}
