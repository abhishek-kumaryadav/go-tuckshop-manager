package services

func Configure(key string, value string) {
	envMap = GetEnvProperties()
	isKeyPresent := false
	for k, _ := range envMap {
		if k == key {
			envMap[k] = value
			isKeyPresent = true
		}
	}
	if !isKeyPresent {
		envMap[key] = value
	}
	UpdateEnvFile(envMap)
}

func Delete(key string) {
	envMap = GetEnvProperties()
	delete(envMap, key)
	UpdateEnvFile(envMap)
}
