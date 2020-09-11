package service

import (
	"fmt"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/go_utils-lib/validstruct"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/domain/entity"
	"github.com/safra-team-35/backend/utils/cryptoutils"
)

type qrcodeService struct {
	svc *Service
}

//newQRCodeService return a new instance of the service
func newQRCodeService(svc *Service) contract.QRCodeService {
	return &qrcodeService{
		svc: svc,
	}
}

func (s *qrcodeService) CreateQRCode(qrcode entity.QRCode, uuid string) (string, resterrors.RestErr) {

	fmt.Println("uuid: ", uuid)
	companyID, err := s.svc.db.Company().GetCompanyIDByUUID(uuid)
	if err != nil {
		return "", err
	}

	qrcode.CompanyID = companyID

	qrcode.Hash = cryptoutils.GenerateHashFromStruct(qrcode)

	err = validstruct.ValidateStruct(qrcode)
	if err != nil {
		return "", err
	}

	err = s.svc.db.QRCode().CreateQRCode(qrcode)
	if err != nil {
		return "", err
	}

	return qrcode.Hash, nil
}
