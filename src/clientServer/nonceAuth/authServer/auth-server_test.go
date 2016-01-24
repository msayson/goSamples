package authServer

import (
	"os"
	"testUtil" //Custom test utils for running tests with expected exit errors
	"testing"
	"time"
)

func runTestWithExpectedExitErr(t *testing.T, testName string) {
	var timeoutPeriod time.Duration = 3 //timeout before kill process (seconds)
	testUtil.RunTestWithExpectedError(t, testName, "exit status", timeoutPeriod)
}

func TestRunAuthServerWithInvalidIp(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		RunAuthServer("192.1.1.1.1:1234", 123456)
	} else {
		runTestWithExpectedExitErr(t, "TestRunAuthServerWithInvalidIp")
	}
}

func TestRunAuthServerWithInvalidPort(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		RunAuthServer("localhost:65536", 123456)
	} else {
		runTestWithExpectedExitErr(t, "TestRunAuthServerWithInvalidPort")
	}
}
