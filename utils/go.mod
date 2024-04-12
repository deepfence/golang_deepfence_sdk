module github.com/deepfence/golang_deepfence_sdk/utils

go 1.20

replace github.com/deepfence/golang_deepfence_sdk/client => ../client

require (
	github.com/deepfence/golang_deepfence_sdk/client v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-retryablehttp v0.7.5
)

require github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
