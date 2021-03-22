package openjava

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderJavaNet(sPlatform)
	case "darwin":
		provider, _ = NewProviderJavaNet(sPlatform)
	case "linux":
		provider, _ = NewProviderJavaNet(sPlatform)
	default:
		provider, _ = NewProviderJavaNet(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:     "OpenJava",
		LinkName: "Java",
		Command:  "-oj",
		// Env:        "JAVA_HOME",
		EnvBin:     "bin",
		EnvChannel: "UVM_JAVA_CHANNEL",
		Provider:   provider,
	}, nil
}
