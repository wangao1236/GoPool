package util

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Stdout(cmd *exec.Cmd) error {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Printf("Exec command stdout\n%s", string(output))
	return nil
}

func Exec(cmdStr string) {
	argv := strings.Fields(cmdStr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, argv[0], argv[1:]...)
	err := Stdout(cmd)

	select {
	case <-ctx.Done():
		log.Printf("Exec command %s timeout", cmdStr)
	default:
		if err != nil {
			log.Printf("Exec command %s failed, err: %v", cmdStr, err)
		} else {
			log.Printf("Exec command %s succeed", cmdStr)
		}
	}
}
