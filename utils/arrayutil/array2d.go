package arrayutil

type Array2D struct {
	data   []int
	width  int
	height int
	length int
}

func NewArray2D(width int, height int, defaultValue int) *Array2D {
	var a = &Array2D{
		width:  width,
		height: height,
	}

	var l = width * height
	a.length = l
	a.data = make([]int, l)
	if defaultValue != 0 {
		for i := 0; i < l; i++ {
			a.data[i] = defaultValue
		}
	}

	return a
}

func (a2d *Array2D) GetIndex(col int, row int) int {
	return row*a2d.width + col
}

func (a2d *Array2D) GetValueByIndex(index int) int {
	if index > a2d.length {
		return 0
	}

	return a2d.data[index]
}

func (a2d *Array2D) SetValue(col int, row int, value int) {
	var index = a2d.GetIndex(col, row)
	if index > a2d.length {
		return
	}

	a2d.data[index] = value
}

func (a2d *Array2D) GetValue(col int, row int) int {
	var index = a2d.GetIndex(col, row)
	if index > a2d.length {
		return 0
	}

	return a2d.data[index]
}
