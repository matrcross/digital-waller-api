package user

type Service struct {
	userRepo Repository
}

func NewService(userRepo Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) Get(id int) (*User, error) {
	return &User{}, nil
}

func (s *Service) Create(user User) (*User, error) {
	return &User{}, nil
}

func (s *Service) Deposit(amount float64) (int, error) {
	return 1, nil
}
