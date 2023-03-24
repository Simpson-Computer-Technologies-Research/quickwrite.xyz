package main

import "github.com/realTristan/QuickWrite/server/api"

// Main Function
func main() {
	api.ListenAndServe(":8000")
}
