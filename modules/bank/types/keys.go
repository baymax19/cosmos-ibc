package types

func UserKey(chainID string) []byte {
	return []byte("users/" + chainID)
}
