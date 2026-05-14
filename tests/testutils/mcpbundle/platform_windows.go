// Copyright 2026 The MathWorks, Inc.

//go:build windows

package mcpbundle

import (
	"context"
	"os/exec"
)

const launcherFilename = "launch-matlab-mcp.cmd"
const pathWithSpaces = `C:\Program Files\MATLAB`

func execLauncherCommand(ctx context.Context, launcherPath string) *exec.Cmd {
	return exec.CommandContext(ctx, "cmd", "/c", launcherPath)
}
