package main

import "jsonerator/Jsonerator"

func main() {
	//"{\"encryptionMethod\":\"xtsAes256\",\"requireEncryptionForWriteAccess\":true,\"blockCrossOrganizationWriteAccess\":false}"
	//"{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"
	Jsonerator.GetKeyVals("{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}")
}
