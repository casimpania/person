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

func (p person) ShowFirstName() string {
        return p.firstName
}

func (p person) ShowLastName() string {
        return p.lastName
}

func (p person) ShowAge() int {
        return p.age
}
