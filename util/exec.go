package util

import (
	"context"
	"log"
	"os/exec"
	"regexp"
	"time"
)

func Stdout(cmd *exec.Cmd) error {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Print("Exec command stdout\n" + string(output))
	return nil
}

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	lastStart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[lastStart:element[0]]
		lastStart = element[1]
	}
	result[len(indexes)] = text[lastStart:]
	return result
}

func Exec(cmdStr string) {
	argv := RegSplit(cmdStr, "\\s+")

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
