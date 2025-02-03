package app

func (a *App) Run(addr string) error {
	a.Setup()
	return a.router.Run(":" + addr)
}
