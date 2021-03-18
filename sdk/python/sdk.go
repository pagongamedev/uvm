package python

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderPythonOrg(sPlatform)
	case "darwin":
		provider, _ = NewProviderPythonOrg(sPlatform)
	case "linux":
		provider, _ = NewProviderPythonOrg(sPlatform)
	default:
		provider, _ = NewProviderPythonOrg(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Python",
		LinkName:   "Python",
		Command:    "-p",
		Env:        "",
		EnvBin:     "",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
