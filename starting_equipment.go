package main

import (
	"regexp"
	"strconv"
)

var ringRegex = regexp.MustCompile(` ring(?: L-[1-3])?$`)
var equipRegex = regexp.MustCompile(`^([^,]+)(?:, *(\d+))?$`)
var labelRegex = regexp.MustCompile(`^([^+]+)(?:\+(0x[[:xdigit:]]+|\d+))?$`)
var keyRegex = regexp.MustCompile(`^([^_]+)_(0x[[:xdigit:]]+|\d+)?$`)

type startingEquipment struct {
	data      map[int]byte
	items     []string
	inventory []byte
	seedItem  bool
	seedAddrs []int
	satchel   int
}

func parseStartingEquipment(rom *romState, ropts *randomizerOptions) *startingEquipment {
	eq := &startingEquipment{seedItem: false, satchel: 0}

	eq.data = make(map[int]byte)
	eq.items = make([]string, 0)
	eq.inventory = make([]byte, 0, 18)
	eq.seedAddrs = make([]int, 0)

	treasureAddr := rom.lookupLabel("rando_startingObtainedTreasureFlags").fullOffset()
	eq.setBit(treasureAddr, 2) // TREASURE_PUNCH

	ringAddr := rom.lookupLabel("rando_startingRings").fullOffset()
	equipmentData := loadStartingEquipmentData(rom.game)
	equipmentPacks := loadStartingEquipmentPacks(rom.game)

	starting := make(map[string]int)

	for _, s := range ropts.starting {
		if ringRegex.MatchString(s) {
			if id := getStringIndex(rings, s); id != -1 {
				eq.setBit(ringAddr, id)
			} else {
				panic("no such ring: " + s)
			}
		} else if pack, found := equipmentPacks[s]; found {
			for _, v := range pack {
				equipMatch := equipRegex.FindStringSubmatch(v)
				if equipMatch == nil {
					panic("no such item: " + v)
				}
				itemName := equipMatch[1]
				count := 1
				if equipMatch[2] != "" {
					count, _ = strconv.Atoi(equipMatch[2])
				}
				starting[itemName] += count
			}
		} else {
			equipMatch := equipRegex.FindStringSubmatch(s)
			if equipMatch == nil {
				panic("no such item: " + s)
			}
			itemName := equipMatch[1]
			count := 1
			if equipMatch[2] != "" {
				count, _ = strconv.Atoi(equipMatch[2])
			}
			starting[itemName] += count
			item, ok := equipmentData[itemName]
			if !ok {
				panic("no such item: " + itemName)
			}
			if item["include"] != "" {
				starting[item["include"]] += 1
			}
		}
	}

	for itemName, count := range starting {
		for i := 0; i < count; i++ {
			eq.items = append(eq.items, itemName)
		}
		if itemName == "satchel" {
			eq.satchel = count
		}
		item, _ := equipmentData[itemName]
		eq.addStartingItem(rom, item, count)
	}

	bigSword := false
	baseAddr := rom.lookupLabel("rando_startingInventory").fullOffset()
	invAddr := baseAddr
	for _, v := range eq.inventory {
		if v == 0x0c {
			bigSword = true
		} else {
			eq.data[invAddr] = v
			invAddr += 1
		}
	}
	if bigSword {
		if invAddr == baseAddr + 1 {
			invAddr += 1
		}
		eq.data[invAddr] = 0x0c
	}

	return eq
}

func (eq *startingEquipment) addStartingItem(rom *romState, item map[string]string, count int) {
	for k, v := range item {
		if k == "itemId" {
			value, _ := strconv.ParseInt(v, 0, 0)
			eq.inventory = append(eq.inventory, byte(value))
		} else if k == "treasure" {
			addr := rom.lookupLabel("rando_startingObtainedTreasureFlags").fullOffset()
			value, _ := strconv.ParseInt(v, 0, 0)
			eq.setBit(addr, int(value))
		} else if k == "setCount" {
			labelMatch := labelRegex.FindStringSubmatch(v)
			if labelMatch == nil {
				panic("invalid label expression: " + v)
			}
			addr := rom.lookupLabel(labelMatch[1]).fullOffset()
			if labelMatch[2] != "" {
				offset, _ := strconv.ParseInt(labelMatch[2], 0, 0)
				addr += int(offset)
			}
			eq.data[addr] = byte(count)
		} else if k == "remove" {
			eq.items = append(eq.items, v)
		} else if k == "startingSeeds" {
			eq.seedItem = true
		} else if k == "fillSatchel" {
			addr := rom.lookupLabel(v).fullOffset()
			eq.seedAddrs = append(eq.seedAddrs, addr)
		} else if keyMatch := keyRegex.FindStringSubmatch(k); keyMatch != nil {
			firstBit := keyMatch[1]
			secondBit, _ := strconv.ParseInt(keyMatch[2], 0, 0)
			if firstBit == "countAs" && count >= int(secondBit) {
				eq.items = append(eq.items, v)
			} else if firstBit == "countTreasure" && count >= int(secondBit) {
				addr := rom.lookupLabel("rando_startingObtainedTreasureFlags").fullOffset()
				value, _ := strconv.ParseInt(v, 0, 0)
				eq.setBit(addr, int(value))
			} else if firstBit == "set" {
				addr := rom.lookupLabel(v).fullOffset()
				eq.data[addr] = byte(secondBit)
			} else if firstBit == "setBit" {
				addr := rom.lookupLabel(v).fullOffset()
				eq.setBit(addr, int(secondBit))
			}
		}
	}
}

func (eq *startingEquipment) setStartingSeeds(rom *romState, seedId byte) {
	if eq.seedItem {
		seedTreasure := seedId + 0x20
		addr := rom.lookupLabel("rando_startingObtainedTreasureFlags").fullOffset()
		eq.setBit(addr, int(seedTreasure))

		seedAddr := rom.lookupLabel("rando_startingNumEmberSeeds").fullOffset() + int(seedId)
		eq.seedAddrs = append(eq.seedAddrs, seedAddr)
	}

	capacity := 0x20

	if eq.satchel == 2 {
		capacity = 0x50
	} else if eq.satchel > 2 {
		capacity = 0x99
	}

	for _, addr := range eq.seedAddrs {
		eq.data[addr] = byte(capacity)
	}
}

func (eq *startingEquipment) setBit(baseAddr int, bitNum int) {
	addr := baseAddr + bitNum/8
	oldValue, _ := eq.data[addr]
	eq.data[addr] = oldValue | (1 << (bitNum % 8))
}

func loadStartingEquipmentData(game int) map[string]map[string]string {
	rawData := make(map[string]map[string]map[string]string)
	if err := readYaml("romdata/startingData.yaml", rawData); err != nil {
		panic(err)
	}

	gameData := make(map[string]map[string]string)
	for k, v := range rawData["common"] {
		gameData[k] = v
	}
	for k, v := range rawData[gameNames[game]] {
		gameData[k] = v
	}

	return gameData
}

func loadStartingEquipmentPacks(game int) map[string][]string {
	rawData := make(map[string]map[string][]string)
	if err := readYaml("romdata/startingPacks.yaml", rawData); err != nil {
		panic(err)
	}

	gameData := make(map[string][]string)
	for k, v := range rawData["common"] {
		gameData[k] = v
	}
	for k, v := range rawData[gameNames[game]] {
		gameData[k] = v
	}

	return gameData
}
