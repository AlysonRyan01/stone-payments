package transaction

import "context"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, t *Transaction) (string, error) {
	return s.repo.Create(ctx, t)
}