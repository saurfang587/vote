package appV0

import (
	"saurfang/vote/appV0/model"
	"saurfang/vote/appV0/router"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	r := router.New()
	_ = r.Run(":8080")
}
