module cloudeng.io/glean/gleancli

go 1.25

require (
	cloudeng.io/aws v0.0.0-20251204190401-6fa1f48d333e
	cloudeng.io/cmdutil v0.0.0-20251204190401-6fa1f48d333e
	cloudeng.io/errors v0.0.13-0.20251104042927-f7e1e5e3ef21
	cloudeng.io/file v0.0.0-20251204190401-6fa1f48d333e
	cloudeng.io/glean/crawlindex v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/extensions/benchling v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/extensions/biorxiv v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/extensions/papersapp v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/extensions/protocolsio v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/extensions/testcmd v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/glean/gleansdk v0.0.0-20240327164852-a1bc676b508c
	cloudeng.io/path v0.0.10-0.20251104042927-f7e1e5e3ef21
	cloudeng.io/security v0.0.0-20251204190401-6fa1f48d333e
	cloudeng.io/sync v0.0.9-0.20251104042927-f7e1e5e3ef21
	cloudeng.io/webapi/operations v0.0.0-20251201181813-e5a358b3f6fd
	github.com/oasdiff/yaml v0.0.0-20250309154309-f31be36b4037
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloudeng.io/algo v0.0.0-20251204190401-6fa1f48d333e // indirect
	cloudeng.io/glean/gleanclientsdk v0.0.0-20240327164852-a1bc676b508c // indirect
	cloudeng.io/logging v0.0.0-20251204190401-6fa1f48d333e // indirect
	cloudeng.io/net v0.0.0-20251204190401-6fa1f48d333e // indirect
	cloudeng.io/os v0.0.0-20251204190401-6fa1f48d333e // indirect
	cloudeng.io/sys v0.0.0-20251204190401-6fa1f48d333e // indirect
	cloudeng.io/text v0.0.13 // indirect
	cloudeng.io/webapi/benchling v0.0.0-20240304015302-6256cc401d35 // indirect
	cloudeng.io/webapi/biorxiv v0.0.0-20240304015302-6256cc401d35 // indirect
	cloudeng.io/webapi/clients/benchling v0.0.0-20251201181813-e5a358b3f6fd // indirect
	cloudeng.io/webapi/clients/biorxiv v0.0.0-20251201181813-e5a358b3f6fd // indirect
	cloudeng.io/webapi/clients/papersapp v0.0.0-20251201181813-e5a358b3f6fd // indirect
	cloudeng.io/webapi/clients/protocolsio v0.0.0-20251201181813-e5a358b3f6fd // indirect
	cloudeng.io/webapi/papersapp v0.0.0-20240304015302-6256cc401d35 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2 v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.4 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.32.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.93.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.0.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.41.3 // indirect
	github.com/aws/smithy-go v1.24.0 // indirect
	github.com/getkin/kin-openapi v0.133.0 // indirect
	github.com/go-openapi/errors v0.22.4 // indirect
	github.com/go-openapi/jsonpointer v0.22.3 // indirect
	github.com/go-openapi/strfmt v0.25.0 // indirect
	github.com/go-openapi/swag v0.25.4 // indirect
	github.com/go-openapi/swag/cmdutils v0.25.4 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/fileutils v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/mangling v0.25.4 // indirect
	github.com/go-openapi/swag/netutils v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/labstack/echo/v4 v4.13.4 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mailru/easyjson v0.9.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oapi-codegen/runtime v1.1.2 // indirect
	github.com/oasdiff/yaml3 v0.0.0-20250309153720-d2182401db90 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/woodsbury/decimal128 v1.4.0 // indirect
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/oauth2 v0.33.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
