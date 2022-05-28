package s771

func numJewelsInStones(jewels string, stones string) int {
	var c int
	m := make(map[byte]bool)
	for _, v := range []byte(jewels) {
		m[v] = true
	}
	for _, v := range []byte(stones) {
		_, exists := m[v]
		if exists {
			c++
		}
	}
	return c
}
