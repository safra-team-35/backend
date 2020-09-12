package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain/entity"
)

// PingService holds a ping service operations
type PingService interface {
}

// QRCodeService holds a qrcode service operations
type QRCodeService interface {
	CreateQRCode(qrcode entity.QRCode, uuid string) (hash string, err resterrors.RestErr)
	GetQRCodeDataByHash(hash string) (qrcode entity.QRCode, err resterrors.RestErr)
}

// UserService holds a user service operations
type UserService interface {
	GetUserAddress(userID int64) (address []entity.Address, err resterrors.RestErr)
	CreateOrder(order entity.Order) (orderNumber string, err resterrors.RestErr)
}
