package main

import "fmt"

type Set map[string]bool

func NewSet() Set {
	s := make(Set)
	return s
}

func (s Set) Add(key string) {
	s[key] = true
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func (s Set) Size() int {
	return len(s)
}
func (s Set) Exists(key string) bool {
	return s[key]
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()
	for _, str := range arr {
		set.Add(str)
	}
	fmt.Println(set)
}
