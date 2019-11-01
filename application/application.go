package application

// Application is a interface describe what a application requires (application life cycle, event handler)
type Application interface {
	// WillLaunch call before application is launch
	WillLaunch(mode string, configuration map[string]string)

	// DidFinishLaunching call after all application launch logic is completed
	DidFinishLaunching()

	// WillTerminate call before application Terminate
	WillTerminate()
}

// BaseApplication implements base application logic
type BaseApplication struct {
	Application
	Mode          string
	Configuration map[string]string
}

// WillLaunch call before application is launch
func (b *BaseApplication) WillLaunch(mode string, configuration map[string]string) {
	b.Mode = mode
	b.Configuration = configuration
}

// DidFinishLaunching call after all application launch logic is completed
func (b *BaseApplication) DidFinishLaunching() {

}

// WillTerminate call before application Terminate
func (b *BaseApplication) WillTerminate() {

}
