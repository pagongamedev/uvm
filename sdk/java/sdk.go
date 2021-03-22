package java

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderOracle(sPlatform)
	case "darwin":
		provider, _ = NewProviderOracle(sPlatform)
	case "linux":
		provider, _ = NewProviderOracle(sPlatform)
	default:
		provider, _ = NewProviderOracle(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:     "Java",
		LinkName: "Java",
		Command:  "-oj",
		// Env:        "JAVA_HOME",
		EnvBin:     "bin",
		EnvChannel: "UVM_JAVA_CHANNEL",
		Provider:   provider,
	}, nil
}
