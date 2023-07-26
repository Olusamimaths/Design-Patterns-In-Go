package strategy

func StrategyExample() {
	lfu := &LFU{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")


	cache.add("c", "3")

	lru := &LRU{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}

