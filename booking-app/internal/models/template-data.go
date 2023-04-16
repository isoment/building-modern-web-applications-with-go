package models

// These are all the different data types we can pass to templates
type TemplateData struct {
	StringMap map[string]string
	IntMat    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
