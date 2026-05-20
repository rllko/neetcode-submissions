type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		fmt.Fprintf(&builder, "%d#%s,", len(str), str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
    final := []string{}
	i := 0

	for i < len(encoded) {
		hashPos := strings.Index(encoded[i:], "#")
		if hashPos == -1 {
			break
		}

		hashPos += i

		num, _ := strconv.Atoi(encoded[i:hashPos])
		i = hashPos + 1

		substr := encoded[i : i+num]
		final = append(final, substr)
		i += num + 1
	}

	return final
}
