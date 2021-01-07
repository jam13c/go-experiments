package main

func findTimestamp(buses map[int]int) int {
	t := buses[0]
	d := buses[0]

	for dt, bus := range buses {
		for {
			if (t+dt)%bus == 0 {
				break
			}
			t += d
		}
		d = lcm(d, bus)
	}

	return t
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	return a * b / gcd(a, b)
}
