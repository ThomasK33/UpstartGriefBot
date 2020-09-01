package src

import (
	"fmt"
	"strconv"
	"time"

	"github.com/reiver/go-oi"
)

// GameCommands - Structure holding all game related commands
type GameCommands struct {
	Caller *TelnetCaller
}

// SendCommand -
func (gc *GameCommands) SendCommand(command string) {
	if gc.Caller != nil && gc.Caller.Writer != nil {
		oi.LongWrite(*gc.Caller.Writer, []byte(command))
		oi.LongWrite(*gc.Caller.Writer, []byte("\n"))
		oi.LongWrite(*gc.Caller.Writer, []byte("dac_test_player_chat_global "+command))
		oi.LongWrite(*gc.Caller.Writer, []byte("\n"))
	}
}

// BuySlot -
func (gc *GameCommands) BuySlot(arg string) {
	if slot, err := strconv.Atoi(arg); err == nil {
		if slot >= 1 && slot <= 5 {
			gc.SendCommand(fmt.Sprint("dac_buy_slot", slot))
		}
	}
}

// Levelup -
func (gc *GameCommands) Levelup(arg string) {
	gc.SendCommand("dac_shop_levelup")
}

// Lock -
func (gc *GameCommands) Lock(arg string) {
	gc.SendCommand("dac_shop_lock")
}

// CameraDown -
func (gc *GameCommands) CameraDown(arg string) {
	gc.SendCommand("dac_scroll_camera_down")
}

// CameraUp -
func (gc *GameCommands) CameraUp(arg string) {
	gc.SendCommand("dac_scroll_camera_up")
}

// EnemiesDown -
func (gc *GameCommands) EnemiesDown(arg string) {
	gc.SendCommand("dac_scroll_enemies_down")
}

// EnemiesUp -
func (gc *GameCommands) EnemiesUp(arg string) {
	gc.SendCommand("dac_scroll_enemies_up")
}

// Away -
func (gc *GameCommands) Away(arg string) {
	gc.SendCommand("dac_view_away")
}

// Home -
func (gc *GameCommands) Home(arg string) {
	gc.SendCommand("dac_view_home")
}

// Opponent -
func (gc *GameCommands) Opponent(arg string) {
	gc.SendCommand("dac_view_opponent")
}

// Dps -
func (gc *GameCommands) Dps(arg string) {
	gc.SendCommand("dac_tab_dps")
}

// Reroll -
func (gc *GameCommands) Reroll(arg string) {
	gc.SendCommand("dac_shop_reroll")
}

// Toggle -
func (gc *GameCommands) Toggle(arg string) {
	gc.SendCommand("dac_shop_toggle")
}

// Disconnect -
func (gc *GameCommands) Disconnect(arg string) {
	gc.SendCommand("disconnect")
}

// Quit -
func (gc *GameCommands) Quit(arg string) {
	gc.SendCommand("quit")
}

// SellUnit -
func (gc *GameCommands) SellUnit(arg string) {
	gc.SendCommand("dac_sell_unit")
}

// BoardSpray -
func (gc *GameCommands) BoardSpray(arg string) {
	gc.SendCommand("dac_board_spray_at_cursor")
}

// BenchUnit -
func (gc *GameCommands) BenchUnit(arg string) {
	gc.SendCommand("dac_send_unit_to_bench")
}

// FakeGCDown -
func (gc *GameCommands) FakeGCDown(arg string) {
	gc.SendCommand("dev_simulate_gcdown 1")

	go func() {
		time.Sleep(10 * time.Second)
		gc.SendCommand("dev_simulate_gcdown 0")
	}()
}

// Sharecode -
func (gc *GameCommands) Sharecode(arg string) {
	gc.SendCommand("dac_generate_sharecode")
}
