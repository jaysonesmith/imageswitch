package main

import (
	"os"

	"github.com/DATA-DOG/godog"
)

type tester interface {
	aBaseImage(imageType string) error
	thatImageIsConverted(imageType string) error
	aMustBeReturned(imageType string) error
}

type test struct {
	tester
	SC *ScenarioContext
}

type unit struct {
	test
}

type acceptance struct {
	test
}

func FeatureContext(s *godog.Suite) {
	var t test
	var sc *ScenarioContext

	switch os.Getenv("TEST_DEPTH") {
	case "acceptance":
		a := &acceptance{test: test{SC: NewScenarioContext()}}
		s.Step(`^a base (\w+) image$`, a.aBaseImage)
		s.Step(`^that image is converted to a (\w+)$`, a.thatImageIsConverted)
		s.Step(`^a (\w+) must be returned$`, a.aMustBeReturned)

		t.tester = a
	default:
		u := &unit{test: test{SC: NewScenarioContext()}}
		s.Step(`^a base (\w+) image$`, u.aBaseImage)
		s.Step(`^that image is converted to a (\w+)$`, u.thatImageIsConverted)
		s.Step(`^a (\w+) must be returned$`, u.aMustBeReturned)

		t.tester = u
	}

	s.BeforeScenario(func(interface{}) {
		sc = NewScenarioContext()
		t.SC = sc
	})
}

// Unit test level
func (u *unit) aBaseImage(imageType string) error {
	u.SC.BaseImageType = imageType
	return godog.ErrPending
}

func (u *unit) thatImageIsConverted(imageType string) error {
	return godog.ErrPending
}

func (u *unit) aMustBeReturned(imageType string) error {
	return godog.ErrPending
}

// Acceptance test level
func (a *acceptance) aBaseImage(imageType string) error {
	return godog.ErrPending
}

func (a *acceptance) thatImageIsConverted(imageType string) error {
	return godog.ErrPending
}

func (a *acceptance) aMustBeReturned(imageType string) error {
	return godog.ErrPending
}
