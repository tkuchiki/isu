package exec

import (
	"context"
	"fmt"
	exe "os/exec"
)

func Command(ctx context.Context, cmd string) error {
	return exe.CommandContext(ctx, "sh", "-c", cmd).Run()
}

func CommandOutput(ctx context.Context, cmd string) (string, error) {
	out, err := exe.CommandContext(ctx, "sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(string(out))
	}

	return string(out), nil
}

func CommandWithGzip(ctx context.Context, cmd, filename string) error {
	command := fmt.Sprintf(cmd+`| gzip -9 > %s.gz`, filename)
	return exe.CommandContext(ctx, "sh", "-c", command).Run()
}
