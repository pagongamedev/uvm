package nodejs

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	envBin := ""
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderNodejsOrg(sPlatform)
	case "darwin":
		provider, _ = NewProviderNodejsOrg(sPlatform)
		envBin = "bin"
	case "linux":
		provider, _ = NewProviderNodejsOrg(sPlatform)
		envBin = "bin"
	default:
		provider, _ = NewProviderNodejsOrg(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "NodeJS",
		LinkName:   "NodeJS",
		Command:    "-n",
		Env:        "",
		EnvBin:     envBin,
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
