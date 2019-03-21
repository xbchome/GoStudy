package bin

import "fmt"

type Pupil struct {
	Student
}

func (pup *Pupil) Exam()  {
	fmt.Println("小学生在考试")
}
