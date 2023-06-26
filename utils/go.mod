module github.com/deepfence/golang_deepfence_sdk/utils

go 1.20

replace github.com/deepfence/golang_deepfence_sdk/client => ../client

require (
	github.com/deepfence/golang_deepfence_sdk/client v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-retryablehttp v0.7.4
	github.com/rs/zerolog v1.29.1
)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
)
