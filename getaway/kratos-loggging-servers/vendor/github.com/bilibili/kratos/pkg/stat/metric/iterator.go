package metric

import "fmt"

// Iterator iterates the buckets within the window.
type Iterator struct {
	count         int
	iteratedCount int
	cur           *Bucket
}

// Next returns true util all of the buckets has been iterated.
func (i *Iterator) Next() bool {
	return i.count != i.iteratedCount
}

// Bucket gets current bucket.
func (i *Iterator) Bucket() Bucket {
	if !(i.Next()) {
		panic(fmt.Errorf("stat/metric: iteration out of range iteratedCount: %d count: %d", i.iteratedCount, i.count))
	}
	bucket := *i.cur
	i.iteratedCount++
	i.cur = i.cur.Next()
	return bucket
}

func (i *Iterator)Print()  {
	fmt.Println("....Iterator ...")
	fmt.Println("Iterator count",i.count)
	fmt.Println("Iterator cur=",i.cur)
	fmt.Println("Iterator iteratedCount=",i.iteratedCount)
}