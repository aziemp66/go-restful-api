package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Value  interface{} `json:"value"`
}
