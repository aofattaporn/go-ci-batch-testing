package server

import "github.com/aofattaporn/go-cobra/configs"

type IServer interface {
	Start()
}

type fiberServer struct {
	cfg configs.IConfig
}

func NewFiberServer(cfg configs.IConfig) (IServer, error) {

	// TODO : database connection

	// TODO : Azure Identify

	// TODO : Blob connection

	return &fiberServer{
		cfg: cfg,
	}, nil
}

func (s *fiberServer) Start() {

}
