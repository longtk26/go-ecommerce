package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}	

func (ur *UserRepo) GetUserByID() string {
	return "id_user"
}