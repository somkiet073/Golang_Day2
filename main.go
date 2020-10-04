package main

import (
	"hellogoDay2/api"
	"hellogoDay2/route"
)

func main() {
	r := route.Env{api.Env{}}.Router()
	r.Run(":8080")
}
