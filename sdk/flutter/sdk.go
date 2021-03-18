package flutter

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderFlutterDev(sPlatform)
	case "darwin":
		provider, _ = NewProviderFlutterDev(sPlatform)
	case "linux":
		provider, _ = NewProviderFlutterDev(sPlatform)
	default:
		provider, _ = NewProviderFlutterDev(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Flutter",
		LinkName:   "Flutter",
		Command:    "-f",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
