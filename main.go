package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tkak/terraform-provider-conoha/conoha"
)

func main() {
	plugin.Serve(conoha.Provider())
}
