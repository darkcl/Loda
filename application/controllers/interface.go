package controllers

// Controller implements all logic between UI and main application
type Controller interface {
	// Load is called when application is loaded
	Load(context map[string]interface{})
}
