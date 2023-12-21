package piscine

func PodiumPosition(podium [][]string) [][]string {
	for x := 0; x < len(podium)/2; x++ {
		y := len(podium) - x - 1
		podium[x], podium[y] = podium[y], podium[x]
	}
	return podium
}
