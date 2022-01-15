package gondlr

type Signer interface {
	PublicKey() []byte
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) bool
}
