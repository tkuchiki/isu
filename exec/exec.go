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
	return string(out), err
}

func CommandWithGzip(ctx context.Context, cmd, filename string) error {
	command := fmt.Sprintf(cmd+`| gzip -9 > %s.gz`, filename)
	return exe.CommandContext(ctx, "sh", "-c", command).Run()
}
