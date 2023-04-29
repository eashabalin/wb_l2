package pattern

import "fmt"

// Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов
// и помещает каждый из них в собственный класс, после чего алгоритмы
// можно взаимозаменять прямо во время исполнения программы.

// Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу. Программа может подменить
// этот объект другим, если требуется иной способ решения задачи.

type EvictionAlgorithm interface {
	evict(c *Cache)
}

type Fifo struct{}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

type Cache struct {
	evictionAlgorithm EvictionAlgorithm
	cap               int
	maxCap            int
}

func NewCache(evictionAlgorithm EvictionAlgorithm) *Cache {
	return &Cache{evictionAlgorithm, 0, 3}
}

func (c *Cache) evict() {
	c.evictionAlgorithm.evict(c)
	c.cap--
}

func (c *Cache) Add() {
	if c.cap == c.maxCap {
		c.evict()
	}
	c.cap++
	fmt.Println("Added", c.cap)
}

func Run07() {
	lfu := &Lfu{}
	cache := NewCache(lfu)

	cache.Add()
	cache.Add()
	cache.Add()
	cache.Add()
}
