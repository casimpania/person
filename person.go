package person

type person struct {
        firstName string
        lastName  string
        age       int
}

func New(firstName string, lastName string, age int) person {
        p := person{
                firstName: firstName,
                lastName:  lastName,
                age:       age,
        }

        return p
}

func (p person) GetFirstName() string {
        return p.firstName
}

func (p person) GetLastName() string {
        return p.lastName
}

func (p person) GetAge() int {
        return p.age
}
