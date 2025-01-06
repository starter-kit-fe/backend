package app

import (
	"admin/internal/service"
)

type Services struct {
	App         service.AppService
	User        service.UserService
	Lookup      service.LookupService
	Permissions service.PermissionsService
}

func (a *App) initServices() {
	a.services = &Services{
		App: service.NewAppService(),
		User: service.NewUserService(
			service.UserServiceConfig{
				UserRepo:     a.repos.User,
				CfClient:     a.turnstile,
				EmailService: a.emailClient,
				GoogleServer: a.googleService,
				TotpClient:   a.totpClient,
				RDB:          a.rdb,
				JWT:          a.jwtClient,
			},
		),
		Lookup:      service.NewLookupService(a.repos.Lookup),
		Permissions: service.NewPermissionsService(a.repos.Permissions, a.repos.Lookup),
	}
}
