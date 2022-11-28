package app

type YourApp struct {
	// List Of Services
}

func NewYourApp() *YourApp {
	return &YourApp{}
}

func (y *YourApp) SetYourService() *YourApp {
	// Inject services
	return y
}
