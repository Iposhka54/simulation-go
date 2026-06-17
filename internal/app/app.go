package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simulation/internal/game/simulation"
)

type App struct {
	simulation *simulation.Simulation
	cancel     context.CancelFunc
}

func New(cancel context.CancelFunc, sim *simulation.Simulation) *App {
	return &App{
		simulation: sim,
		cancel:     cancel,
	}
}

func (a *App) Run(ctx context.Context) {
	go a.handleSignals(ctx)
	go a.handleConsoleInput(ctx)
	go a.simulation.Start(ctx)

	<-ctx.Done()

	println("Симуляция завершена!")
}

func (a *App) handleSignals(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Завершаем выполнение программы!")
	}
}

func (a *App) handleConsoleInput(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	a.printStartMenu()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if scanner.Scan() {
				switch scanner.Text() {
				case "p":
					a.simulation.Pause()
					fmt.Println("⏸️ Пауза")
				case "r":
					a.simulation.Resume()
					fmt.Println("▶️ Продолжение")
				case "q":
					fmt.Println("Выход...")
					a.cancel()
					return
				default:
					fmt.Println("Неизвестная команда. Доступно: p(pause), r(resume), q(quit)")
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
