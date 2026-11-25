package main

import (
	"os/exec"
	"testing"
)

func probe(t *testing.T, name string, args ...string) {
	t.Helper()
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		t.Logf("$ %s %v -> unavailable (%v)", name, args, err)
		return
	}
	t.Logf("$ %s %v\n%s", name, args, string(out))
}

func TestAcceleratorInventory(t *testing.T) {
	t.Log("=== accelerator self-check ===")
	probe(t, "nvidia-smi", "-L")
	probe(t, "rocm-smi")
	probe(t, "hostname")
	probe(t, "id")
	t.Log("=== end accelerator self-check ===")
}
