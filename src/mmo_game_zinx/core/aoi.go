package core

import "fmt"

/*
	AOI
*/

type AOI struct {
	//区域的左边界坐标
	MinX int
	//区域的右边界坐标
	MaxX int
	//区域的上边界坐标
	MinY int
	//区域的下边界坐标
	MaxY int
	//x方向格子的数量
	NumsX int
	//y方向格子的数量
	NumsY int
	//当前区域有哪些格子 map[格子ID]格子对象
	grids map[int]*Grid
}

func NewAOI(minX, maxX, minY, maxY, numsX, numsY int) *AOI {
	aoi := &AOI{
		MinX:  minX,
		MaxX:  maxX,
		MinY:  minY,
		MaxY:  maxY,
		NumsX: numsX,
		NumsY: numsY,
		grids: make(map[int]*Grid),
	}
	//给AOI区域的所有格子  编号并初始化
	for y := 0; y < numsY; y++ {
		for x := 0; x < numsX; x++ {
			gid := y*numsY + x
			aoi.grids[gid] = NewGrid(gid,
				aoi.MinX+x*aoi.gridWidth(),
				aoi.MinX+(x+1)*aoi.gridWidth(),
				aoi.MinY+y*aoi.gridHeight(),
				aoi.MinY+(y+1)*aoi.gridHeight(),
			)
		}
	}
	return aoi
}

func (a *AOI) gridWidth() int {
	return (a.MaxX - a.MinX) / a.NumsX
}

func (a *AOI) gridHeight() int {
	return (a.MaxY - a.MinY) / a.NumsY
}

func (a AOI) String() string {
	s := fmt.Sprintf(`= = = = = = = = =  [%d X %d] AOI = = = = = = = = =
%d					%d

		

%d					%d
`+"\n该AOI区域下格子编号： ", a.NumsX, a.NumsY, a.MinX, a.MaxX, a.MinX, a.MinY)
	for key, _ := range a.grids {
		s += fmt.Sprint(key, " ")
	}
	return s
}
