/*
Custom test utils for running tests with expected exit errors
*/

package testUtil

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func failWithTimeoutErr(t *testing.T, cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		t.Fatal("Failed to kill process: ", err)
	}
	t.Fatal("Process did not exit within expected timeout")
}

func verifyProcessErr(t *testing.T, err error, expectedErrMsg string) {
	if err != nil {
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Fatalf("Expected Process to exit with err message %s, but received err %v instead", expectedErrMsg, err)
		}
	} else {
		t.Fatal("Process did not exit with expected error, returned silently")
	}
}

func WaitAndVerifyErr(t *testing.T, cmd *exec.Cmd, timeoutPeriod time.Duration, expectedErrMsg string) {
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeoutPeriod * time.Second):
		failWithTimeoutErr(t, cmd)
	case err := <-done:
		verifyProcessErr(t, err, expectedErrMsg)
	}
}

func RunTestWithExpectedError(t *testing.T, testName string, expectedErrMsg string, timeoutPeriod time.Duration) {
	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start test in crasher mode, err: ", err)
	}
	WaitAndVerifyErr(t, cmd, timeoutPeriod, expectedErrMsg)
}
