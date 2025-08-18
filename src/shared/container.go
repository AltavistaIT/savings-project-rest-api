package shared

import (
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	infra_redis "github.com/ssssshel/sp-api/src/infraestructure/redis"
)

type Container struct {
	DB    *infra_db.DBConnections
	Redis *infra_redis.RedisConnection
}
