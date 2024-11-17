package models

type Route struct {
	Path          string
	TargetService string
	Method        string
	RateLimit     int
	RequirePath   bool
	CacheDuration int //in sec
	Timeout       int // in sec
}

type ServiceResponse struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}
