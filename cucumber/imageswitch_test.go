package main

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/imageswitch/cucumber/support"
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

func FeatureContext(s *godog.Suite) {
	var t test

	switch os.Getenv("TEST_DEPTH") {
	case "acceptance":
		a := &acceptance{test: test{SC: NewScenarioContext()}}
		s.Step(`^a base (\w+) image$`, a.aBaseImage)
		s.Step(`^that image is converted to a (\w+)$`, a.thatImageIsConverted)
		s.Step(`^a (\w+) must be returned$`, a.aMustBeReturned)

		t.tester = a
	case "integrated":
		i := &integrated{test: test{SC: NewScenarioContext()}}
		s.Step(`^a base (\w+) image$`, i.aBaseImage)
		s.Step(`^that image is converted to a (\w+)$`, i.thatImageIsConverted)
		s.Step(`^a (\w+) must be returned$`, i.aMustBeReturned)

		t.tester = i
	default:
		u := &unit{test: test{SC: NewScenarioContext()}}
		s.Step(`^a base (\w+) image$`, u.aBaseImage)
		s.Step(`^that image is converted to a (\w+)$`, u.thatImageIsConverted)
		s.Step(`^a (\w+) must be returned$`, u.aMustBeReturned)

		t.tester = u
	}

	s.BeforeScenario(func(interface{}) {
		t.SC = NewScenarioContext()
	})
}

// Unit test level
// Call functions directly
type unit struct {
	test
}

func (u *unit) aBaseImage(imageType string) error {
	u.SC.BaseImageType = imageType

	return nil
}

func (u *unit) thatImageIsConverted(desiredFormat string) error {
	imageURL := fmt.Sprintf("https://localhost/gopher.%s", u.SC.BaseImageType)
	// request := support.ConversionRequestBody(imageURL, desiredFormat)
	// call the conversion function directly, for now hard code
	response := support.ConversionResponse{
		OriginalImage: imageURL,
		DesiredFormat: desiredFormat,
		NewImage:      fmt.Sprintf("https://localhost/gopher.%s", desiredFormat),
	}

	u.SC.ConversionResponse = response

	return nil
}

func (u *unit) aMustBeReturned(imageType string) error {
	convertedExtention := support.ImageExtention(u.SC.ConversionResponse.NewImage)
	if convertedExtention != imageType {
		return fmt.Errorf("image conversion failed. expected: %s found: %s", imageType, convertedExtention)
	}

	return nil
}

// Acceptance test level
// Call mocked urls
type acceptance struct {
	test
}

func (a *acceptance) aBaseImage(imageType string) error {
	return godog.ErrPending
}

func (a *acceptance) thatImageIsConverted(imageType string) error {
	return godog.ErrPending
}

func (a *acceptance) aMustBeReturned(imageType string) error {
	return godog.ErrPending
}

// Integration test level
// Call real urls
type integrated struct {
	test
}

func (i *integrated) aBaseImage(imageType string) error {
	return godog.ErrPending
}

func (i *integrated) thatImageIsConverted(imageType string) error {
	return godog.ErrPending
}

func (i *integrated) aMustBeReturned(imageType string) error {
	return godog.ErrPending
}
