package etrace

type (
	FuncParams map[string]interface{}

	Method struct {
		Package      string     `json:"package"`
		Type         string     `json:"type"`
		Name         string     `json:"name"`
		MethodParams FuncParams `json:"method_parameters,omitempty"`
		CauseFunc    string     `json:"cause_function,omitempty"`
		CauseMethod  string     `json:"cause_method,omitempty"`
		CauseParams  FuncParams `json:"cause_parameters,omitempty"`
	}

	Func struct {
		Package     string     `json:"package"`
		Name        string     `json:"name"`
		CauseFunc   string     `json:"cause_function,omitempty"`
		CauseParams FuncParams `json:"cause_parameters,omitempty"`
	}
)
