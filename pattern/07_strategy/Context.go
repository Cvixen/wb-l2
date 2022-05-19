package main

type Context struct {
	strategy StrategySort
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
