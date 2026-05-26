package simulation

import (
	"simulation/internal/game/action"
	"simulation/internal/game/renderer"
	"simulation/internal/game/world"
	"time"
)

type Simulation struct {
	turn        uint64
	world       *world.World
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	interval    time.Duration
	pauseChan   chan struct{}
	stopChan    chan struct{}
	running     bool
	paused      bool
}

func New(world *world.World, delayMs int, render renderer.Renderer, initActions, turnActions []action.Action) *Simulation {
	simulation := &Simulation{
		turn:        0,
		world:       world,
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

// Start blocking operation, need invoke in a separate goroutine
func (s *Simulation) Start() {
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
			s.nextTurn()
		}
	}
}

func (s *Simulation) Pause() {
	if s.running && !s.paused {
		s.pauseChan <- struct{}{}
	}
}

func (s *Simulation) Resume() {
	if s.running && s.paused {
		s.pauseChan <- struct{}{}
	}
}

func (s *Simulation) Stop() {
	if !s.running {
		return
	}
	close(s.stopChan)
}

func (s *Simulation) nextTurn() {
	s.turn++
	for _, turnAction := range s.turnActions {
		turnAction.Execute(s.world)
	}

	s.printTurnHeader()
	s.renderer.Render(s.world)
}

func (s *Simulation) init() {
	for _, initAction := range s.initActions {
		initAction.Execute(s.world)
	}
}

func (s *Simulation) printTurnHeader() {
	println("Turn: ", s.turn)
}
