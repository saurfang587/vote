package appV2

import (
	"saurfang/vote/appV2/model"
	"saurfang/vote/appV2/router"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	r := router.New()
	_ = r.Run(":8080")
}
