package base

import "testing"

func TestAurayExecutableName(t *testing.T) {
	if CommandEnv.Exec != "auray" {
		t.Fatalf("CommandEnv.Exec = %q, want %q", CommandEnv.Exec, "auray")
	}
}
