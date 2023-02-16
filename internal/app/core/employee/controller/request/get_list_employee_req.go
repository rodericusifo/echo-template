package request

type GetListEmployeeRequestQuery struct {
	Page  int `query:"page" validate:"omitempty,min=0"`
	Limit int `query:"limit" validate:"omitempty,min=0"`
}

func (r *GetListEmployeeRequestQuery) CustomValidateRequestQuery() error {
	return nil
}
