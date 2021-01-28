rm -rf temp/temp_resolver.go
go run github.com/99designs/gqlgen generate
go run runner/runner.go