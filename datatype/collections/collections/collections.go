package collections

type WorkWith struct {
	Data    string
	Version int
}

func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	res := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			res = append(res, w)
		}
	}
	return res
}

func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	res := make([]WorkWith, len(ws))
	for pos, w := range ws {
		newW := f(w)
		res[pos] = newW
	}
	return res
}
