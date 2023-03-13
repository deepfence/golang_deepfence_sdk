package directory

import (
	"context"
	"sync"

	postgresqlDb "github.com/deepfence/golang_deepfence_sdk/utils/postgresql/postgresql-db"
	"github.com/minio/minio-go/v7"

	"github.com/redis/go-redis/v9"
)

var locks []sync.Mutex

func init() {
	locks = make([]sync.Mutex, 5)
}

func getLock[T any]() *sync.Mutex {
	var val T
	switch any(val).(type) {
	case *redis.Client:
		return &locks[1]
	case *CypherDriver:
		return &locks[2]
	case *postgresqlDb.Queries:
		return &locks[3]
	case *minio.Client:
		return &locks[4]
	}
	// Global fallback that should never be reached
	return &locks[0]
}

func getClient[T *redis.Client | *CypherDriver | *postgresqlDb.Queries | *minio.Client](ctx context.Context, pool map[NamespaceID]T, newClient func(DBConfigs) (T, error)) (T, error) {
	key, err := ExtractNamespace(ctx)
	if err != nil {
		return nil, err
	}

	lock := getLock[T]()
	lock.Lock()
	defer lock.Unlock()

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
