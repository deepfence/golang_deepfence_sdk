package directory

import (
	"context"

	postgresqlDb "github.com/deepfence/golang_deepfence_sdk/utils/postgresql/postgresql-db"
	"github.com/minio/minio-go/v7"

	"github.com/redis/go-redis/v9"
)

func getClient[T *redis.Client | *CypherDriver | *postgresqlDb.Queries | *minio.Client](ctx context.Context, pool map[NamespaceID]T, newClient func(DBConfigs) (T, error)) (T, error) {
	key, err := ExtractNamespace(ctx)
	if err != nil {
		return nil, err
	}
	val, has := pool[key]
	if has {
		return val, nil
	}

	client, err := newClient(directory[key])
	if err != nil {
		return nil, err
	}
	pool[key] = client
	return client, nil
}
