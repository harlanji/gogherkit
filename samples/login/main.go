package main

type LoginAPI struct {
	credentials map[string]string
}

func (loginAPI *LoginAPI) Init() {
	loginAPI.credentials = make(map[string]string, 100)
}

func (loginAPI *LoginAPI) Register(username string, password string) bool {
	if _, present := loginAPI.credentials[username]; present {
		return false
	}

	loginAPI.credentials[username] = password

	return true
}

func (loginAPI *LoginAPI) ChangePassword(username string, newPassword string) bool {
	if _, present := loginAPI.credentials[username]; !present {
		return false
	}

	loginAPI.credentials[username] = newPassword

	return true
}

func (loginAPI *LoginAPI) AttemptLogin(username string, password string) bool {
	if pw, present := loginAPI.credentials[username]; !present || pw != password {
		return false
	}

	return true
}
