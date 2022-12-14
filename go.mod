module gitlab.com/mcsolutions/find-psy/back/article-v1

go 1.18

require (
	entgo.io/ent v0.11.2
	github.com/go-kratos/kratos/v2 v2.5.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/grokify/html-strip-tags-go v0.0.1
	github.com/prometheus/client_golang v1.12.1
	gitlab.com/mcsolutions/find-psy/back/common v0.0.0-00010101000000-000000000000
	gitlab.com/mcsolutions/find-psy/proto v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require (
	ariga.io/atlas v0.5.1-0.20220717122844-8593d7eb1a8e // indirect
	github.com/MicahParks/keyfunc v1.2.2 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/go-kratos/sentry v0.0.0-20211021071616-de3a2011c4e4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-redis/redis/extra/rediscmd v0.2.0 // indirect
	github.com/go-redis/redis/extra/redisotel v0.3.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.opentelemetry.io/otel v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0 // indirect
	go.opentelemetry.io/otel/sdk v1.9.0 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20211008194852-3b03d305991f // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.49.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace gitlab.com/mcsolutions/find-psy/proto => ../../proto

replace gitlab.com/mcsolutions/find-psy/back/common => ../common
