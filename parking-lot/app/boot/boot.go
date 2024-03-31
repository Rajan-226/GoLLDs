package boot

import (
	"main/app/api"
	"main/app/models/slot"
	"main/app/modules/admin"
	"main/app/modules/display"
)

var (
	slotRepo slot.IRepo
)

var (
	adminProcessor   *admin.Processor
	displayProcessor *display.Processor
)

func Init() *api.Server {
	initRepos()
	initProcessor()

	return api.NewServer(adminProcessor)
}

func initRepos() {
	slotRepo = slot.NewRepo()
}

func initProcessor() {
	adminProcessor = admin.NewProcessor(slotRepo)

	displayProcessor = display.NewProcessor()
}
