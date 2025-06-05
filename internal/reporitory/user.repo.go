package reporitory

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserById() string {
	return "my user"
}
