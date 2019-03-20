package factory

type student struct {
	Name string
	Age int
	Sex string
}

func newStudent(name string,age int,sex string ) *student {
	return &student{name,age,sex}
}