package btree

type BTree struct {
	// 指针（非0页码）
	root uint64

	// 用于管理磁盘页面的回调
	get func(uint64) BNode // 取消引用指针
	new func(BNode) uint64 // 分配一个新页面
	del func(uint64)       // 取消分配页面
}
