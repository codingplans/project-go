module testgo/modgo

go 1.23.0

toolchain go1.24.1

require (
	github.com/Darrenzzy/person-go v1.0.7
	github.com/IBM/sarama v1.45.2
	github.com/RoaringBitmap/roaring v1.3.0
	github.com/fsnotify/fsnotify v1.5.4
	github.com/getsentry/sentry-go v0.14.0
	github.com/go-kratos/kratos v1.0.0
	github.com/jinzhu/copier v0.4.0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.12
	github.com/lib/pq v1.10.2
	github.com/looplab/fsm v0.2.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/panjf2000/ants/v2 v2.7.1
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.13.6
	github.com/prometheus/client_golang v1.17.0
	github.com/samber/lo v1.35.0
	github.com/shopspring/decimal v1.2.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.10.0
	github.com/tidwall/gjson v1.18.0
	github.com/valyala/fasthttp v1.40.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.38.0
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8
	golang.org/x/mod v0.18.0
	golang.org/x/sync v0.14.0
	golang.org/x/text v0.25.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
	testgo/modgo/crypto v0.0.0-20211102110000-000000000000
	xorm.io/xorm v1.3.2
)

replace testgo/modgo/crypto v0.0.0-20211102110000-000000000000 => ./crypto

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.8.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	xorm.io/builder v0.3.11-0.20220531020008-1bd24a7dc978 // indirect
)
