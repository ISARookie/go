package person

type Person struct {
	firstName string
	lastName  string
}

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
