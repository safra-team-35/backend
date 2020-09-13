package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_company_partners",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_company_partners (
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
				INSERT INTO wallet_db.tab_company_partners 
					(id,uuid,name,document_number)
				VALUES
					(1,'66a4917e-39d1-463f-84b8-a31cded3dc20','Mercado Livre','12434595000172');
			`,
		},
		{
			Version:     3,
			Description: "Creating table tab_qrcode",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_qrcode (
				id INT NOT NULL AUTO_INCREMENT,
				company_id INT NOT NULL,
				price DECIMAL(7,2) NOT NULL,
				expiration_time DATE NOT NULL,
				hash VARCHAR(5000) NOT NULL,
				product_id INT NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_qrcode_tab_company_partners_idx (id ASC),
				CONSTRAINT fk_company_qrcode
					FOREIGN KEY (company_id)
					REFERENCES wallet_db.tab_company_partners (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     4,
			Description: "Creating table tab_payment_type",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_payment_type (
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
				INSERT INTO wallet_db.tab_payment_type 
					(id,name)
				VALUES
				(1,'saldo'),
				(2,'cartão de crédito'),
				(3,'saldo + cartão de crédito');
			`,
		},
		{
			Version:     6,
			Description: "Creating table tab_user",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_user (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL,
				birthdate DATE NOT NULL,
				document_number VARCHAR(14) NOT NULL,
				email VARCHAR(100) NOT NULL,
				phone_number VARCHAR(20) NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     7,
			Description: "Inserting data on table tab_user",
			Script: `
				INSERT INTO wallet_db.tab_user 
					(id,name,birthdate,document_number,email,phone_number)
				VALUES
				(1,'Diego Clair','1993-10-25','01234567890','diego93rodrigues@gmail.com','991313476');
			`,
		},
		{
			Version:     8,
			Description: "Creating table tab_user_address",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_user_address (
				id INT NOT NULL AUTO_INCREMENT,
				user_id INT NOT NULL,
				country VARCHAR(100) NOT NULL,
				street VARCHAR(3000) NOT NULL,
				number VARCHAR(11) NOT NULL,
				complement VARCHAR(100) NULL DEFAULT '',
				zip_code VARCHAR(20) NOT NULL,
				city VARCHAR(100) NOT NULL,
				federative_unit CHAR(2) NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_user_address_tab_user_idx (id ASC),
				CONSTRAINT fk_user_user_address
					FOREIGN KEY (user_id)
					REFERENCES wallet_db.tab_user (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     9,
			Description: "Inserting data on table tab_user_address",
			Script: `
				INSERT INTO wallet_db.tab_user_address 
					(id,user_id,country,street,number,complement,zip_code,city,federative_unit)
				VALUES
				(1,1,'Brasil','Av Paulista','2100','5 andar','01310300','São Paulo','SP'),
				(2,1,'Brasil','Rua Bela Cintra','2182','','01415008','São Paulo','SP');
			`,
		},
		{
			Version:     10,
			Description: "Creating table tab_order",
			Script: `CREATE TABLE IF NOT EXISTS wallet_db.tab_order (
				id INT NOT NULL AUTO_INCREMENT,
				user_id INT NOT NULL,
				company_id INT NOT NULL,
				price DECIMAL(7,2) NOT NULL,
				qrcode_id INT NOT NULL,
				product_id INT NOT NULL,
				payment_type INT NOT NULL,
				address_id INT NOT NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				INDEX fk_tab_order_tab_company_partners_idx (id ASC),
				INDEX fk_tab_order_tab_user_address_idx (id ASC),
				INDEX fk_tab_order_tab_user_idx (id ASC),
				CONSTRAINT fk_company_order
					FOREIGN KEY (company_id)
					REFERENCES wallet_db.tab_company_partners (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_address_order
					FOREIGN KEY (address_id)
					REFERENCES wallet_db.tab_user_address (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_user_order
					FOREIGN KEY (user_id)
					REFERENCES wallet_db.tab_user (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
				CONSTRAINT fk_qrcode_order
					FOREIGN KEY (qrcode_id)
					REFERENCES wallet_db.tab_qrcode (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     11,
			Description: "Adding column freight to table order",
			Script:      "ALTER TABLE tab_order ADD COLUMN freight DECIMAL(7,2) NOT NULL AFTER price;",
		},
		{
			Version:     12,
			Description: "Changing column freight",
			Script:      "ALTER TABLE tab_order MODIFY freight DECIMAL(7,2) NOT NULL DEFAULT 0.00 AFTER price;",
		},
	}
)
