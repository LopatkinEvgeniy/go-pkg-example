package main

import (
	"fmt"
	"os"

	"github.com/LopatkinEvgeniy/go-pkg-example/pkg/deep/level1/level2/deeplib"
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
			sum := mylib.Add(a, b)
			fmt.Printf("sum of %d and %d is %d\n", a, b, sum)

			if deeplib.IsOdd(sum) {
				fmt.Println("sum is odd")
			} else {
				fmt.Println("sum is even")
			}
		},
	}

	cmd.Flags().Int64Var(&a, "a", 0, "first number")
	cmd.Flags().Int64Var(&b, "b", 0, "second number")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
