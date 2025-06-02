package shared

import infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"

type Container struct {
	DB *infra_db.DBConnections
}
