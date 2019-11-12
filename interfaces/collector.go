package interfaces

type ICollector interface {
	Collect(url []string)
	AddUrl(url string)
	Store(response string)
}

type Collector struct {
	url []string
}
