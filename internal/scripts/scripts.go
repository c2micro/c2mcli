package scripts

// инициализация движка скриптинга
func Init() error {
	// регистрация api
	registerApi()
	// регистрация builtin скриптов
	if err := registerBuilin(); err != nil {
		return err
	}
	return nil
}
