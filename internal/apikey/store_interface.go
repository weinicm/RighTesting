package apikey

import "time"

type ApiKey struct {
	ID         int
	Key        string
	ProviderID int
	CreatedAt  time.Time
	LastUsedAt *time.Time
	ExpiresAt  *time.Time
}

// ApiKeyStore defines the interface for managing API keys in the storage.
type ApiKeyStore interface {
	// Create creates a new API key.
	Create(key *ApiKey) error

	// Get retrieves the API key object based on the given key string.
	Get(key string) (*ApiKey, error)

	// Update updates an existing API key.
	Update(key *ApiKey) error

	// Delete deletes the specified API key.
	Delete(key string) error
}
