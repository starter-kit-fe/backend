package app

import "admin/internal/repository"

type Repositories struct {
	User        repository.UserRepository
	Lookup      repository.LookupRepository
	Permissions repository.PermissionsRepository
}

func (a *App) initRepositories() {
	user := repository.NewUserRepository(a.db)
	lookup := repository.NewLookupRepository(a.db, user)
	permissions := repository.NewPermissionRepository(a.db)

	a.repos = &Repositories{
		User:        user,
		Lookup:      lookup,
		Permissions: permissions,
	}
}
