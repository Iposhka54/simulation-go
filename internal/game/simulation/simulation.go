package simulation

import (
	"simulation/internal/game/action"
	_map "simulation/internal/game/map"
	"simulation/internal/game/renderer"
	"time"
)

type Simulation struct {
	turn        uint64
	worldMap    *_map.Map
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	interval    time.Duration
	pauseChan   chan struct{}
	stopChan    chan struct{}
	running     bool
	paused      bool
}

func New(worldMap *_map.Map, delayMs int, render renderer.Renderer, initActions, turnActions []action.Action) *Simulation {
	simulation := &Simulation{
		turn:        0,
		worldMap:    worldMap,
		renderer:    render,
		initActions: initActions,
		turnActions: turnActions,
		interval:    time.Duration(delayMs) * time.Millisecond,
		pauseChan:   make(chan struct{}),
		stopChan:    make(chan struct{}),
		running:     false,
		paused:      false,
	}
	simulation.init()
	return simulation
}

// StartSimulation blocking operation, need invoke in a separate goroutine
func (s *Simulation) StartSimulation() {
	if s.running {
		return
	}

	s.running = true
	s.paused = false

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()
	for {
		if s.paused {
			select {
			case <-s.pauseChan:
				s.paused = false
			case <-s.stopChan:
				s.paused = false
				s.running = false
				return
			}
			continue
		}

		select {
		case <-s.stopChan:
			s.running = false
			return
		case <-s.pauseChan:
			s.paused = true
		case <-ticker.C:
			s.NextTurn()
		}
	}
}

func (s *Simulation) PauseSimulation() {
	if s.running && !s.paused {
		s.pauseChan <- struct{}{}
	}
}

func (s *Simulation) ResumeSimulation() {
	if s.running && s.paused {
		s.pauseChan <- struct{}{}
	}
}

func (s *Simulation) StopSimulation() {
	if !s.running {
		return
	}
	close(s.stopChan)
}

func (s *Simulation) NextTurn() {
	s.turn++
	for _, turnAction := range s.turnActions {
		turnAction.Execute(s.worldMap)
	}

	s.printTurnHeader()
	s.renderer.Render(s.worldMap)
}

func (s *Simulation) init() {
	for _, initAction := range s.initActions {
		initAction.Execute(s.worldMap)
	}
}

func (s *Simulation) printTurnHeader() {
	println("Turn: ", s.turn)
}
