package store

import ("database/sql"
"context"
)
type PostsStore struct {
	db *sql.DB
}
func (s *PostsStore) Create (ctx  context.Context)error{
	return nil
}