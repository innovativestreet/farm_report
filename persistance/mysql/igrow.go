package mysql

import (
	"absolutetech/farm_report/global-lib/absmysql"
	"absolutetech/farm_report/global-lib/envutils"
	"context"
	"database/sql"
)

func NewIgrowDBClient(env envutils.Env, appName string) (*sql.DB, absmysql.Err, error) {
	client, err := absmysql.ConnectIgrowlDB(context.Background(), env, appName)
	if err != nil {
		return nil, nil, err
	}
	cl := func() error {
		return absmysql.ForceDisconnect(client)
	}
	return client, cl, nil
}
