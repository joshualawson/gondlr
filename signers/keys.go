package signers

type PublicKey []byte

func (p PublicKey) String() string {
	return string(p)
}

func (p PublicKey) Byte() []byte {
	return p
}

type PrivateKey []byte

func (p PrivateKey) String() string {
	return string(p)
}

func (p PrivateKey) Byte() []byte {
	return p
}
