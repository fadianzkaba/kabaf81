package main

type test123 struct {
	s1 []int
	s2 []interface{}
}

type Config struct {
	TaskQueue string
}

func New(config Config) *test123 {
	return &test123{
		s1: make([]int, 3),
		s2: make([]interface{}, 1),
	}
}

func main() {

	s00 := test123{}

	s00.s2 = append(s00.s2, addTest123())

}

func addTest123(s string) []string {
	s8 := []string{"Hello", "World"}
	return s8

}
func addTest1234() []string {
	s8 := []string{"Hello", "World"}
	return s8

}
