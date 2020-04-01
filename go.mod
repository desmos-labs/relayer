module github.com/iqlusioninc/relayer

go 1.13

require (
	github.com/avast/retry-go v2.6.0+incompatible
	github.com/bitsongofficial/go-bitsong v0.3.1-0.20200401075836-7d426a6c432f
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200327170214-3b48464bb4dc
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/desmos-labs/desmos v0.3.1-0.20200401075749-37c04e1f2bba
	github.com/gorilla/mux v1.7.4
	github.com/ory/dockertest/v3 v3.5.4
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/sirkon/goproxy v1.4.8
	github.com/sirupsen/logrus v1.5.0 // indirect
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.2
	github.com/tendermint/tm-db v0.5.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
