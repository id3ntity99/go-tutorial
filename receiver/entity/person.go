package person

/*
Learning difference between value receiver and pointer receiver.
*/
type Person struct {
	name, gender string
	age          int
}

// A constructor for Person
func NewPerson(name string, gender string, age int) *Person {
	return &Person{
		name:   name,
		gender: gender,
		age:    age,
	}
}

func (person *Person) GetName() string {
	return person.name
}

func (person *Person) GetGender() string {
	return person.gender
}

func (person *Person) GetAge() int {
	return person.age
}

func (person *Person) GetOlderUsingPointerRecv(n int) int {
	person.age += n
	return person.age
}

func (person Person) GetOlderUsingValueRecv(n int) int {
	person.age += n
	return person.age
}
