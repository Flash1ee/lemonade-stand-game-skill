package usecase

type LemonadeGameUsecase struct {
	repo GameRepo
}

// New -.
func New(r GameRepo) *LemonadeGameUsecase {
	return &LemonadeGameUsecase{
		repo: r,
	}
}

func (u *LemonadeGameUsecase) CreateUser() string {
	return u.repo.CreateUser()
}
