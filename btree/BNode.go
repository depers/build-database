package btree

import "encoding/binary"

type BNode struct {
	data []byte
}

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

// 一个节点存储为字节数组，存储的主要内容是：
// 2个字节：用于存储节点类型，是内部节点还是叶子节点
// 2个字节：用于存储节点存储的键的数量
// 键的数量 * 8个字节：用于存储保存指向子节点的指针列表
// 键的数量 * 2个字节：用于存储指向每个键值对的偏移量列表
// 2个字节：用于存储key
// 2个字节：用于存储value
// 2个字节：用于存储key
// 2个字节：用于存储value
// 2个字节：用于存储key
// 2个字节：用于存储value
//.....

// 节点头的数据是4个字节，分别是节点类型和键的数量
const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	assert(node1max <= BTREE_PAGE_SIZE)
}

// header
// 获取节点的类型
func (node BNode) btype() uint16 {
	// 从字节切片中读取一个16位的无符号整数（uint16），并假设该整数是以小端字节序存储的
	// 我的理解是读取了前两个字节，也就是节点的类型
	return binary.LittleEndian.Uint16(node.data)
}

// 读取节点存储的键的数量
func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}

// 设置节点类型和键的数量
func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

// pointers
// 获取执行子节点的指针
func (node BNode) getPtr(idx uint16) uint64 {
	assert(idx < node.nkeys())
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint64(node.data[pos:])
}

// 保存指向子节点的指针
func (node BNode) setPtr(idx uint16, val uint64) {
	assert(idx < node.nkeys())
	pos := HEADER + 8*idx
	binary.LittleEndian.PutUint64(node.data[pos:], val)
}

// offset list
// 获取第i-1个键的偏移下标，值得注意的是该偏移量是相对于第一个 KV 对的位置而言的。 第一个 KV 对的偏移量始终为零，因此不存储在列表中。
func offsetPos(node BNode, idx uint16) uint16 {
	assert(1 <= idx && idx <= node.nkeys())
	return HEADER + 8*node.nkeys() + 2*(idx-1)
}

// 获取第i-1个键的偏移量
func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(node.data[offsetPos(node, idx):])
}

// 设置键的偏移量
func (node BNode) setOffset(idx uint16, offset uint16) {
	binary.LittleEndian.PutUint16(node.data[offsetPos(node, idx):], offset)
}

// key-values
func (node BNode) kvPos(idx uint16) uint16 {
	assert(idx <= node.nkeys())
	return HEADER + 8*node.nkeys() + 2*node.nkeys() + node.getOffset(idx)
}
func (node BNode) getKey(idx uint16) []byte {
	assert(idx < node.nkeys())
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	return node.data[pos+4:][:klen]
}
func (node BNode) getVal(idx uint16) []byte {
	assert(idx < node.nkeys())
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos+0:])
	vlen := binary.LittleEndian.Uint16(node.data[pos+2:])
	return node.data[pos+4+klen:][:vlen]
}

func assert(b bool) {

}
