package etrace

func (f Func) OfName(funcName string) Func {
	f.Name = funcName
	return f
}
