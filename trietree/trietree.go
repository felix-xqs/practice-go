package trietree

//核心思想，空间换时间，利用字符串公共前缀来降低查询时间的开销提高效率
//根节点不包括字符，路径上字符连接起来，为对应字符串
//每个节点的子节点包含字符都不同

//搜索方法
//从根结点开始第一次搜索
//取得查找关键词的第一个字幕，并根据字幕选择对应的子树，并转到该子树继续进行检索
//响应的子树上，取得要查找关键词的第二个字母，并进一步选择对应的子树进行检索
//迭代
//某个结点出，关键词所有字母都已取出，则读取附在该结点上的信息完成操作

const MAXCAP = 26

type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

//Initialize
func Constructor() Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie, MAXCAP)
	root.isWord = false
	return *root
}

//Insert
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v] == nil {
			node := new(Trie)
			node.next = make(map[rune]*Trie, MAXCAP)
			node.isWord = false
			this.next[v] = node
		}
		this = this.next[v]
	}
	this.isWord = true
}

//Search
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return true
}
