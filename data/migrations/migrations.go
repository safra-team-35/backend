package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_company_partners",
			Script: `CREATE TABLE IF NOT EXISTS tab_company_partners (
				id INT NOT NULL AUTO_INCREMENT,
				uuid CHAR(36) NOT NULL,
				name VARCHAR(300) NOT NULL,
				document_number VARCHAR(14) NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				UNIQUE INDEX UUID_UNIQUE (uuid ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     2,
			Description: "Inserting data on table tab_company_partners",
			Script: `
				INSERT INTO tab_company_partners 
					(id,uuid,name,document_number)
				VALUES
					(1,'66a4917e-39d1-463f-84b8-a31cded3dc20','Mercado Livre','12434595000172');
			`,
		},
		{
			Version:     3,
			Description: "Creating table tab_qrcode",
			Script: `CREATE TABLE IF NOT EXISTS tab_qrcode (
				id INT NOT NULL AUTO_INCREMENT,
				company_id INT NOT NULL,
				price DECIMAL(7,2) NOT NULL,
				expiration_time DATE NOT NULL,
				hash VARCHAR(5000) NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_qrcode_tab_company_partners_idx (id ASC),
				CONSTRAINT fk_company_business
					FOREIGN KEY (company_id)
					REFERENCES wallet_db.tab_company_partners (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     4,
			Description: "Creating table tab_payment_type",
			Script: `CREATE TABLE IF NOT EXISTS tab_payment_type (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     5,
			Description: "Inserting data on table tab_payment_type",
			Script: `
				INSERT INTO tab_payment_type 
					(id,name)
				VALUES
				(1,'saldo'),
				(2,'cartão de crédito'),
				(3,'saldo + cartão de crédito');
			`,
		},
	}
)
