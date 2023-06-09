package appV0

import (
	"saurfang/vote/appV0/model"
	"saurfang/vote/appV0/router"
)

func Start() {
	model.New()
	defer func() {
		model.Close()
	}()

	r := router.New()
	_ = r.Run(":8080")
}
