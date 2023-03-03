package directory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/XSAM/otelsql"
	postgresqlDb "github.com/deepfence/golang_deepfence_sdk/utils/postgresql/postgresql-db"
	_ "github.com/lib/pq"
)

const (
	ConsoleURLSettingKey = "console_url"
)

type SettingValue struct {
	Label       string      `json:"label"`
	Value       interface{} `json:"value"`
	Description string      `json:"description"`
}

var postgresClientsPool map[NamespaceID]*postgresqlDb.Queries

func init() {
	postgresClientsPool = map[NamespaceID]*postgresqlDb.Queries{}
}

func newPostgresClient(endpoints DBConfigs) (*postgresqlDb.Queries, error) {
	if endpoints.Postgres == nil {
		return nil, errors.New("No defined postgres config")
	}
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	endpoints.Postgres.Host, endpoints.Postgres.Port, endpoints.Postgres.Username, endpoints.Postgres.Password, endpoints.Postgres.Database, endpoints.Postgres.SslMode)
	// db, err := sql.Open("postgres", psqlInfo)
	psqlDSN := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		endpoints.Postgres.Username, endpoints.Postgres.Password,
		endpoints.Postgres.Host, endpoints.Postgres.Port,
		endpoints.Postgres.Database, endpoints.Postgres.SslMode)
	db, err := otelsql.Open("postgres", psqlDSN)
	if err != nil {
		return nil, err
	}
	return postgresqlDb.New(db), nil
}

func PostgresClient(ctx context.Context) (*postgresqlDb.Queries, error) {
	driver, err := getClient(ctx, postgresClientsPool, newPostgresClient)
	if err != nil {
		return nil, err
	}
	return driver, err
}

func ListSettings(ctx context.Context, visible bool) ([]postgresqlDb.Setting, error) {
	pgClient, err := PostgresClient(ctx)
	if err != nil {
		return nil, err
	}
	var settings []postgresqlDb.Setting
	if visible == true {
		settings, err = pgClient.GetVisibleSettings(ctx)
	} else {
		settings, err = pgClient.GetSettings(ctx)
	}
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func GetManagementConsoleURL(ctx context.Context) (string, error) {
	pgClient, err := PostgresClient(ctx)
	if err != nil {
		return "", err
	}
	setting, err := pgClient.GetSetting(ctx, ConsoleURLSettingKey)
	if err != nil {
		return "", err
	}
	var settingVal SettingValue
	err = json.Unmarshal(setting.Value, &settingVal)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", settingVal.Value), nil
}

func GetManagementHost(ctx context.Context) (string, error) {
	url, err := GetManagementConsoleURL(ctx)
	if err != nil {
		return "", err
	}
	// Remove scheme if present
	if strings.HasPrefix(url, "http://") {
		return url[7:], nil
	} else if strings.HasPrefix(url, "https://") {
		return url[8:], nil
	}
	return url, nil
}

func GetSetting(ctx context.Context, key string) (*postgresqlDb.Setting, error) {
	pgClient, err := PostgresClient(ctx)
	if err != nil {
		return nil, err
	}
	setting, err := pgClient.GetSetting(ctx, key)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}
