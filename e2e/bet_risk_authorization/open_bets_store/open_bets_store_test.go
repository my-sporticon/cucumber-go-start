package open_bets_store

import (
	"github.com/cucumber/godog"
	"github.com/my-sporticon/cucumber-go-start/internal/steps"
	"testing"
)

func aHungryCat() error {
	//return godog.ErrPending
	steps.Bar()
	return nil
}

func aSatiatedCat() error {
	//return godog.ErrPending
	return nil
}

func iFeedTheCat() error {
	//return godog.ErrPending
	return nil
}

func theCatIsNotHungry() error {
	//return godog.ErrPending
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a hungry cat$`, aHungryCat)
	ctx.Step(`^a satiated cat$`, aSatiatedCat)
	ctx.Step(`^I feed the cat$`, iFeedTheCat)
	ctx.Step(`^the cat is not hungry$`, theCatIsNotHungry)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
