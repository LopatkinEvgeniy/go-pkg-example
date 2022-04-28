module github.com/LopatkinEvgeniy/go-pkg-example

go 1.17

require (
	github.com/LopatkinEvgeniy/go-pkg-example/pkg/mylib v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.4.0
)

replace github.com/LopatkinEvgeniy/go-pkg-example/pkg/mylib => ./pkg/mylib

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
