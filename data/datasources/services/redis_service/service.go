// Package defines a service to acces a redis database.
package redis_service

// RedisService provides the type to access this server.
type RedisService struct {
	host string
}

// New creates a RedisService from a host.
func New(host string) *RedisService {
	return &RedisService{host}
}
