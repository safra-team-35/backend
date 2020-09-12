package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/domain/entity"
)

type userService struct {
	svc *Service
}

//newUserService return a new instance of the service
func newUserService(svc *Service) contract.UserService {
	return &userService{
		svc: svc,
	}
}

func (s *userService) GetUserAddress(userID int64) (address []entity.Address, err resterrors.RestErr) {
	return s.svc.db.User().GetUserAddress(userID)
}
