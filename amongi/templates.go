package amongi

var Amongi = []AmongUs{
	{[][]pixel{
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{true, false, true},
	}},
}

type AmongUs struct {
	pixels [][]pixel
}

type pixel bool
