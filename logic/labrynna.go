package logic

// this all assumes that you start in the forest of time and that the time
// portals on the screens next to the maku tree are always active.

var labrynnaNodes = map[string]*Node{
	// forest of time
	"start":          And(),
	"starting chest": AndSlot("start"),
	"nayru's house":  AndSlot("start"),

	// lynna / south shore / palace
	"lynna city":         Or("break bush", "flute", "echoes"),
	"lynna village":      Or("lynna city", "echoes"),
	"black tower worker": AndSlot("lynna village"),
	"maku tree": OrSlot("rescue nayru",
		And("lynna village", "shovel", "kill normal")),
	"south lynna tree": AndSlot("lynna city", "seed item",
		Or("sword", "punch object", "dimitri's flute", Hard("break bush"))),
	"lynna city chest": OrSlot("ember seeds", "currents"),
	"shore present": Or("flute", "ricky's gloves",
		And("break bush", "feather"), And("lynna city", "bracelet"),
		And("currents", Or("feather", "flippers", "raft")),
		And("ages", "break bush"),
		And("lynna city", "mermaid suit")),
	"south shore dirt": AndSlot("shore present", Or("shovel", "flute")),
	"balloon guy": And("feather", Or("sword", "boomerang"),
		Or("currents", "ricky's gloves", "ricky's flute",
			And("shore present", Or("any seed shooter",
				HardAnd(Or("pegasus satchel", "bombs"), "boomerang"))))),
	"balloon guy's gift": AndSlot("balloon guy"),
	"balloon guy's upgrade": AndSlot("balloon guy", Or( // 3 types of seeds
		And("ember seeds", Or(
			And("scent seeds",
				Or("pegasus seeds", "gale seeds", "mystery seeds")),
			And("pegasus seeds", Or("gale seeds", "mystery seeds")),
			And("gale seeds", "mystery seeds"))),
		And("scent seeds", Or(
			And("pegasus seeds", Or("gale seeds", "mystery seeds")),
			And("gale seeds", "mystery seeds"))),
		And("pegasus seeds", "gale seeds", "mystery seeds"))),
	"raft":             And("lynna village", "cheval rope", "island chart"),
	"shop, 30 rupees":  AndSlot("lynna city"),
	"shop, 150 rupees": AndSlot("lynna city"),
	"ambi's palace tree": AndSlot("lynna village", Or("sword", "punch object"),
		"seed item"),
	"ambi's palace chest": AndSlot("lynna village", Or("ages",
		HardAnd("satchel", "scent seeds", "pegasus seeds"),
		And("break bush safe", "mermaid suit"))),
	"rescue nayru": AndSlot("ambi's palace chest", "mystery seeds",
		"switch hook", Or("sword", "punch enemy")),
	"mayor plen's house": AndSlot("long hook"),
	"maku seed": And("d1 boss", "d2 boss", "d3 boss", "d4 boss", "d5 boss",
		"d6 boss", "d7 boss", "d8 boss"),

	// yoll graveyard
	"yoll graveyard": And("ember seeds"),
	"cheval's grave": And("yoll graveyard", Or("kill ghini", "bomb jump 3")),
	"cheval's test": AndSlot("cheval's grave", "bracelet",
		Or("feather", "flippers")),
	"cheval's invention": AndSlot("cheval's grave", "flippers"),
	"grave under tree":   AndSlot("yoll graveyard"),
	"syrup": And("yoll graveyard", "graveyard key",
		Or("flippers", "bomb jump 2", "dimitri's flute", "long hook")),
	"graveyard poe": AndSlot("yoll graveyard", "graveyard key", "bracelet"),
	"enter d1":      And("yoll graveyard", "graveyard key"),

	// western woods
	// it's possible to switch hook the octorok through the boulder to enter
	// fairies' woods (easier if you have scent seeds), but it's not in logic
	// for the same reason that cucco clip isn't in seasons logic.
	"fairies' woods": Or("bracelet", "dimitri's flute", "ages",
		And("lynna city", "flippers"),
		And("currents", Or("hit lever", "ricky's flute", "moosh's flute"))),
	"fairies' woods chest": OrSlot(And("deku forest", "currents"),
		And("fairies' woods",
			Or("feather", "ricky's flute", "moosh's flute", "switch hook"))),
	"deku forest":           Or("bracelet", "ages"),
	"deku forest cave east": AndSlot("deku forest"),
	"deku forest cave west": AndSlot("deku forest", "bracelet",
		Or("feather", "switch hook", "ember seeds", "ages", "gale satchel")),
	"deku forest tree": AndSlot("deku forest", Or("sword", "punch object"),
		"seed item", Or("ember seeds", "ages", "switch hook", "gale satchel",
			HardAnd("feather", Or("sword", "bombs")))),
	"deku forest soldier": AndSlot("deku forest", "mystery seeds"),
	"enter d2":            And("deku forest", Or("bombs", "currents")),

	// crescent island
	// the western present portal responds to currents only, in order to
	// prevent softlocks.
	"crescent past": Or("raft", And("lynna city", "mermaid suit"),
		And("crescent present west", "currents")),
	"tokay crystal cave": AndSlot("crescent past",
		Or("shovel", "break crystal"), "feather"),
	"tokay bomb cave": AndSlot("crescent past", "bracelet", "bombs"),
	"wild tokay game": AndSlot("crescent past", "bombs", "bracelet"),
	// can get the warp point by swimming under crescent island, but that's
	// pretty unintuitive.
	"crescent island tree": AndSlot("crescent past", "scent seedling",
		Or("sword", "punch object"), "seed item", Or("ages", And("bracelet",
			Or("echoes", HardAnd("gale satchel", "mermaid suit"))))),
	"crescent present west": Or("dimitri's flute",
		And("lynna city", "mermaid suit"),
		And("crescent past", Or("currents", And("shovel", "echoes")))),
	"enter d3":              And("crescent present west"),
	"hidden tokay cave":     AndSlot("lynna city", "mermaid suit"),
	"under crescent island": AndSlot("lynna city", "mermaid suit"),
	"tokay pot cave":        AndSlot("crescent past", "long hook"),

	// nuun / symmetry city / talus peaks
	"ricky nuun":   Root(),
	"dimitri nuun": Root(),
	"moosh nuun":   Root("start"),
	"nuun": And("lynna city", Or("currents",
		And("fairies' woods", "ember shooter"))),
	"nuun highlands cave": AndSlot("nuun", Or("dimitri's flute",
		And(Or("ricky nuun", "moosh nuun"), Or("flute", "currents")))),
	"symmetry present": And("nuun", Or("currents", "flute")),
	"symmetry city tree": AndSlot(Or("sword", "punch object",
		And("dimitri's flute", "enter d4")), "seed item", "symmetry present"),
	"symmetry past": And("symmetry present",
		Or("ages", And("break bush safe", "echoes"))),
	"symmetry city brother": AndSlot("symmetry past"),
	"tokkey's composition":  AndSlot("symmetry past", "flippers"),
	"restoration wall": Or("ages",
		And("symmetry past", "currents", "bracelet", "flippers")),
	"patch": And("restoration wall", Or("sword",
		HardOr("shield", "boomerang", "switch hook", "scent seeds", "shovel"))),
	"talus peaks chest": OrSlot("restoration wall"),
	"enter d4":          And("symmetry present", "tuni nut", "patch"),

	// rolling ridge. what a nightmare
	"goron elder": AndSlot("bomb flower", "switch hook",
		Or("feather", "ages")),
	"ridge west past": Or("goron elder",
		And("ridge west present", Or("ages", And("bracelet", "echoes")))),
	"ridge west present": Or("ridge upper present",
		And("switch hook", "currents", Or("feather", "ages")),
		And("currents", "ridge west past")),
	"ridge west cave": AndSlot("ridge west present"),
	"rolling ridge west tree": AndSlot(Or("sword", "punch object"), "seed item",
		"ridge west past"),
	"under moblin keep": AndSlot("ridge west present", "feather",
		"flippers"),
	"defeat great moblin": AndSlot("ridge west present", "pegasus satchel",
		"bracelet"),
	"ridge upper present": Or(And("ridge base present", "switch hook"),
		And("defeat great moblin", "feather")),
	"enter d5": And("crown key", "ridge upper present"),
	"ridge base present": Or("ridge upper present",
		"ridge mid present",
		And("currents", Or("ridge base past east", "ridge base past west"))),
	"enter d6 present":         And("old mermaid key", "ridge base present"),
	"pool in d6 entrance":      AndSlot("ridge base present", "mermaid suit"),
	"goron dance present":      AndSlot("ridge base present"),
	"goron dance, with letter": AndSlot("ridge base past", "goron letter"),
	"ridge mid past": Or(And("ridge base past west", "switch hook"),
		And("ridge upper present", "ages"),
		And("ridge mid present", "ages"),
		And("ridge base past east", "brother emblem", "feather")),
	"ridge mid present": Or(
		And("ridge mid past", "currents"),
		And("ridge base present", "brother emblem",
			Or("switch hook", "jump 3"))),
	"target carts":           And("ridge mid past", "switch hook", "currents"),
	"target carts 1":         AndSlot("target carts"),
	"target carts 2":         AndSlot("target carts"),
	"goron shooting gallery": AndSlot("target carts", "sword"),
	"rolling ridge east tree": AndSlot(Or("sword", "punch object"), "seed item",
		Or("target carts", And("ridge mid present", "ages"),
			And("ridge mid past", "gale satchel"))),
	"ridge base past east": Or("target carts",
		And("lynna city", Or("feather", "ages"), "mermaid suit"),
		And("ridge mid past", "feather", "brother emblem"),
		"rolling ridge east tree",
		And("ridge base present", "ages"),
		And("ridge base past west", Or("flippers", Hard("jump 3")))),
	"ridge base past west": Or(
		And("ridge base present", Or("ages", And("break bush safe", "echoes"))),
		And("ridge base past east", Or("flippers", Hard("bomb jump 2"))),
		"ridge mid past"), // Ledge added to prevent softlocks
	"ridge base past": AndSlot("ridge base past west", "bombs"),
	"enter d6 past": And("mermaid key", "ridge base past west",
		Or("flippers", And("ages", "feather"), Hard("bomb jump 2"))),
	"ridge diamonds past": AndSlot("ridge base past west", "switch hook"),
	"bomb goron head": AndSlot("bombs", Or(
		And("ridge base past west", "switch hook"),
		And("ridge upper present", "ages"))),
	"big bang game":         AndSlot("goronade", "ridge mid present"),
	"ridge NE cave present": AndSlot("ridge mid present"),
	"trade rock brisket": AndSlot("brother emblem", "rock brisket",
		"ridge base present"),
	"trade goron vase": AndSlot("brother emblem", "goron vase",
		"ridge base past east"),
	"trade lava juice":     AndSlot("lava juice", "ridge mid past"),
	"goron's hiding place": AndSlot("ridge west present", "bombs"),
	"ridge base chest":     AndSlot("ridge west present"),
	"goron diamond cave": AndSlot("ridge mid present",
		Or("switch hook", "jump 3")),
	"ridge bush cave": AndSlot("ridge mid past", "switch hook"),

	// zora village / zora seas. only accessible with tune of ages, so no
	// distinctions between past and present are necessary.
	"zora village": And("mermaid suit", "ages", "switch hook"),
	"zora village tree": AndSlot("zora village", "seed item",
		Or("sword", "punch object", And("dimitri's flute", "clean seas"))),
	"zora village present": AndSlot("zora village"),
	"zora palace chest":    AndSlot("zora village"),
	"zora NW cave":         AndSlot("zora village", "bombs", "power glove"),
	"fairies' coast chest": AndSlot("zora village"),
	// in hard logic, farm kills and get a potion off maple
	"king zora":       AndSlot("zora village", Or("syrup", Hard())),
	"library present": AndSlot("zora village", "library key"),
	"library past": AndSlot("zora village", "library key",
		Or("book of seals", "bomb jump 3")),
	"clean seas":           And("zora village", "fairy powder"),
	"zora seas chest":      AndSlot("clean seas"),
	"enter d7":             And("king zora", "clean seas"),
	"fisher's island cave": AndSlot("mermaid suit", "ages", "long hook"),
	"zora's reward":        AndSlot("d7 boss"),

	// sea of storms / sea of no return
	"piratian captain":   AndSlot("lynna city", "mermaid suit", "zora scale"),
	"sea of storms past": AndSlot("lynna city", "mermaid suit", "zora scale"),
	"enter d8": And("crescent past", "tokay eyeball", "kill normal", "break pot",
		"bombs", Or("cane", Hard()), "mermaid suit", "feather"),
	"sea of no return": AndSlot("enter d8", "power glove"),
}
