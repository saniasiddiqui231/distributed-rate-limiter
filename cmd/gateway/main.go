package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/limiter"
	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/metrics"
	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/middleware"
	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/proxy"
)

func main() {

	metrics.Register()

	l, err := limiter.New()
	if err != nil {
		log.Fatal(err)
	}

	p, err := proxy.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", middleware.RateLimiter(l, p))

	fmt.Println("Gateway running on", config.AppConfig.ServerPort)

	log.Fatal(http.ListenAndServe(config.AppConfig.ServerPort, mux))
}
