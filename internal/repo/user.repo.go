package repo

type UserRepo struct { }

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

type UserInfo struct { 
	Id int `json:"id"`
	Name string `json:"name"`
	Old int `json:"old"`

}

func (ur *UserRepo) GetInfoUser() UserInfo {	
	return UserInfo  {	 
		Id: 1,
		Name: "Alice",
		Old: 18,
	}
}