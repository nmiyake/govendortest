package govendortest

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "github.com/jteeuwen/go-bindata"
)

func PublicFunc() {
  cli.NewApp().Run(os.Args)
  asset := bindata.Asset{}
  fmt.Println(asset)
}
