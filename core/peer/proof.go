package peer

type AuthorityProver interface {
	Prove(proof []byte) (bool, error)
}
