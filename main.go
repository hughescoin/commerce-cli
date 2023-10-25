package main

import (
	cli "github.com/hughescoin/commerce-cli/cmd"
	"github.com/hughescoin/commerce-cli/sdk"
)

func main() {
	sdk.InitClient()
	cli.Execute()
}
