package funcs

type Decision struct {
	Id     int
	Result string
}

var Decisions_1 = []Decision{
	{Id: 1, Result: "approved"},
	{Id: 3, Result: "approved"},
	{Id: 15, Result: "approved"},
	{Id: 20, Result: "approved"},
	{Id: 26, Result: "approved"},
	{Id: 30, Result: "approved"},
}

var Decisions_2 = []Decision{
	{Id: 2, Result: "approved"},
	{Id: 4, Result: "approved"},
	{Id: 14, Result: "approved"},
	{Id: 16, Result: "approved"},
	{Id: 23, Result: "approved"},
	{Id: 32, Result: "approved"},
}

func GetLastDecisions(d1 []Decision, d2 []Decision, k int) []Decision {
	d1Ptr := len(d1) - 1
	d2Ptr := len(d2) - 1

	count := 0
	res := make([]Decision, k)

	for count < k {
		d1EndEl := d1[d1Ptr]
		d2EndEl := d2[d2Ptr]

		if d1EndEl.Id > d2EndEl.Id {
			res[count] = d1EndEl
			d1Ptr--
		} else {
			res[count] = d2EndEl
			d2Ptr--
		}

		count++
	}

	return res
}
