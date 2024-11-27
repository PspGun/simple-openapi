package main

import "backend/cmd/api/di"

func main() {
	app, err := di.InitializedApp()
	if err != nil {
		return
	}
	err = app.RunApp()
	if err != nil {
		return
	}
}
