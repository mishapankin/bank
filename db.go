package main

import (
	"context"
	"log"
)

func getClientInfo(ctx *appContext, id int) (name string, balance int, last int, err error) {
	err = ctx.dbpool.QueryRow(context.Background(),
		"select name, balance, last_operation from clients where id=$1", id).Scan(&name, &balance, &last)
	return
}

func computeOperations(ctx *appContext, id int) error {
	_, balance, prev_last, err := getClientInfo(ctx, id)

	if err != nil {
		log.Printf("Can't get client info")
		return err
	}

	var last int
	rows, _ := ctx.dbpool.Query(context.Background(),
		"select change, id from operations where id>$1 order by id asc", prev_last)
	defer rows.Close()
	for rows.Next() {
		var change int
		if err := rows.Scan(&change, &last); err != nil {
			log.Printf("Error computing balance")
			return err
		}
		if balance+change >= 0 {
			balance += change
		}
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error computing balance")
		return err
	}

	log.Printf("Result: %d %d", balance, last)
	_, err = ctx.dbpool.Exec(context.Background(),
		"update clients set balance=$1, last_operation=$2 where id=$3",
		balance, last, id)
	return err
}
