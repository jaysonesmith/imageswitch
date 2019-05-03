package main

import "github.com/jaysonesmith/imageswitch/cucumber/support"

// ScenarioContext can contain data to be persisted across steps during a scenario
type ScenarioContext struct {
	BaseImageType      string
	DesiredImageType   string
	ConversionResponse support.ConversionResponse
}

// NewScenarioContext returns a new ScenarioContext
func NewScenarioContext() *ScenarioContext {
	return &ScenarioContext{}
}
