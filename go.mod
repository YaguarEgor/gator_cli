module github.com/YaguarEgor/gator_cli

go 1.23.4

replace github.com/YaguarEgor/gator_cli/internal/config v0.0.0 => ./internal/config

require github.com/YaguarEgor/gator_cli/internal/config v0.0.0

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)
