package utils

import (
	"context"
	"encoding/json"
	"fmt"

	postgresqlDb "github.com/deepfence/golang_deepfence_sdk/utils/postgresql/postgresql-db"
)

const (
	ConsoleURLSettingKey = "console_url"
)

type SettingValue struct {
	Label       string      `json:"label"`
	Value       interface{} `json:"value"`
	Description string      `json:"description"`
}

func ListSettings(ctx context.Context, pgClient *postgresqlDb.Queries, visible bool) ([]postgresqlDb.Setting, error) {
	var settings []postgresqlDb.Setting
	var err error
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

func GetManagementConsoleURL(ctx context.Context, pgClient *postgresqlDb.Queries) (string, error) {
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

func GetSetting(ctx context.Context, pgClient *postgresqlDb.Queries, key string) (*postgresqlDb.Setting, error) {
	setting, err := pgClient.GetSetting(ctx, key)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}
