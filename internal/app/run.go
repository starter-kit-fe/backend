package app

func (a *App) Run(addr string) error {
	return a.router.Run(":" + addr)
}
