module github.com/deepfence/golang_deepfence_sdk/utils

go 1.19

replace github.com/deepfence/golang_deepfence_sdk/client => ../client

require (
	github.com/XSAM/otelsql v0.22.0
	github.com/deepfence/golang_deepfence_sdk/client v0.0.0-00010101000000-000000000000
	github.com/go-chi/jwtauth/v5 v5.1.0
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/lestrrat-go/jwx/v2 v2.0.9
	github.com/lib/pq v1.10.9
	github.com/minio/minio-go/v7 v7.0.52
	github.com/neo4j/neo4j-go-driver/v4 v4.4.7
	github.com/redis/go-redis/v9 v9.0.4
	github.com/rs/zerolog v1.29.1
	github.com/twmb/franz-go v1.13.4
	github.com/twmb/franz-go/pkg/kadm v1.8.1
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/trace v1.15.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/lestrrat-go/blackmagic v1.0.1 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/httprc v1.0.4 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.4.0 // indirect
	go.opentelemetry.io/otel/metric v0.38.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
