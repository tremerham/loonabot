package runner

import (
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestService(t *testing.T) {
	go Start("localhost:5252", nil)
	time.Sleep(time.Second)
	dial, err := grpc.Dial("localhost:5252", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.FailNow()
	}
	defer dial.Close()
	err = UploadExec(dial, "./repl")
	if err != nil {
		t.FailNow()
	}
	err = Restart(dial)
	if err != nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 2)
}
