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
)

func (method Method) OfName(name string) Method {
	method.Name = name
	return method
}

func (method Method) OfParams(params FuncParams) Method {
	method.MethodParams = params
	return method
}

func (method Method) OfCauseFunc(funcName string) Method {
	method.CauseFunc = funcName
	return method
}
func (method Method) OfCauseMethod(methodName string) Method {
	method.CauseMethod = methodName
	return method
}

func (method Method) OfCauseParams(params FuncParams) Method {
	method.CauseParams = params
	return method
}
