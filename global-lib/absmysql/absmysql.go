package absmysql

import (
	"absolutetech/farm_report/global-lib/envutils"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Connect connects with the default options.
func Connect(_ context.Context, url string, databaseName string, appName string) (*sql.DB, error) {
	if appName == "" {
		panic("application name must not be empty")
	}
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "abcom",            //os.Getenv("root"),
		Passwd:               "6EhsGWcMrJRB2gEJ", //os.Getenv("root"),
		Net:                  "tcp",
		Addr:                 url,
		DBName:               databaseName,
		AllowNativePasswords: true,
	}
	println(cfg.FormatDSN())

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectURIsEnv connects with a list of URIs per environments.
func ConnectURIsEnv(ctx context.Context, uris map[envutils.Env]string, env envutils.Env, databaseName string, appName string) (*sql.DB, error) {
	url, err := getURIEnv(uris, env)
	if err != nil {
		return nil, err
	}
	return Connect(ctx, url, databaseName, appName)
}

func getURIEnv(uris map[envutils.Env]string, env envutils.Env) (string, error) {
	u, ok := uris[env]
	if !ok {
		return "", nil
	}
	return u, nil
}

var uriIgrowDB = map[envutils.Env]string{
	envutils.Testing:     "uat.absolute.ag",
	envutils.Development: "uat.absolute.ag",
	envutils.Staging:     "uat.absolute.ag",
	envutils.Production:  "",
}

// ConnectIgrowDB connects to Igrow DB.
func ConnectIgrowlDB(ctx context.Context, env envutils.Env, appName string) (*sql.DB, error) {
	return ConnectURIsEnv(ctx, uriIgrowDB, env, DatabaseIGROW, appName)
}

// Database names.
const (
	DatabaseIGROW = "igrow"
)

// ForceDisconnect disconnects the client with a context that is already canceled.
// It forces closes all opened connections.
func ForceDisconnect(clt *sql.DB) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	return ctx.Err()
}

type Err func() error
