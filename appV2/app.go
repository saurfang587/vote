package appV2

import (
	"saurfang/vote/appV2/model"
	"saurfang/vote/appV2/router"
	"saurfang/vote/appV2/tools"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	tools.NewToken("")

	r := router.New()
	_ = r.Run(":8080")
}
