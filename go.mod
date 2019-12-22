module github.com/wealdtech/walletd

go 1.13

require (
	github.com/aws/aws-sdk-go v1.26.7 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.4.0
	github.com/wealdtech/go-ecodec v1.1.0 // indirect
	github.com/wealdtech/go-eth2-wallet v1.7.0
	github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4 v1.0.0
	github.com/wealdtech/go-eth2-wallet-hd v1.8.0
	github.com/wealdtech/go-eth2-wallet-store-filesystem v1.4.0
	github.com/wealdtech/go-eth2-wallet-store-s3 v1.4.0
	github.com/wealdtech/go-eth2-wallet-store-scratch v1.2.0
	github.com/wealdtech/go-eth2-wallet-types v1.7.0
	github.com/yuin/gopher-lua v0.0.0-20191220021717-ab39c6098bdb
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20191220142924-d4481acd189f // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20191220175831-5c49e3ecc1c1 // indirect
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace github.com/wealdtech/go-eth2-wallet => ../go-eth2-wallet