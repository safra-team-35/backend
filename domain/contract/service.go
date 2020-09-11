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
}
