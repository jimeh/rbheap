package main

// diffsect removes all items in `a` from `b`, then removes all items from `b`
// which are not in `c`. Effectively: intersect(difference(b, a), c)
func diffsect(a, b, c *[]string) *[]string {
	result := []string{}
	mapA := map[string]bool{}
	mapC := map[string]bool{}

	for _, x := range *a {
		mapA[x] = true
	}

	for _, x := range *c {
		mapC[x] = true
	}

	for _, x := range *b {
		_, okA := mapA[x]
		_, okC := mapC[x]

		if !okA && okC {
			result = append(result, x)
		}
	}

	return &result
}
