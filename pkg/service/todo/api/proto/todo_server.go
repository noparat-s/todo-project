package todo

import context "context"

type server struct {
}

func (s *server) GetTodo(ctx context.Context, req *GetByID) (*Todo, error) {
	return nil, nil
}
