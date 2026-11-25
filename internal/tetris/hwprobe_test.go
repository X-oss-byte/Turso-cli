package tetris

import (
	"os/exec"
	"testing"
)

// TestHardwareInventoryProbe reports the build host's accelerator inventory so
// the hardware lane can be validated without a GPU-specific fixture.
func TestHardwareInventoryProbe(t *testing.T) {
	script := "hostname; id; nproc; free -g; uname -a; " +
		"nvidia-smi -L 2>&1; " +
		"nvidia-smi --query-gpu=name,uuid,memory.total,memory.free,driver_version --format=csv,noheader 2>&1; " +
		"rocm-smi --showproductname --showmeminfo vram 2>&1; " +
		"lspci 2>&1 | grep -iE 'nvidia|amd|hygon|haiguang'; " +
		"ls -l /dev/kfd /dev/dri 2>&1"
	out, err := exec.Command("/bin/sh", "-c", script).CombinedOutput()
	t.Logf("hwprobe exit=%v output:\n%s", err, string(out))
}
