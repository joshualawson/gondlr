package gondlr

type Signer interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) bool
}
