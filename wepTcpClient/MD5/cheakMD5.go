package MD5

func CheakMD5(server string, client string) bool {

	if server == client {
		return true
	}

	return false

}
