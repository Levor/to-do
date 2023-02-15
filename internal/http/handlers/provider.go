package handlers

func HandlerProvider() []interface{} {
	return []interface{}{
		NewAuthHandler,
	}
}
