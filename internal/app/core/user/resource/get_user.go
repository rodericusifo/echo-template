package resource

import (
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (r *UserResource) GetUser(query *types.Query) (*sql.User, error) {
	return r.UserDatabaseSQLRepository.GetUser(query)
}
