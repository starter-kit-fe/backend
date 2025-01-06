package app

import "admin/internal/handler"

type Handlers struct {
	App         *handler.AppHandler
	User        *handler.UserHandler
	Lookup      *handler.LookupHandler
	Permissions *handler.PermissionsHandler
}

func (a *App) initHandlers() {
	a.handlers = &Handlers{
		App:         handler.NewAppHandler(a.services.App),
		User:        handler.NewUserHandler(a.services.User),
		Lookup:      handler.NewLookupHandler(a.services.Lookup),
		Permissions: handler.NewPermissionsHandler(a.services.Permissions),
	}
}
