package Common

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func ExecCmd(cmd string, timeout ...time.Duration) ([]byte, error) {
	var t time.Duration
	if len(timeout) > 0 {
		t = timeout[0]
	}
	var ec *exec.Cmd
	if t > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), timeout[0])
		defer cancel()
		ec = exec.CommandContext(ctx, "bash", "-c", cmd)
	} else {
		ec = exec.Command("bash", "-c", cmd)
	}
	result, err := ec.Output()
	log.Panicf("result:%s", string(result))
	//log.Debug(zap.String("cmd", cmd), zap.String("result", string(result)))
	return result, err
}