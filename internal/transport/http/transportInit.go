package http

import (
	"forum/internal/business"
	businessrealiz "forum/internal/business/businessRealiz"
	"forum/internal/transport"
)

type Transport struct {
	service businessrealiz.Service
}

func InitTransport(b business.Business) (transport.Transport, error) {
	var t Transport

	return t, nil
}
