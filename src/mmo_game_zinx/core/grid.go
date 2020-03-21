package core

import (
	"fmt"
	"sync"
)

/*
	AOI地图的格子类型
*/

type Grid struct {
	GID  int
	MinX int
	MaxX int
	MinY int
	MaxY int
	//当前格子内的玩家/物体的ID集合
	PlayersID map[int]bool
	pIDLock   sync.RWMutex
}

//初始化当前格子
func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{GID: gid,
		MinX:      minX,
		MaxX:      maxX,
		MaxY:      maxY,
		MinY:      minY,
		PlayersID: make(map[int]bool),
	}
}

//给当前格子添加一个玩家
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.PlayersID[playerID] = true
}

//给当前格子删除一个玩家
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.PlayersID, playerID)
}

//获取当前格子中所有玩家
func (g *Grid) GetPlayersIDs() (players []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for key := range g.PlayersID {
		players = append(players, key)
	}
	return
}

func (g Grid) String() string {
	return fmt.Sprintf("<< GRID ID : %d ,minX : %d ,maxX : %d ,minY : %d maxY : %d ,playersID :%v  >>\n",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayersID)
}
