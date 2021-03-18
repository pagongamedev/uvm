package dart

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderDartDev(sPlatform)
	case "darwin":
		provider, _ = NewProviderDartDev(sPlatform)
	case "linux":
		provider, _ = NewProviderDartDev(sPlatform)
	default:
		provider, _ = NewProviderDartDev(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Dart",
		LinkName:   "Dart",
		Command:    "-d",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}
