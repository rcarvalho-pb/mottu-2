package rpc_client

import (
	"fmt"
	"net/rpc"
)

func Call[K, T any](port, service string, data K, entry *T) (err error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Printf("Erro ao conectar ao servidor na port [%s]: %s\n", port, err)
		return
	}
	defer client.Close()
	err = client.Call(service, &data, &entry)
	if err != nil {
		return
	}
	return
}
