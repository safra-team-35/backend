package service

import "github.com/safra-team-35/backend/domain/contract"

// Service holds the domain service repositories
type Service struct {
	db contract.RepoManager
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	QRCodeService(svc *Service) contract.QRCodeService
	UserService(svc *Service) contract.UserService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) QRCodeService(svc *Service) contract.QRCodeService {
	return newQRCodeService(svc)
}

func (s *serviceManager) UserService(svc *Service) contract.UserService {
	return newUserService(svc)
}
