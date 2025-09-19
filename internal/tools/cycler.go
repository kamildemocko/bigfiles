package tools

type Cycler struct {
	elements []string
	index    int
}

func NewCycler() Cycler {
	elements := []string{"-", "\\", "|", "/"}

	return Cycler{
		elements: elements,
		index:    0,
	}
}

func (c *Cycler) Next() string {
	el := c.elements[c.index]
	c.index = (c.index + 1) % len(c.elements)

	return el
}
