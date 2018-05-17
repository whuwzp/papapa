package deffer




type PaserResult struct {
	Requests []Request
	Items []interface{}
}



type Request struct {
	Url string
	PaserFunc func([]byte) PaserResult
}

func NilPaser([]byte) PaserResult{
	return PaserResult{}
}