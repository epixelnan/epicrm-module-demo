default: oapi-api.gen.go oapi-server.gen.go oapi-types.gen.go

oapi-api.gen.go: api.yml
	~/go/bin/oapi-codegen --old-config-style -generate spec -o oapi-api.gen.go -package app api.yml

oapi-server.gen.go: api.yml
	~/go/bin/oapi-codegen --old-config-style -generate chi-server -o oapi-server.gen.go -package app api.yml

oapi-types.gen.go: api.yml
	~/go/bin/oapi-codegen --old-config-style -generate types,skip-prune -o oapi-types.gen.go -package app api.yml

clean:
	rm -f oapi-api.gen.go oapi-server.gen.go oapi-types.gen.go
