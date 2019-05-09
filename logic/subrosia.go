package logic

// subrosia has several large areas which are guaranteed to be traverseable as
// long as you can get there in the first place:
//
// 1. "temple": rosa portal, dance hall, temple, smithy
// 2. "beach": swamp portal, market, beach
// 3. "hide and seek": H&S, mountain portal, spring tower
// 4. "pirate house": village portal, pirates
// 5. "bridge": bridge area (large but not visited in any%)
//
// "furnace" used to be on this list, but you can get there using an animal
// companion to jump over the holes at eyeglass lake, so you won't necessarily
// have feather.

var subrosiaNodes = map[string]*Node{
	"temple": Or("exit volcanoes east portal",
		And("beach", "ribbon"),
		And("beach", "jump 2"),
		And("hide and seek", "bomb jump 4"),
		And("bridge", "jump 2")),

	"beach": Or("exit subrosia market portal",
		And("hide and seek", "jump 2", "bracelet",
			Or("bomb jump 2", "magnet gloves")),
		And("furnace", "bracelet", "jump 2"),
		And("furnace", Or("jump 4", And("hard", "bomb jump 3"))),
		And("furnace", "jump 2", "magnet gloves"),
		And("temple", "jump 2")),

	"hide and seek": And("feather", Or(
		"exit strange brothers portal", "pirate house",
		And("bomb jump 4", Or("temple", "bridge")))),

	"pirate house": Or("exit house of pirates portal",
		And("hide and seek", "jump 2")),

	"furnace": Or("exit great furnace portal",
		And("beach", Or("jump 4", And("hard", "bomb jump 3"))),
		And("beach", "magnet gloves", "jump 2")),

	"bridge": Or(
		And("temple", "jump 2"),
		And("exit volcanoes west portal", "bracelet", "bomb jump 3"),
		And("hide and seek", "bomb jump 4")),

	"subrosian dance hall": AndSlot("temple"),
	"temple of seasons":    AndSlot("temple"),
	"subrosia seaside":     AndSlot("beach", "shovel"),
	"pirate's bell":        And("temple", "rusty bell"),
	"tower of winter":      AndSlot("temple", Or("hit far switch", "jump 2")),
	"tower of summer":      AndSlot("beach", "ribbon", "bracelet"),
	"tower of spring":      AndSlot("hide and seek", "jump 2"),
	"tower of autumn":      AndSlot("temple", "jump 2", "bomb flower"),
	"subrosian wilds chest": AndSlot("hide and seek",
		Or("jump 4", "magnet gloves")),
	"subrosia village chest": OrSlot(
		And("beach", "magnet gloves"),
		And("furnace", "jump 2", Or("jump 4", "magnet gloves"))),
	"subrosia, open cave":       AndSlot("bridge"),
	"subrosia, locked cave":     AndSlot("beach", "ribbon", "jump 2"),
	"subrosia market, 1st item": AndSlot("beach", "star ore"),
	"subrosia market, 2nd item": AndSlot("beach", "ore chunks", "ember seeds"),
	"subrosia market, 5th item": AndSlot("beach", "ore chunks"),
	"great furnace": AndSlot("furnace", "red ore", "blue ore",
		"temple", "bomb flower"),
	"subrosian smithy": AndSlot("temple", "hard ore"),

	"d8 entrance": And("exit D8 entrance portal"),
}
