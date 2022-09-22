package main

//go:generate go run -mod=mod ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen generate --config database/graphql/gqlgen.yml
