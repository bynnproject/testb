package main

func main() {

}

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcb(x, y) == 0
}

func gcb(x, y int) int {
	if x > y {
		x, y = y, x
	}
	tmp := y % x
	for tmp != 0 {
		y = x
		x = tmp % x
		tmp = y % x
	}
	return x
}
