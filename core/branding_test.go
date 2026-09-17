package core

import (
	"strings"
	"testing"
)

func TestAurayVersionStatement(t *testing.T) {
	statement := strings.Join(VersionStatement(), "\n")

	for _, expected := range []string{
		DistributionName(),
		DistributionVersion(),
	} {
		if !strings.Contains(statement, expected) {
			t.Fatalf("version statement %q does not contain %q", statement, expected)
		}
	}
	if strings.Contains(strings.ToLower(statement), "xray") {
		t.Fatalf("version statement %q unexpectedly exposes the upstream brand", statement)
	}
}
