package service

import (
	"context"
	"fmt"
	"time"

	"URL-Shortner/internal/repository"
	"URL-Shortner/pkg/base62"
)

type URLService interface {
	ShortenURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortKey string) (string, error)
}

type urlService struct {
	repo  repository.URLRepository
	cache repository.CacheRepository
}

func NewURLService(repo repository.URLRepository, cache repository.CacheRepository) URLService {
	return &urlService{
		repo:  repo,
		cache: cache,
	}
}

func (s *urlService) ShortenURL(ctx context.Context, originalURL string) (string, error) {
	id, err := s.repo.Save(ctx, originalURL)
	if err != nil {
		return "", fmt.Errorf("failed to save url: %w", err)
	}

	shortKey := base62.Encode(id)

	// Cache the mapping asynchronously or synchronously with 24h TTL
	_ = s.cache.Set(ctx, shortKey, originalURL, 24*time.Hour)

	return shortKey, nil
}

func (s *urlService) GetOriginalURL(ctx context.Context, shortKey string) (string, error) {
	// 1. Try fetching from Redis cache first
	cachedURL, err := s.cache.Get(ctx, shortKey)
	if err == nil && cachedURL != "" {
		return cachedURL, nil
	}

	// 2. Cache miss: Decode shortKey back to ID
	id, err := base62.Decode(shortKey)
	if err != nil {
		return "", fmt.Errorf("invalid short key: %w", err)
	}

	// 3. Fetch from PostgreSQL
	originalURL, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("url not found: %w", err)
	}

	// 4. Backfill cache for future hits
	_ = s.cache.Set(ctx, shortKey, originalURL, 24*time.Hour)

	return originalURL, nil
}
