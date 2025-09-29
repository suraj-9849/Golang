package user

type User struct{
	Email string
	Name string
}

func UserData(email string, name string) User{
	return User{email, name}
}