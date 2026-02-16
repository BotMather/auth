package utils

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Client struct {
	requests int
	window   time.Time
	lastSeen time.Time
}

type RateLimiter struct {
	clients map[string]*Client
	limit   int
	logger  *zap.Logger
	mu      sync.Mutex
}

func NewRateLimiter(ctx context.Context, logger *zap.Logger, limit int) *RateLimiter {
	rl := &RateLimiter{
		clients: make(map[string]*Client),
		limit:   limit,
		logger:  logger,
	}

	// Cleanup goroutine
	go rl.cleanup(ctx)

	return rl
}

func (rl *RateLimiter) cleanup(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	rl.logger.Info("Rate limitter cleanup started")
	for {
		select {
		case <-ctx.Done():
			rl.logger.Info("Rate limiter cleanup goroutine stopped")
			return
		case <-ticker.C:
			rl.mu.Lock()
			for ip, client := range rl.clients {
				// 2 minutdan beri ishlatilmagan IP o‘chadi
				if time.Since(client.lastSeen) > 2*time.Minute {
					delete(rl.clients, ip)
				}
			}
			rl.mu.Unlock()
		}
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	client, exists := rl.clients[ip]
	if !exists {
		rl.clients[ip] = &Client{
			requests: 1,
			window:   now,
			lastSeen: now,
		}
		return true
	}

	// yangi minut boshlangan bo‘lsa reset
	if now.Sub(client.window) > time.Minute {
		client.requests = 0
		client.window = now
	}

	if client.requests >= rl.limit {
		return false
	}

	client.requests++
	client.lastSeen = now
	return true
}
