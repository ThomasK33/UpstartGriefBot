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

// SendCommand - Function writing commands to the game console
func (gc *GameCommands) SendCommand(command string) {
	if gc.Caller != nil && gc.Caller.Writer != nil {
		oi.LongWrite(*gc.Caller.Writer, []byte("dac_test_player_chat_global "+command))
		oi.LongWrite(*gc.Caller.Writer, []byte("\n"))
		oi.LongWrite(*gc.Caller.Writer, []byte(command))
		oi.LongWrite(*gc.Caller.Writer, []byte("\n"))
	}
}

// BuySlot - Buy a unit from the shop from slot number 1-5
func (gc *GameCommands) BuySlot(arg string) {
	if slot, err := strconv.Atoi(arg); err == nil {
		if slot >= 1 && slot <= 5 {
			gc.SendCommand(fmt.Sprint("dac_buy_slot", slot))
		}
	}
}

// Levelup - Send a levelup command
func (gc *GameCommands) Levelup(arg string) {
	gc.SendCommand("dac_shop_levelup")
}

// Lock - Lock the shop
func (gc *GameCommands) Lock(arg string) {
	gc.SendCommand("dac_shop_lock")
}

// CameraDown - Move the camera down
func (gc *GameCommands) CameraDown(arg string) {
	gc.SendCommand("dac_scroll_camera_down")
}

// CameraUp - Move the camera up
func (gc *GameCommands) CameraUp(arg string) {
	gc.SendCommand("dac_scroll_camera_up")
}

// EnemiesDown - Move on enemy board down
func (gc *GameCommands) EnemiesDown(arg string) {
	gc.SendCommand("dac_scroll_enemies_down")
}

// EnemiesUp - Move one enemy board up
func (gc *GameCommands) EnemiesUp(arg string) {
	gc.SendCommand("dac_scroll_enemies_up")
}

// Away - Spectate the away board
func (gc *GameCommands) Away(arg string) {
	gc.SendCommand("dac_view_away")
}

// Home - Spectate own board
func (gc *GameCommands) Home(arg string) {
	gc.SendCommand("dac_view_home")
}

// Opponent - Spectate current's opponents board
func (gc *GameCommands) Opponent(arg string) {
	gc.SendCommand("dac_view_opponent")
}

// Dps - Switch to the dps tab
func (gc *GameCommands) Dps(arg string) {
	gc.SendCommand("dac_tab_dps")
}

// Reroll - Reroll the shop
func (gc *GameCommands) Reroll(arg string) {
	gc.SendCommand("dac_shop_reroll")
}

// Toggle - Toggle the shop
func (gc *GameCommands) Toggle(arg string) {
	gc.SendCommand("dac_shop_toggle")
}

// Disconnect - Disconnect from the current server
func (gc *GameCommands) Disconnect(arg string) {
	gc.SendCommand("disconnect")
}

// Quit - Quit the game client
func (gc *GameCommands) Quit(arg string) {
	gc.SendCommand("quit")
}

// SellUnit - Sell the currently selected unit
func (gc *GameCommands) SellUnit(arg string) {
	gc.SendCommand("dac_sell_unit")
}

// BoardSpray - Spray a board spray at the current cursor location
func (gc *GameCommands) BoardSpray(arg string) {
	gc.SendCommand("dac_board_spray_at_cursor")
}

// BenchUnit - Bench the currently selected unit
func (gc *GameCommands) BenchUnit(arg string) {
	gc.SendCommand("dac_send_unit_to_bench")
}

// FakeGCDown - Fake a game coordinator downtime
func (gc *GameCommands) FakeGCDown(arg string) {
	gc.SendCommand("dev_simulate_gcdown 1")

	go func() {
		time.Sleep(10 * time.Second)
		gc.SendCommand("dev_simulate_gcdown 0")
	}()
}

// Sharecode - Generate a board sharecode and print it in the game's console
func (gc *GameCommands) Sharecode(arg string) {
	gc.SendCommand("dac_generate_sharecode")
}
