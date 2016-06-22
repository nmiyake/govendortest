package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "github.com/jteeuwen/go-bindata"
)

func main() {
  cli.NewApp().Run(os.Args)
  asset := bindata.Asset{}
  fmt.Println(asset)
}
