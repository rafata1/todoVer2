package core

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	protopack "todoapp/generated_proto"
)

type Server struct {
	protopack.UnimplementedManagePostServiceServer
	myDB *sqlx.DB
}

func NewServer() *Server {
	return &Server{
		myDB: connectToDB(),
	}
}

func (s *Server) AddPost(ctx context.Context, req *protopack.AddPostRequest) (*protopack.AddPostResponse, error) {
	title := req.Title
	content := req.Content
	// TODO: add new post item into db

	queryCode := `INSERT INTO tuts (title, content) VALUES (?, ?)`
	s.myDB.MustExec(queryCode, title, content)

	time.Sleep(5 * time.Second)

	return &protopack.AddPostResponse{
		StatusCode: 200,
	}, nil
}
