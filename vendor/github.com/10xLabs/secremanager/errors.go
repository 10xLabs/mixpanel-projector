package secremanager

// NilSecretStringError ...
type NilSecretStringError struct{}

func (e NilSecretStringError) Error() string {
	return "The decrypted part of the protected secret information that you asked for is nil"
}
