package conf_ent

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	// init postgres driver
	_ "github.com/lib/pq"

	"gitlab.com/mcsolutions/find-psy/back/common/cent"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

func GetRelationalDriver(relational *common.Relational) (*sql.Driver, error) {
	if relational == nil {
		return nil, cerrors.GetBadConfigError("relational")
	}
	if relational.Dialect == "" {
		return nil, cerrors.GetBadConfigError("relational.dialect")
	}
	if relational.Host == "" {
		return nil, cerrors.GetBadConfigError("relational.host")
	}
	if relational.Port == 0 {
		return nil, cerrors.GetBadConfigError("relational.port")
	}
	if relational.Dbname == "" {
		return nil, cerrors.GetBadConfigError("relational.dbname")
	}
	if relational.User == "" {
		return nil, cerrors.GetBadConfigError("relational.user")
	}
	if relational.Password == "" {
		return nil, cerrors.GetBadConfigError("relational.password")
	}
	if relational.SslMode == "" {
		return nil, cerrors.GetBadConfigError("relational.ssl_mode")
	}
	if relational.Dialect != dialect.Postgres {
		return nil, cerrors.NoPostgres
	}
	driver, err := sql.Open(
		dialect.Postgres,
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			relational.Host,
			relational.Port,
			relational.User,
			relational.Password,
			relational.Dbname,
			relational.SslMode,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := cent.SetSchema(driver, relational.Schema); err != nil {
		return nil, err
	}
	return driver, err
}
