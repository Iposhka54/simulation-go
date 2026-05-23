package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"simulation/internal/game/simulation"
	"syscall"
)

type App struct {
	simulation *simulation.Simulation
	ctx        context.Context
	cancel     context.CancelFunc
}

func New(sim *simulation.Simulation) *App {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &App{
		simulation: sim,
		ctx:        ctx,
		cancel:     cancelFunc,
	}
}

func (a *App) Run() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go a.handleSignals(sigChan)
	go a.handleConsoleInput()
	go a.simulation.StartSimulation()

	<-a.ctx.Done()

	a.simulation.StopSimulation()
	println("Симуляция завершена!")
}

func (a *App) handleSignals(sigChan <-chan os.Signal) {
	select {
	case <-a.ctx.Done():
		return
	case sig := <-sigChan:
		println("Завершаем выполнение программы из-за сигнала: ", sig)
		a.cancel()
	}
}

func (a *App) handleConsoleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	a.printStartMenu()
	for {
		select {
		case <-a.ctx.Done():
			return
		default:
			if scanner.Scan() {
				switch scanner.Text() {
				case "p":
					a.simulation.PauseSimulation()
					fmt.Println("▶️ Пауза")
				case "r":
					a.simulation.ResumeSimulation()
					fmt.Println("▶️ Продолжение")
				case "q":
					fmt.Println("Выход...")
					a.cancel()
					return
				default:
					fmt.Println("Неизвестная команда. Доступно: pause, resume, quit")
				}
			}
		}
	}
}

func (a *App) printStartMenu() {
	println("\n=== Управление ===")
	println("  p  - пауза")
	println("  r - продолжить")
	println("  q   - выход")
	println("==================")
}
