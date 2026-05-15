package action

import _map "simulation/internal/game/map"

type Action interface {
	Execute(worldMap *_map.Map)
}
