package user

type UserService interface {
	CreateUser(name, email string) (*User, error)
	GetUser(id uint) (*User, error)
	UpdateUser(id uint, name, email string) (*User, error)
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) UserService {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(name, email string) (*User, error) {
	user := &User{Name: name, Email: email}
	err := s.repo.CreateUser(user)
	return user, err
}

func (s *Service) GetUser(id uint) (*User, error) {
	return s.repo.GetUser(id)
}

func (s *Service) UpdateUser(id uint, name, email string) (*User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	err = s.repo.UpdateUser(user)
	return user, err
}
