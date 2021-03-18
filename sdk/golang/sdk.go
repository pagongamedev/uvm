package golang

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderGolangOrg(sPlatform)
	case "darwin":
		provider, _ = NewProviderGolangOrg(sPlatform)
	case "linux":
		provider, _ = NewProviderGolangOrg(sPlatform)
	default:
		provider, _ = NewProviderGolangOrg(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Golang",
		LinkName:   "Golang",
		Command:    "-g",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
