package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
- Given a list of [FileName, FileSize, [Collection]] - Collection is optional, i.e., a collection can have 1 or more files. Same file can be a part of more than 1 collection.How would you design a system
    - [https://leetcode.com/discuss/interview-question/1508471/code-design-file-processing-scenario](https://leetcode.com/discuss/interview-question/1508471/code-design-file-processing-scenario)
    - To calculate total size of files processed.
    - To calculate Top K collections based on size.Example:
    - file1.txt(size: 100)
    file2.txt(size: 200) in collection "collection1"
    file3.txt(size: 200) in collection "collection1"
    file4.txt(size: 300) in collection "collection2"
    file5.txt(size: 100)
    - Output:
    - Total size of files processed: 900
    Top 2 collections:
    - collection1 : 400
    - collection2 : 300
*/

type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File {
	return &File{name, size}
}

type Collection struct {
	name      string
	index     int
	files     []*File
	totalSize int64
}

func NewCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}

func (c *Collection) TotalSize() int64 {
	return c.totalSize
}

func (c *Collection) AddFile(f *File) {
	c.files = append(c.files, f)
	c.totalSize += f.size
}

type Catalog struct {
	files       []*File
	collections []*Collection
	catalog     map[string]*Collection
}

func NewCatalog() *Catalog {
	c := &Catalog{
		catalog: map[string]*Collection{},
	}
	return c
}

func (c *Catalog) Add(data [][]string) {
	var filename string
	var size int64
	for _, row := range data {
		filename = row[0]
		size = parseInt(row[1])
		file := NewFile(filename, size)
		c.files = append(c.files, file)

		if len(row) == 3 {
			colName := row[2]
			c.addFileToCollection(colName, file)
		}
	}
}

func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (c *Catalog) addFileToCollection(name string, f *File) *Collection {
	if _, ok := c.catalog[name]; !ok {
		col := &Collection{
			name:  name,
			index: len(c.collections),
		}
		col.AddFile(f)
		c.collections = append(c.collections, col)
		c.catalog[name] = col
		return col
	}

	col := c.catalog[name]
	col.AddFile(f)
	return col
}

func (c *Catalog) TotalSizeOfFiles() int64 {
	var total int64
	for _, f := range c.files {
		total += f.size
	}
	return total
}

func (c *Catalog) TopKCollections(k int) map[string]int64 {
	if len(c.collections) < k {
		return nil
	}
	colHeap := CollectionsHeap(c.collections)
	heap.Init(&colHeap)

	res := map[string]int64{}
	for i := 0; i < k; i++ {
		col := heap.Pop(&colHeap).(*Collection)
		res[col.name] = col.TotalSize()
	}
	return res
}

type CollectionsHeap []*Collection

// Len is the number of elements in the collection.
func (c *CollectionsHeap) Len() int {
	return len(*c)
}

func (c *CollectionsHeap) Less(i int, j int) bool {
	return (*c)[i].TotalSize() > (*c)[j].TotalSize()
}

// Swap swaps the elements with indexes i and j.
func (c *CollectionsHeap) Swap(i int, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
	(*c)[i].index = j
	(*c)[j].index = i
}

func (c *CollectionsHeap) Push(x any) {
	col := x.(*Collection)
	*c = append(*c, col)
}

func (c *CollectionsHeap) Pop() any {
	if c.Len() > 0 {
		old := *c
		last := old[c.Len()-1]
		old[c.Len()-1] = nil
		*c = old[:c.Len()-1]
		return last
	}
	return nil
}
