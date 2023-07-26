package strategy

import "fmt"

type Fifo struct {

}
func(l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}


type LRU struct {

}
func(l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by LRU strtegy")
}

type LFU struct {
	
}
func(l *LFU) evict(c *Cache) {
	fmt.Println("Evicting by LFU strtegy")
}