package patterns

import "fmt"

/*
	Паттерн "Состояние" позволяет объекту изменять свое поведение при изменении внутреннего состояния.
	Это помогает убрать громоздкие условные операторы и сделать код более читаемым и расширяемым.

	Плюсы:
	Уменьшение зависимости
	Читабельность
	Гибкость

	Минусы:
	Усложнение
	Избыточность
*/

// ? Интерфейс состояния

type State interface {
	play(player *MusicPlayer)
	pause(player *MusicPlayer)
	stop(player *MusicPlayer)
}

// ? Музыкальный плеер

type MusicPlayer struct {
	state State
}

func (p *MusicPlayer) setState(state State) {
	p.state = state
}

func (p *MusicPlayer) play() {
	p.state.play(p)
}

func (p *MusicPlayer) pause() {
	p.state.pause(p)
}

func (p *MusicPlayer) stop() {
	p.state.stop(p)
}

// ? Состояние воспроизведения

type PlayingState struct{}

func (s *PlayingState) play(player *MusicPlayer) {
	fmt.Println("Music is already playing")
}

func (s *PlayingState) pause(player *MusicPlayer) {
	fmt.Println("Music paused")
	player.setState(&PausedState{})
}

func (s *PlayingState) stop(player *MusicPlayer) {
	fmt.Println("Music stopped")
	player.setState(&StoppedState{})
}

// ? Состояние паузы

type PausedState struct{}

func (s *PausedState) play(player *MusicPlayer) {
	fmt.Println("Resuming music")
	player.setState(&PlayingState{})
}

func (s *PausedState) pause(player *MusicPlayer) {
	fmt.Println("Music is already paused")
}

func (s *PausedState) stop(player *MusicPlayer) {
	fmt.Println("Music stopped from pause")
	player.setState(&StoppedState{})
}

// ? Состояние остановки

type StoppedState struct{}

func (s *StoppedState) play(player *MusicPlayer) {
	fmt.Println("Playing music from stop")
	player.setState(&PlayingState{})
}

func (s *StoppedState) pause(player *MusicPlayer) {
	fmt.Println("Music is already stopped")
}

func (s *StoppedState) stop(player *MusicPlayer) {
	fmt.Println("Music is already stopped")
}

func CheckState() {
	player := &MusicPlayer{state: &StoppedState{}}

	player.play()
	player.pause()
	player.play()
	player.stop()
	player.stop()
}
