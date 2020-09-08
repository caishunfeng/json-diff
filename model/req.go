package model

type DiffReq struct {
	TraceId      string
	EndPoints    EndPoint
	QueryParams  map[string]string
	HeaderParams map[string]string
}

type EndPoint struct {
	Live  string
	Local string
}
