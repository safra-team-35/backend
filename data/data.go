package data

import (
	"github.com/safra-team-35/backend/data/mysql"
	"github.com/safra-team-35/backend/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
