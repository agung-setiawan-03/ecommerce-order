package main

import (
	"ecommerce-order/cmd"
	"ecommerce-order/helpers"
)

func main() {
	// Load config
	helpers.SetupConfig()

	// Load log 
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// load redis 
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	// run http
	cmd.ServeHTTP()
}