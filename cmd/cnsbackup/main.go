package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chiradeep/go-nitro/netscaler"
)

func main() {
	fmt.Println("Citrix NetScaler (ADC) Backup")
	fmt.Println("-----------------------------")

	var nitroURL = flag.String("url", "https://<hostname>", "Nitro URL - http(s)://<hostname>")
	var nitroUSer = flag.String("username", "nsroot", "Nitro API - username")
	var nitroPass = flag.String("password", "nsroot", "Nitro API - password")
	var nitroSSLVerify = flag.Bool("sslVerify", true, "Nitro API - Verify SSL Certificate")

	flag.Parse()

	nitroParams := netscaler.NitroParams{
		Url:       *nitroURL,
		Username:  *nitroUSer,
		Password:  *nitroPass,
		SslVerify: *nitroSSLVerify,
	}

	log.SetOutput(ioutil.Discard)

	nitroClient, err := netscaler.NewNitroClientFromParams(nitroParams)

	if err != nil {
		fmt.Println(err)
	}

	lbvserver, err := nitroClient.FindAllResources(netscaler.Lbvserver.Type())

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range lbvserver {
		fmt.Println(v["name"])
	}

}
