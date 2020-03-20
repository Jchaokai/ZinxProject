package core

import (
	"fmt"
	"sync"
)

/*
	AOI地图的格子类型
*/

type Grid struct {
	//格子id
	GID int
	//格子的左边边界坐标
	MinX int
	//格子的右边边界坐标
	MaxX int
	//格子的上边边界坐标
	MinY int
	//格子的下边边界坐标
	MaxY int
	//当前格子内的玩家/物体的ID集合
	PlayersID map[int]bool
	//保护格子内的玩家/物体的ID集合的锁
	pIDLock sync.RWMutex
}

//初始化当前格子init
func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{GID: gid,
		MinX:      minX,
		MaxX:      maxX,
		MaxY:      maxY,
		MinY:      minY,
		PlayersID: make(map[int]bool),
	}
}

//给当前格子添加一个玩家addUser
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.PlayersID[playerID] = true
}

//给当前格子删除一个玩家delUser
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.PlayersID, playerID)
}

//获取当前格子中所有玩家getUsers
func (g *Grid) GetPlayersIDs() (players []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for key, _ := range g.PlayersID {
		players = append(players, key)
	}
	return
}

//调试使用-打印格子的信息
func (g Grid) String() string {
	return fmt.Sprintf("<< GRID ID : %d ,minX : %d ,maxX : %d ,minY : %d maxY : %d ,playersID :%v  >>\n",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayersID)
}
