package concat

func Concat(strings []string) string {
	var s, sep string
	for _, str := range strings {
		s += sep + str
		sep = " "
	}
	return s
}
