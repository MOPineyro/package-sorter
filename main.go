package main

func sort(width, height, length, mass float64) string {
	volume := width * height * length
	isBulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20

	if isBulky && isHeavy {
		return "REJECTED"
	}
	if isBulky || isHeavy {
		return "SPECIAL"
	}
	return "STANDARD"
}

func main() {
	_ = sort(10, 10, 10, 5)
}
