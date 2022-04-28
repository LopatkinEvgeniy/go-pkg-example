package main

import (
	"fmt"
	"os"

	"github.com/LopatkinEvgeniy/go-pkg-example/pkg/mylib"
	"github.com/spf13/cobra"
)

func main() {
	var a int64
	var b int64

	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("sum of %d and %d is %d\n", a, b, mylib.Add(a, b))
		},
	}

	cmd.Flags().Int64Var(&a, "a", 0, "first number")
	cmd.Flags().Int64Var(&b, "b", 0, "second number")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
