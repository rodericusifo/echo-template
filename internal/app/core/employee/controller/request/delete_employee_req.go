package request

type DeleteEmployeeRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *DeleteEmployeeRequestParams) CustomValidateRequestParams() error {
	return nil
}
