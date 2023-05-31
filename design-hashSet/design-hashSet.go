package designhashset

type MyHashSet struct {
	Mp map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{Mp: make(map[int]bool)}
}

func (this *MyHashSet) Add(key int) {
	this.Mp[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.Mp[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.Mp[key]
}
