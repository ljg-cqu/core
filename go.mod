module github.com/ljg-cqu/core

go 1.18

replace (
	github.com/gin-gonic/gin v1.7.7 => github.com/ljg-cqu/gin v1.7.8
	github.com/long2ice/swagin v0.1.0 => github.com/ljg-cqu/swagin v0.1.4
	github.com/vgarvardt/gue/v3 v3.3.0 => github.com/ljg-cqu/gue/v3 v3.3.11-0.20220225033707-56cf5c188166
)

require (
	github.com/allegro/bigcache/v3 v3.0.2
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/appleboy/gofight/v2 v2.1.2
	github.com/brianvoe/gofakeit/v6 v6.14.5
	github.com/cockroachdb/cockroach-go/v2 v2.2.8
	github.com/danielgtaylor/huma v1.0.2
	github.com/deepmap/oapi-codegen v1.9.1
	github.com/georgysavva/scany v0.3.0
	github.com/getkin/kin-openapi v0.87.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/requestid v0.0.3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.0
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/go-resty/resty/v2 v2.7.0
	github.com/golang-jwt/jwt/v4 v4.3.0
	github.com/google/uuid v1.3.0
	github.com/jackc/pgconn v1.11.0
	github.com/jackc/pgtype v1.10.0
	github.com/jackc/pgx/v4 v4.15.0
	github.com/json-iterator/go v1.1.12
	github.com/lestrrat-go/backoff/v2 v2.0.8
	github.com/long2ice/swagin v0.1.0
	github.com/orandin/lumberjackrus v1.0.1
	github.com/pdfcpu/pdfcpu v0.3.13
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.6.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
	github.com/swaggest/jsonschema-go v0.3.24
	github.com/swaggest/openapi-go v0.2.15
	github.com/swaggest/rest v0.2.21
	github.com/swaggest/swgui v1.4.4
	github.com/swaggest/usecase v1.1.2
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.1
	github.com/swaggo/swag v1.8.0
	github.com/toorop/gin-logrus v0.0.0-20210225092905-2c785434f26f
	github.com/vgarvardt/gue/v3 v3.3.0
	github.com/wI2L/fizz v0.18.1
	github.com/xhit/go-simple-mail/v2 v2.10.0
	github.com/zc2638/swag v1.1.5
	github.com/zput/zxcTool v1.3.10
	go.uber.org/zap v1.19.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/Jeffail/gabs/v2 v2.6.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/fxamacker/cbor/v2 v2.2.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/goccy/go-json v0.9.0 // indirect
	github.com/goccy/go-yaml v1.8.1 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hhrutter/lzw v0.0.0-20190829144645-6f07a24e8650 // indirect
	github.com/hhrutter/tiff v0.0.0-20190829141212-736cae8d0bc7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/jinzhu/copier v0.3.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/option v1.0.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkgms/go v0.0.0-20201028070800-899b81726496 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v3 v3.1.0 // indirect
	github.com/shurcooL/httpgzip v0.0.0-20190720172056-320755c1c1b0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/swaggest/form/v5 v5.0.1 // indirect
	github.com/swaggest/refl v1.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/vearutop/statigz v1.1.5 // indirect
	github.com/vgarvardt/backoff v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
