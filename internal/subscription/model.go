package subscription

import "apricescrapper/internal/apperror"

type SubscribtionDTO struct {
	Url   string
	Email string
}

type Subscribtion struct {
	Url         string
	Subscribers []string
}

func (dto *SubscribtionDTO) Validate() error {
	err := apperror.UnprocessableEntity

	if dto.Email == "" && dto.Url == "" {
		err.Message = "Email and URL are not valid"

		return err
	}

	if dto.Email == "" {
		err.Message = "Email is not valid"

		return err
	}

	if dto.Url == "" {
		err.Message = "URL is not valid"

		return err
	}

	return nil
}
