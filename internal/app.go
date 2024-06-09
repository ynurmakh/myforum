package internal

import (
	"fmt"

	"forum/internal/storage/sqlite3"
	"forum/internal/transport/http"

	businessrealiz "forum/internal/business/businessRealiz"
)

func Run() error {
	// TODO Logger create/init

	storage, err := sqlite3.InitStorage()
	if err != nil {
		return err
	}

	service, err := businessrealiz.InitService(storage)
	if err != nil {
		return err
	}

	transport, err := http.InitTransport(service)
	if err != nil {
		return err
	}

	fmt.Println("Server START")
	err = transport.Start()
	if err != nil {
		return err
	}
	return nil
}
