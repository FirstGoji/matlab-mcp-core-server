// Copyright 2026 The MathWorks, Inc.

//go:build !windows

package mcpbundle

import (
	"context"
	"os/exec"
)

const launcherFilename = "launch-matlab-mcp.sh"
const pathWithSpaces = "/opt/my matlab/R2025b"

func execLauncherCommand(ctx context.Context, launcherPath string) *exec.Cmd {
	return exec.CommandContext(ctx, "bash", launcherPath)
}
