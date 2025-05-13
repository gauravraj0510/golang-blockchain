
package blockchain

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock checks if the transaction input can be unlocked with the given data.
// It will return true if the signature matches the data, and false otherwise.
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked checks if the output can be unlocked with the given data.
// It will return true if the PubKey matches the data, and false otherwise.
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}