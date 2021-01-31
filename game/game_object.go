package game

import "time"

type GameObject interface {
	Update(delta time.Time)
}
