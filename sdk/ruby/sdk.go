package ruby

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderRubyinstallerOrg(sPlatform)
	// case "darwin":
	// case "linux":
	default:
		provider, _ = NewProviderRubyinstallerOrg(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Ruby",
		LinkName:   "Ruby",
		Command:    "-r",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
