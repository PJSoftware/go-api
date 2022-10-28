package api

type QRYResult struct {
	Body string
}

func (r *QRYResult) Parse() string {
	return "-x-"
}