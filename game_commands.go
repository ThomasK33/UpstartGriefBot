package main

import (
	"fmt"
	"time"

	"github.com/reiver/go-oi"
)

// GameCommands Structure holding all game related commands
type GameCommands struct {
	tc *telnetCaller
}

func (gc *GameCommands) sendCommand(command string) {
	oi.LongWrite(*gc.tc.w, []byte(command))
	oi.LongWrite(*gc.tc.w, []byte("\n"))
}

func (gc *GameCommands) buySlot(slot int) {
	gc.sendCommand(fmt.Sprint("dac_buy_slot", slot))
}

func (gc *GameCommands) levelup() {
	gc.sendCommand("dac_shop_levelup")
}

func (gc *GameCommands) lock() {
	gc.sendCommand("dac_shop_lock")
}

func (gc *GameCommands) cameraDown() {
	gc.sendCommand("dac_scroll_camera_down")
}

func (gc *GameCommands) cameraUp() {
	gc.sendCommand("dac_scroll_camera_up")
}

func (gc *GameCommands) enemiesDown() {
	gc.sendCommand("dac_scroll_enemies_down")
}

func (gc *GameCommands) enemiesUp() {
	gc.sendCommand("dac_scroll_enemies_up")
}

func (gc *GameCommands) away() {
	gc.sendCommand("dac_view_away")
}

func (gc *GameCommands) home() {
	gc.sendCommand("dac_view_home")
}

func (gc *GameCommands) opponent() {
	gc.sendCommand("dac_view_opponent")
}

func (gc *GameCommands) dps() {
	gc.sendCommand("dac_tab_dps")
}

func (gc *GameCommands) reroll() {
	gc.sendCommand("dac_shop_reroll")
}

func (gc *GameCommands) toggle() {
	gc.sendCommand("dac_shop_toggle")
}

func (gc *GameCommands) disconnect() {
	gc.sendCommand("disconnect")
}

func (gc *GameCommands) quit() {
	gc.sendCommand("quit")
}

func (gc *GameCommands) sellUnit() {
	gc.sendCommand("dac_sell_unit")
}

func (gc *GameCommands) boardSpray() {
	gc.sendCommand("dac_board_spray_at_cursor")
}

func (gc *GameCommands) benchUnit() {
	gc.sendCommand("dac_send_unit_to_bench")
}

func (gc *GameCommands) fakeGCDown() {
	gc.sendCommand("dev_simulate_gcdown 1")

	go func() {
		time.Sleep(10 * time.Second)
		gc.sendCommand("dev_simulate_gcdown 0")
	}()
}

func (gc *GameCommands) sharecode() {
	gc.sendCommand("dac_generate_sharecode")
}
