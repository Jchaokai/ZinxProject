package core

import (
	"fmt"
	"testing"
)

func TestAOI_NewAOI(t *testing.T) {
	//初始化一个AOI区域
	aoi := NewAOI(100, 400, 100, 300, 6, 4)
	//打印
	fmt.Println(aoi.String())
}

func TestAOI_GetSurroundingByGid(t *testing.T) {
	aoi := NewAOI(100, 400, 100, 300, 6, 4)
	for gid := range aoi.grids {
		fmt.Print("中心格子：", aoi.grids[gid])
		//得到每个格子所属的九宫格信息
		grids := aoi.GetSurroundingByGid(gid)
		for _, grid := range grids {
			fmt.Print(grid)
		}
		fmt.Println()
	}
}
