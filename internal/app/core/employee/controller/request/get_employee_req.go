package request

type GetEmployeeRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *GetEmployeeRequestParams) CustomValidateRequestParams() error {
	return nil
}
