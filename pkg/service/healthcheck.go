package service

import "github.com/mkrs2404/sre-bootcamp/pkg/repo/healthcheck"

// StudentService is a implementation of the StudentService
type HealthCheckService interface {
	GetHealthz() error
}

type healthCheckService struct {
	repo healthcheck.Repository
}

func NewHealthCheckService(hr healthcheck.Repository) HealthCheckService {
	return healthCheckService{hr}
}

// GetHealthz returns health of the dependencies used by the service
func (h healthCheckService) GetHealthz() error {
	return h.repo.DBHealth()
}
