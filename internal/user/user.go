package user

type User struct {
	ID         string
	Name       string
	DocumentID string
	Email      string
	Password   string
	UserType   string
	Balance    float64
}
type Repository interface {
	Get(id int) (*User, error)
	Create(e User) (int, error)
}

type UseCase interface {
	Get(id int) (*User, error)
	Create(User) (*User, error)
	Deposit(amount float64) (int, error)
}
