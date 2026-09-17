package platform_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xtls/xray-core/common"
	. "github.com/xtls/xray-core/common/platform"
)

func TestNormalizeEnvName(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "a",
			output: "A",
		},
		{
			input:  "a.a",
			output: "A_A",
		},
		{
			input:  "A.A.B",
			output: "A_A_B",
		},
	}
	for _, test := range cases {
		if v := NormalizeEnvName(test.input); v != test.output {
			t.Error("unexpected output: ", v, " want ", test.output)
		}
	}
}

func TestEnvFlag(t *testing.T) {
	v := EnvFlag{
		Name: "xxxxx.y",
	}.GetValueAsInt(10)
	if v != 10 {
		t.Error("env value: ", v)
	}
}

func TestAurayLocationEnvironmentNames(t *testing.T) {
	tests := map[string]string{
		AurayConfigLocation:   "AURAY_LOCATION_CONFIG",
		AurayConfdirLocation:  "AURAY_LOCATION_CONFDIR",
		AurayAssetLocation:    "AURAY_LOCATION_ASSET",
		AurayCertLocation:     "AURAY_LOCATION_CERT",
		AurayUseReadV:         "AURAY_BUF_READV",
		AurayUseFreedomSplice: "AURAY_BUF_SPLICE",
		AurayUseVmessPadding:  "AURAY_VMESS_PADDING",
		AurayUseCone:          "AURAY_CONE_DISABLED",
		AurayUseStrictJSON:    "AURAY_JSON_STRICT",
		AurayBufferSize:       "AURAY_RAY_BUFFER_SIZE",
		AurayBrowserDialer:    "AURAY_BROWSER_DIALER",
		AurayXUDPLog:          "AURAY_XUDP_SHOW",
		AurayXUDPBaseKey:      "AURAY_XUDP_BASEKEY",
		AurayTunFdKey:         "AURAY_TUN_FD",
	}
	for input, expected := range tests {
		if value := NormalizeEnvName(input); value != expected {
			t.Errorf("NormalizeEnvName(%q) = %q, want %q", input, value, expected)
		}
	}
}

func TestEnvFlagFallbackPriority(t *testing.T) {
	const (
		primary  = "auray.test.location"
		fallback = "xray.test.location"
	)

	t.Setenv(fallback, "legacy-dotted")
	flag := NewEnvFlagWithFallback(primary, fallback)
	if value := flag.GetValue(func() string { return "default" }); value != "legacy-dotted" {
		t.Fatalf("fallback value = %q, want %q", value, "legacy-dotted")
	}

	t.Setenv(NormalizeEnvName(primary), "auray-normalized")
	if value := flag.GetValue(func() string { return "default" }); value != "auray-normalized" {
		t.Fatalf("normalized primary value = %q, want %q", value, "auray-normalized")
	}

	t.Setenv(primary, "auray-dotted")
	if value := flag.GetValue(func() string { return "default" }); value != "auray-dotted" {
		t.Fatalf("dotted primary value = %q, want %q", value, "auray-dotted")
	}
}

func TestGetAssetLocation(t *testing.T) {
	exec, err := os.Executable()
	common.Must(err)

	loc := GetAssetLocation("t")
	if filepath.Dir(loc) != filepath.Dir(exec) {
		t.Error("asset dir: ", loc, " not in ", exec)
	}

	t.Setenv(AssetLocation, "/xray")
	if v := GetAssetLocation("t"); v != filepath.Join("/xray", "t") {
		t.Error("legacy asset loc: ", v)
	}

	t.Setenv(AurayAssetLocation, "/auray")
	if v := GetAssetLocation("t"); v != filepath.Join("/auray", "t") {
		t.Error("auray asset loc: ", v)
	}
}

func TestAurayConfigurationLocationTakesPriority(t *testing.T) {
	legacyDir := t.TempDir()
	aurayDir := t.TempDir()
	t.Setenv(ConfigLocation, legacyDir)
	t.Setenv(AurayConfigLocation, aurayDir)

	if value := GetConfigurationPath(); value != filepath.Join(aurayDir, "config.json") {
		t.Fatalf("configuration path = %q, want %q", value, filepath.Join(aurayDir, "config.json"))
	}
}
