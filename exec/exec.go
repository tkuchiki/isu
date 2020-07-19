package exec

import (
	"context"
	"fmt"
	exe "os/exec"
)

func Command(ctx context.Context, cmd string) error {
	return exe.CommandContext(ctx, "sh", "-c", cmd).Run()
}

func CommandWithGzip(ctx context.Context, cmd, filename string) error {
	command := fmt.Sprintf(cmd+`| gzip -9 > %s.gz`, filename)
	return exe.CommandContext(ctx, "sh", "-c", command).Run()
}
