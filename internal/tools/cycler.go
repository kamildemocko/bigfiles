package tools

type cycler struct {
	elements []string
	index    int
}

func newCycler() *cycler {
	elements := []string{"-", "\\", "|", "/"}

	return &cycler{
		elements: elements,
		index:    0,
	}
}

func (c *cycler) Next() string {
	el := c.elements[c.index]
	c.index = (c.index + 1) % len(c.elements)

	return el
}
