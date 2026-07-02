package limiter

type Limiter interface {
	Allow(clientID string) (bool, error)
}
