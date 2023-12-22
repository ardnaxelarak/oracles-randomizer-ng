package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Reads a ".sym" file, returns a tuple with 2 elements:
// - Map of addresses ([labels] section)
// - Map of definitions ([definitions] section)
func readSymbolFile(filename string) (map[string]*address, map[string]uint32) {
	symbolPanic := func(line string) {
		panic(fmt.Sprintf("Error parsing symbol file \"%s\": \"%s\"", filename, line))
	}

	labels := make(map[string]*address)
	defs := make(map[string]uint32)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't open symbol file \"%s\".", filename))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \t")
		if strings.HasPrefix(s, ";") || strings.HasPrefix(s, "[") || len(s) == 0 {
			continue
		}
		list := strings.Split(s, " ")
		if len(list) != 2 {
			symbolPanic(s)
		}
		addr := strings.Split(list[0], ":")
		if len(addr) == 2 {
			bank, _ := strconv.ParseUint(addr[0], 16, 8)
			offset, _ := strconv.ParseUint(addr[1], 16, 16)
			labels[list[1]] = &address{uint8(bank), uint16(offset)}
		} else {
			// Not an address, then it must be a definitions
			value, _ := strconv.ParseUint(list[0], 16, 32)
			defs[list[1]] = uint32(value)
		}
	}

	return labels, defs
}
