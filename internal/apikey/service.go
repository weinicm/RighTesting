package apikey

type Service struct {
	store ApiKeyStore
}

func NewService(store ApiKeyStore) *Service {
	return &Service{
		store: store,
	}
}
