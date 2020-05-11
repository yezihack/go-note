package demo_slice

func getData() []string {
	s := make([]string, 0)
	for i := 0; i <= 1000; i++ {
		s = append(s, "hello")
	}
	return s
}
