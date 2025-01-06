package btree

type BNode struct {
	data []byte
}

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

type BTree struct {
	root uint64

	get func(uint64) BNode
	new func(BNode) uint64
	del func(uint64)
}
