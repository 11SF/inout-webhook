package coreline

type service struct {
}

type Service interface {
}

func NewService() *service {
	return &service{}
}
