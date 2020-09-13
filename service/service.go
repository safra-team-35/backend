package service

import (
	"net"
	"net/http"

	"github.com/safra-team-35/backend/domain"
	"github.com/safra-team-35/backend/domain/contract"
)

// Service holds the domain service repositories
type Service struct {
	db         contract.RepoManager
	httpClient *http.Client
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	httpDialer := new(net.Dialer)
	httpDialer.Timeout = domain.DefaultConnectionTimeout

	httpTransport := new(http.Transport)
	httpTransport.TLSHandshakeTimeout = domain.DefaultConnectionTimeout
	httpTransport.Dial = httpDialer.Dial

	svc.httpClient = new(http.Client)
	svc.httpClient.Transport = httpTransport
	svc.httpClient.Timeout = domain.DefaultConnectionTimeout

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	QRCodeService(svc *Service) contract.QRCodeService
	UserService(svc *Service) contract.UserService
	SafraService(svc *Service) contract.SafraService
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

func (s *serviceManager) SafraService(svc *Service) contract.SafraService {
	return newSafraService(svc)
}
