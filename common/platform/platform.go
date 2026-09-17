package platform // import "github.com/xtls/xray-core/common/platform"

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	AurayConfigLocation   = "auray.location.config"
	AurayConfdirLocation  = "auray.location.confdir"
	AurayAssetLocation    = "auray.location.asset"
	AurayCertLocation     = "auray.location.cert"
	AurayUseReadV         = "auray.buf.readv"
	AurayUseFreedomSplice = "auray.buf.splice"
	AurayUseVmessPadding  = "auray.vmess.padding"
	AurayUseCone          = "auray.cone.disabled"
	AurayUseStrictJSON    = "auray.json.strict"
	AurayBufferSize       = "auray.ray.buffer.size"
	AurayBrowserDialer    = "auray.browser.dialer"
	AurayXUDPLog          = "auray.xudp.show"
	AurayXUDPBaseKey      = "auray.xudp.basekey"
	AurayTunFdKey         = "auray.tun.fd"

	// Legacy Xray-core names remain supported for compatibility.
	ConfigLocation  = "xray.location.config"
	ConfdirLocation = "xray.location.confdir"
	AssetLocation   = "xray.location.asset"
	CertLocation    = "xray.location.cert"

	UseReadV         = "xray.buf.readv"
	UseFreedomSplice = "xray.buf.splice"
	UseVmessPadding  = "xray.vmess.padding"
	UseCone          = "xray.cone.disabled"
	UseStrictJSON    = "xray.json.strict"

	BufferSize           = "xray.ray.buffer.size"
	BrowserDialerAddress = "xray.browser.dialer"
	XUDPLog              = "xray.xudp.show"
	XUDPBaseKey          = "xray.xudp.basekey"

	TunFdKey = "xray.tun.fd"
)

type EnvFlag struct {
	Name          string
	AltName       string
	fallbackNames []string
}

func NewEnvFlag(name string) EnvFlag {
	return EnvFlag{
		Name:    name,
		AltName: NormalizeEnvName(name),
	}
}

// NewEnvFlagWithFallback returns an environment flag that checks name first,
// followed by fallbackNames. Both dotted and normalized uppercase forms are
// accepted for every name.
func NewEnvFlagWithFallback(name string, fallbackNames ...string) EnvFlag {
	flag := NewEnvFlag(name)
	flag.fallbackNames = append(flag.fallbackNames, fallbackNames...)
	return flag
}

func lookupEnv(name string) (string, bool) {
	if len(name) == 0 {
		return "", false
	}
	return os.LookupEnv(name)
}

func (f EnvFlag) GetValue(defaultValue func() string) string {
	if v, found := lookupEnv(f.Name); found {
		return v
	}
	if v, found := lookupEnv(f.AltName); found {
		return v
	}
	for _, name := range f.fallbackNames {
		if v, found := lookupEnv(name); found {
			return v
		}
		if v, found := lookupEnv(NormalizeEnvName(name)); found {
			return v
		}
	}

	return defaultValue()
}

func (f EnvFlag) GetValueAsInt(defaultValue int) int {
	useDefaultValue := false
	s := f.GetValue(func() string {
		useDefaultValue = true
		return ""
	})
	if useDefaultValue {
		return defaultValue
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int(v)
}

func NormalizeEnvName(name string) string {
	return strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(name)), ".", "_")
}

func getExecutableDir() string {
	exec, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exec)
}

func GetConfigurationPath() string {
	configPath := NewEnvFlagWithFallback(AurayConfigLocation, ConfigLocation).GetValue(getExecutableDir)
	return filepath.Join(configPath, "config.json")
}

// GetConfDirPath reads Auray's confdir location, falling back to Xray-core's
// legacy environment variable.
func GetConfDirPath() string {
	configPath := NewEnvFlagWithFallback(AurayConfdirLocation, ConfdirLocation).GetValue(func() string { return "" })
	return configPath
}
