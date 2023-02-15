package repositories

func RepositoryProvider() []interface{} {
	return []interface{}{
		NewUserRepository,
	}
}
