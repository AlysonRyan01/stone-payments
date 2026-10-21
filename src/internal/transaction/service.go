package transaction

import (
	"context"
	"log"
	"strings"
	"time"
)

type Service struct {
	repo *Repository
	velocity *RedisVelocity
}

func NewService(r *Repository, v *RedisVelocity) *Service {
	return &Service{
		repo: r,
		velocity: v,
	}
}

const highAmountThreshold int64 = 1_000_000

func (s *Service) Create(ctx context.Context, t *Transaction) (string, error) {
	notStandard, err := s.velocity.IsNotStandard(ctx, t.CardNumber, time.Now())
	if err != nil {
		log.Printf("velocity check failed: %v", err)
	}

	var reasons []string
	if t.Amount > highAmountThreshold {
		reasons = append(reasons, "high amount")
	}
	if notStandard {
		reasons = append(reasons, "not standard")
	}

	if len(reasons) > 0 {
		return s.repo.CreateSuspiciousTransaction(ctx, t, strings.Join(reasons, ", "))
	}
	return s.repo.Create(ctx, t)
}