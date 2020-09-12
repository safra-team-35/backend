package mysql

import (
	"database/sql"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain/entity"
)

type userRepo struct {
	db *sql.DB
}

// newUserRepo returns a instance of dbrepo
func newUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (s *userRepo) GetUserAddress(userID int64) (addresses []entity.Address, restErr resterrors.RestErr) {

	query := `
		SELECT
		tua.id,
		tua.country,
		tua.street,
		tua.number,
		tua.complement,
		tua.zip_code,
		tua.city,
		tua.federative_unit

		FROM 	tab_user_address 	tua
		WHERE  	tua.user_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		logger.Error("GetUserAddress", err)
		return addresses, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		logger.Error("GetUserAddress", err)
		return addresses, resterrors.NewInternalServerError("Database error")
	}

	var address entity.Address
	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.Country,
			&address.Street,
			&address.Number,
			&address.Complement,
			&address.ZipCode,
			&address.City,
			&address.FederativeUnit,
		)
		if err != nil {
			logger.Error("GetUserAddress", err)
			return nil, mysqlutils.HandleMySQLError(err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
