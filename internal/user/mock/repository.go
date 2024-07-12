package mock

import (
	"github.com/matrcross/digital-wallet-api/internal/user"
)

type mockRepository struct {
	db string
}

func NewMockRepository() user.Repository {
	return &mockRepository{
		db: "mock",
	}
}

func (r *mockRepository) Get(id int) (*user.User, error) {
	u := &user.User{}
	return u, nil
}

// Create implements user.Repository.
func (*mockRepository) Create(e user.User) (int, error) {
	return 1, nil
}
