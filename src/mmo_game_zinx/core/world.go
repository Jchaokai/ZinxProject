package core

import "sync"

var (
	WorldObj *World
)

/*
	当前游戏的世界 总管理没模块
*/

type World struct {
	Aoi     *AOI
	Players map[int32]*Player
	plock   sync.RWMutex
}

//初始化
func init() {
	//游戏的全部地图
	WorldObj = &World{
		Aoi:     NewAOI(0, 10000, 0, 10000, 200, 200),
		Players: make(map[int32]*Player),
	}
}

//添加一个玩家
func (w *World) AddPlayer(player *Player) {
	w.plock.Lock()
	defer w.plock.Unlock()
	w.Players[player.Pid] = player
	//将player添加到格子中
	w.Aoi.AddPlayerToGridByPos(int(player.Pid), player.x, player.z)
}

//删除一个玩家
func (w *World) RemovePlayerByPid(pid int32) {
	player := w.Players[pid]
	w.Aoi.RemovePlayFromGridByPos(int(pid), player.x, player.z)
	w.plock.Lock()
	defer w.plock.Unlock()
	delete(w.Players, pid)
}

//通过玩家ID查询玩家对象
func (w *World) GetPlayerByID(pid int32) *Player {
	w.plock.RLock()
	defer w.plock.RUnlock()
	return w.Players[pid]
}

//获取全部的在线玩家
func (w *World) GetAllPlayers() (players []*Player) {
	w.plock.RLock()
	defer w.plock.RUnlock()
	for _, p := range w.Players {
		players = append(players, p)
	}
	return
}
