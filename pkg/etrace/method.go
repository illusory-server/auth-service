package etrace

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
