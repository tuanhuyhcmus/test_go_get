package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

var token = os.Getenv("TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	secret, err := client.Logical().Read("secret/data/foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}
	fmt.Printf("hello: %v\n", m["hello"])
}
