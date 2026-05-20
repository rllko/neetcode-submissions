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
        // see how we find for the index starting at i? 
        // it gets 5#world, no matter how far it is into the str it is
		hashPos := strings.Index(encoded[i:], "#")
        // check if the hashtag exists
		if hashPos == -1 {
			break
		}

        // place the index correctly from the number
		hashPos += i

        // get the number
		num, _ := strconv.Atoi(encoded[i:hashPos])
        // go to hashag + 1 or 5#HERE->world
		i = hashPos + 1

        // get the string
		substr := encoded[i : i+num]
		final = append(final, substr)
        
        // here we have to add one because if it stops before the ,
        // when it tries to to Atoi it will read ,5 and thats invalid
		i += num + 1
	}

	return final
}
