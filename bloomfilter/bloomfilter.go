package bloomfilter

//https://www.cnblogs.com/Hollson/p/12031692.html
import "github.com/willf/bitset"

type BloomFilter struct {
	Set *bitset.BitSet
}
