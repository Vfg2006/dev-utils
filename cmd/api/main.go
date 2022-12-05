package main

import (
	"context"
	"fmt"

	"github.com/vfg2006/dev-utils/config"
	"github.com/vfg2006/dev-utils/hasher"
	"github.com/vfg2006/dev-utils/qrcode"
	"github.com/vfg2006/dev-utils/server"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("INIT - SERVER")
	config, err := config.NewConfigFromFile()
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sha256 := hasher.NewSHA256()
	md5 := hasher.NewMD5()
	hasher := hasher.NewHasher(md5, sha256)

	qrcode := qrcode.New()

	server, err := server.New(
		config,
		hasher,
		qrcode,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := server.Run(ctx); err != nil {
		logrus.Error(err)
	}
}
