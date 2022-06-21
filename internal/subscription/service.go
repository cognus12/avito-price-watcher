package subscription

type service struct {
	store Repository
}

type Service interface {
	Subscribe(url string, email string) error
	Unsubscribe(url string, email string) error
}

type adInfo struct {
	Url   string `json:"url"`
	Price uint64 `json:"price"`
}

func NewService(repository Repository) Service {
	return &service{
		store: repository,
	}
}

func (s *service) Subscribe(url, email string) error {
	err := s.store.CreateSubscibtion(url, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Unsubscribe(url, email string) error {
	err := s.store.DeleteSubscibtion(url, email)

	if err != nil {
		return err
	}

	return nil
}
