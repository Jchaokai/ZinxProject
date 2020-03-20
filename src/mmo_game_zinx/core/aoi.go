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
			gid := y*numsX + x
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

func (a *AOI) GetSurroundingByGid(centerID int) (grids []*Grid) {
	//判断gid是否在AOI中
	if _, ok := a.grids[centerID]; !ok {
		return
	}
	grids = append(grids, a.grids[centerID])
	//通过格子编号获取 x轴编号 idx
	idx := centerID % a.NumsX
	//判断idx 左右是否有格子，如果有就放入结果
	if idx > 0 {
		grids = append(grids, a.grids[centerID-1])
	}
	if idx < a.NumsX-1 {
		grids = append(grids, a.grids[centerID+1])
	}
	//横向格子的 gid slice
	gid := make([]int, 0)
	for _, grid := range grids {
		gid = append(gid, grid.GID)
	}
	//遍历gid slice，判断这些格子的上下是否有格子
	for _, v := range gid {
		//获取x周格子对应的y轴 编号 idy
		idy := v / a.NumsX
		//判断上面有没有
		if idy > 0 {
			grids = append(grids, a.grids[v-a.NumsX])
		}
		//判断下面有没有
		if idy < a.NumsY-1 {
			grids = append(grids, a.grids[v+a.NumsX])
		}
	}
	return
}

func (a AOI) String() string {
	s := fmt.Sprintf(`= = = = =   [%d X %d] AOI  每个块宽:%d,高:%d = = = = = 
     | %d				 %d   |
_____|_________________________|__			
%d   |    					   |
     |						   |
     |						   |
     |						   |
%d  |						   | 
_____|_________________________|__
	 |						   |
该AOI区域下每个格子信息如下： `+"\n", a.NumsX, a.NumsY, a.gridWidth(), a.gridHeight(), a.MinX, a.MaxX, a.MinY, a.MaxY)
	for _, grid := range a.grids {
		s += fmt.Sprint(grid)
	}
	return s
}
