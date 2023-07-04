package cage

func HashPassword(password string) (string, error) {
	params := Params{65536, 8, 1, 16, 64}
	hashed, err := GenerateFromPassword([]byte(password), params)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func VerifyPassword(password, hash string) error {
	if err := CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}
