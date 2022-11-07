package api

type QRYResult struct {
	Body []byte
}

func (r *QRYResult) Parse() string {
	return "-x-"
}