package player

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Play(targetPaths []string, extraFlags ...string) error {
	if len(targetPaths) == 0 {
		return fmt.Errorf("No target paths provided to player.")
	}

	args := append([]string{"--no-terminal"}, extraFlags...)
	args = append(args, targetPaths...)

	player := exec.Command("mpv", args...)

	player.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	if err := player.Start(); err != nil {
		return fmt.Errorf("Failed to start player: %w", err)
	}

	os.Exit(0)
	return nil
}
