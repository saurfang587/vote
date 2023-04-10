package appV1

import (
	"saurfang/vote/appV1/model"
	"saurfang/vote/appV1/router"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	r := router.New()
	_ = r.Run(":8080")
}
