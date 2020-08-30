package main

import (
	"fmt"
	"strconv"
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

func (gc *GameCommands) buySlot(arg string) {
	if slot, err := strconv.Atoi(arg); err == nil {
		if slot >= 1 && slot <= 5 {
			gc.sendCommand(fmt.Sprint("dac_buy_slot", slot))
		}
	}
}

func (gc *GameCommands) levelup(arg string) {
	gc.sendCommand("dac_shop_levelup")
}

func (gc *GameCommands) lock(arg string) {
	gc.sendCommand("dac_shop_lock")
}

func (gc *GameCommands) cameraDown(arg string) {
	gc.sendCommand("dac_scroll_camera_down")
}

func (gc *GameCommands) cameraUp(arg string) {
	gc.sendCommand("dac_scroll_camera_up")
}

func (gc *GameCommands) enemiesDown(arg string) {
	gc.sendCommand("dac_scroll_enemies_down")
}

func (gc *GameCommands) enemiesUp(arg string) {
	gc.sendCommand("dac_scroll_enemies_up")
}

func (gc *GameCommands) away(arg string) {
	gc.sendCommand("dac_view_away")
}

func (gc *GameCommands) home(arg string) {
	gc.sendCommand("dac_view_home")
}

func (gc *GameCommands) opponent(arg string) {
	gc.sendCommand("dac_view_opponent")
}

func (gc *GameCommands) dps(arg string) {
	gc.sendCommand("dac_tab_dps")
}

func (gc *GameCommands) reroll(arg string) {
	gc.sendCommand("dac_shop_reroll")
}

func (gc *GameCommands) toggle(arg string) {
	gc.sendCommand("dac_shop_toggle")
}

func (gc *GameCommands) disconnect(arg string) {
	gc.sendCommand("disconnect")
}

func (gc *GameCommands) quit(arg string) {
	gc.sendCommand("quit")
}

func (gc *GameCommands) sellUnit(arg string) {
	gc.sendCommand("dac_sell_unit")
}

func (gc *GameCommands) boardSpray(arg string) {
	gc.sendCommand("dac_board_spray_at_cursor")
}

func (gc *GameCommands) benchUnit(arg string) {
	gc.sendCommand("dac_send_unit_to_bench")
}

func (gc *GameCommands) fakeGCDown(arg string) {
	gc.sendCommand("dev_simulate_gcdown 1")

	go func() {
		time.Sleep(10 * time.Second)
		gc.sendCommand("dev_simulate_gcdown 0")
	}()
}

func (gc *GameCommands) sharecode(arg string) {
	gc.sendCommand("dac_generate_sharecode")
}
