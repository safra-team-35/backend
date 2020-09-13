package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/safra-team-35/backend/data/migrations"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/infra/config"
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance returns an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")
	log.Println("Connection String: ", dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = mysql.SetLogger(logger.GetLogger())
	if err != nil {
		return nil, err
	}
	logger.Info("Database successfully configured")

	logger.Info("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)

	err = d.Migrate()
	if err != nil {
		return nil, err
	}

	logger.Info("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Ping returns the ping set
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

//QRCode returns the qrcode set
func (c *DBManager) QRCode() contract.QRCodeRepo {
	return newQRCodeRepo(c.db)
}

//Company returns the company set
func (c *DBManager) Company() contract.CompanyRepo {
	return newCompanyRepo(c.db)
}

//User returns the company set
func (c *DBManager) User() contract.UserRepo {
	return newUserRepo(c.db)
}
