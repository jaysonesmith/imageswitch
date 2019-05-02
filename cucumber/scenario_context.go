package main

// ScenarioContext can contain data to be persisted across steps during a scenario
type ScenarioContext struct {
	BaseImageType    string
	DesiredImageType string
}

// NewScenarioContext returns a new ScenarioContext
func NewScenarioContext() *ScenarioContext {
	return &ScenarioContext{}
}
