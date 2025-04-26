package services

import "github.com/go-playground/validator/v10"

type (
	ErrorResponse struct {
		FailedField string `json:"failedField"`
		Tag         string `json:"tag"`
		Value       string `json:"value,omitempty"`
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func (s *Service) Validate(data any) []ErrorResponse {
	var validationErrors []ErrorResponse

	err := s.Validator.Struct(data)
	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorResponse{
				FailedField: fieldErr.Field(),
				Tag:         fieldErr.ActualTag(),
				Value:       fieldErr.Param(),
			})
		}
	}

	return validationErrors
}
