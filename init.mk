
remove-api-third_party:
	@rm -f -r \
	api \
	third_party \
	openapi.yaml \
	generate.go

change-internal:
	@sed -i 's|greeter|article|g' internal/*/*.go
	@sed -i 's|Greeter|Article|g' internal/*/*.go
	@sed -i 's|greeter|article|g' internal/*/*/*.go
	@sed -i 's|Greeter|Article|g' internal/*/*/*.go
	@mv internal/biz/greeter.go internal/biz/article.go
	@mv internal/outside/data/greeter.go internal/outside/data/article.go
	@mv internal/service/greeter.go internal/service/article.go
	@mv internal/service/greeter_test.go internal/service/article_test.go
	@sed -i 's|v1\.|pb\.|g' internal/*/*.go
	@sed -i 's|"gitlab.com/mcsolutions/find-psy/back/article-v1/api/helloworld/v1"|"gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"|g' internal/*/*.go
	@goimports -w internal/server internal/service

change-mod:
	@sed -i "s|go 1.16|go 1.18|g" go.mod
	@sed -i "s|go 1.17|go 1.18|g" go.mod
	@echo "\nreplace gitlab.com/mcsolutions/find-psy/proto => ../../proto" >> go.mod
	@echo "\nreplace gitlab.com/mcsolutions/find-psy/back/common => ../common" >> go.mod

change-main:
	@sed -i 's|Name string|Name = \"article-v1\"|g' cmd/article-v1/main.go
	@sed -i 's|flagconf|flagConf|g' cmd/article-v1/main.go
	@sed -i 's|bc|conf|g' cmd/article-v1/main.go
	@sed -i 's|../../configs|./configs|g' cmd/article-v1/main.go
	@goimports -w cmd

change-config:
	@sed -i 's|8000|$(PORT_HTTP)|g' configs/config.yaml
	@sed -i 's|9000|$(PORT_GRPC)|g' configs/config.yaml
	@sed -i 's|0.0.0.0:|:|g' configs/config.yaml

init-all: remove-api-third_party change-internal change-mod change-main change-config dep wire
	@du -shc vendor/*
	@rm -f -r .git
	@git init
	@git remote add origin https://gitlab.com/mcsolutions/find-psy/back/article-v1.git
	@git add .
	@git commit -m Init
	@git tag -a v0.0.1 -m Init
