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

func (s *userService) GetUserOrdersSummary(userID int64) (summary []entity.OrderSummary, err resterrors.RestErr) {

	return s.svc.db.User().GetOrderSummary(userID)
}

func (s *userService) CreateOrder(order entity.Order) (orderNumber string, err resterrors.RestErr) {

	orderDetail, err := s.svc.db.QRCode().GetByHash(order.Hash)
	if err != nil {
		return orderNumber, err
	}

	order.CompanyID = orderDetail.CompanyID
	order.Price = orderDetail.Price
	order.ProductID = orderDetail.ProductID
	order.QRCodeID = orderDetail.ID

	err = s.svc.db.User().CreateOrder(order)
	if err != nil {
		return orderNumber, err
	}

	orderNumber, err = s.sendOrderToPartner()
	if err != nil {
		return orderNumber, err
	}

	err = s.processPayment()
	if err != nil {
		return orderNumber, err
	}
	return orderNumber, nil
}

func (s *userService) sendOrderToPartner() (orderNumber string, err resterrors.RestErr) {
	return "4587T06B", nil
}

func (s *userService) processPayment() (err resterrors.RestErr) {
	return
}
