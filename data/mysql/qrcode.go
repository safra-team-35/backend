package mysql

import (
	"database/sql"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain/entity"
)

type qrcodeRepo struct {
	db *sql.DB
}

// newQRCodeRepo returns a instance of dbrepo
func newQRCodeRepo(db *sql.DB) *qrcodeRepo {
	return &qrcodeRepo{
		db: db,
	}
}

func (s *qrcodeRepo) CreateQRCode(qrcode entity.QRCode) resterrors.RestErr {

	query := `
		INSERT INTO tab_qrcode (
			company_id,
			price,
			product_id,
			hash,
			expiration_time) 
		VALUES	
			(?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		logger.Error("CreateQRCode", err)
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		qrcode.CompanyID,
		qrcode.Price,
		qrcode.ProductID,
		qrcode.Hash,
		qrcode.ExpirationTime)
	if err != nil {
		logger.Error("CreateQRCode", err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *qrcodeRepo) GetByHash(hash string) (qrcode entity.QRCode, restErr resterrors.RestErr) {
	query := `
		SELECT 	
			tq.company_id,
			tcp.name,
			tq.price
			
		FROM 	tab_qrcode 	tq
		INNER JOIN tab_company_partners tcp
			ON tcp.id 	= tq.company_id

		WHERE  	tq.hash = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		logger.Error("GetByHash", err)
		return qrcode, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(hash)

	err = result.Scan(
		&qrcode.CompanyID,
		&qrcode.CompanyName,
		&qrcode.Price,
	)

	if err != nil {
		logger.Error("GetByHash", err)
		return qrcode, mysqlutils.HandleMySQLError(err)
	}

	return qrcode, nil
}
