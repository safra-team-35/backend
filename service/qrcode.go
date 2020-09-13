package service

import (
	"net/http"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/go_utils-lib/validstruct"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/domain/entity"
	"github.com/safra-team-35/backend/utils/cryptoutils"
)

type qrcodeService struct {
	svc    *Service
	client *http.Client
}

//newQRCodeService return a new instance of the service
func newQRCodeService(svc *Service) contract.QRCodeService {
	client := svc.httpClient
	return &qrcodeService{
		svc:    svc,
		client: client,
	}
}

func (s *qrcodeService) CreateQRCode(qrcode entity.QRCode, uuid string) (string, resterrors.RestErr) {

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

func (s *qrcodeService) GetQRCodeDataByHash(hash string) (qrcode entity.QRCode, err resterrors.RestErr) {
	return s.svc.db.QRCode().GetByHash(hash)
}
