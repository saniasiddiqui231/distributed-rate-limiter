package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
)

func New() (*httputil.ReverseProxy, error) {

	target, err := url.Parse(config.AppConfig.BackendURL)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(target), nil
}
