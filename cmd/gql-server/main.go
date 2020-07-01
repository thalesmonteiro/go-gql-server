package main

import (
	log "github.com/znobrega/go-gql-server/internal/logger"
	"github.com/znobrega/go-gql-server/internal/orm"
	"github.com/znobrega/go-gql-server/pkg/server"
)

func main() {
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	server.Run(orm)
}
