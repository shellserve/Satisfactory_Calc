package tui

import (
	"github.com/charmbracelet/lipgloss"
)

type ColorSet struct {
	// Blacks
	AILimiterBlack            string
	BatteryBlack              string
	CoalBlack                 string
	ComputerBlack             string
	CokeBlack                 string
	EncasedPlutoniumCellBlack string
	IronRodBlack              string
	OilBlack                  string
	RubberBlack               string
	SteelBlack                string
	SteelBeamBlack            string
	SteelPipeBlack            string

	// Whites & Grays
	AluminumCasingGray      string
	AluminumIngotGray       string
	AluminumScrapGray       string
	AluminumSheetGray       string
	AluminaSolutionWhite    string
	BatteryGray             string
	CateriumOreGray         string
	ConcreteWhite           string
	GaseliskGray            string
	HeatSinkGray            string
	HeavyModularFrameBlack  string
	HeavyModularFrameGray   string
	IronGray                string
	IronIngotGray           string
	IronOreGray             string
	IronPlateGray           string
	NitricAcidWhite         string
	PlutoniumPelletGray     string
	RawQuartzGray           string
	ReinforcedIronPlateGray string
	SilicaWhite             string
	StatorGray              string
	SulphurGray             string

	// Reds, Yellows & Browns
	AluminumScrapOrange        string
	BauxiteReddish             string
	CateriumOreYellow          string
	CateriumIngotYellow        string
	ClusteriskOrange           string
	ClusteriskRed              string
	ComputerOrange             string
	ConcreteBrownish           string
	CopperIngotReddish         string
	CopperOreOrange            string
	CopperSheetReddish         string
	CrystalOscillatorYellow    string
	EncasedIndustrialBeamRed   string
	EncasedPlutoniumCellYellow string
	FicsitOrange               string
	FuelOrange                 string
	FusedModularFrameOrange    string
	HeatSinkReddish            string
	IronOreReddish             string
	LimestoneBrownish          string
	NobeliskRed                string
	NobeliskYellow             string
	NonFissileUraniumYellow    string
	NukeliskYellow             string
	PlutoniumFuelRodYellow     string
	QuickwireYellow            string
	RubberYellow               string
	SteelBeamRed               string
	SteelPipeReddish           string
	SulphurYellow              string
	SulphuricAcidYellow        string
	SupercomputerOrange        string
	TurbofuelRed               string
	UraniumCellYellow          string
	UraniumFuelRodYellow       string
	UraniumWasteYellow         string
	WireBrownish               string
	WireRed                    string

	// Greens, Blues & Purples
	AILimiterTeal            string
	CircuitBoardGreen        string
	CopperOreGreenish        string
	CopperSheetBlue          string
	CrystalOscillatorPinkish string
	GaseliskGreen            string
	HeavyOilResidueHORPurple string
	LiquidBiofuelGreen       string
	NonFissileUraniumGreen   string
	NukeliskGreenish         string
	PlasticBlue              string
	PlutoniumFeulRodBlue     string
	PlutoniumPelletBlue      string
	PolymerResinBlue         string
	PulseliskTeal            string
	QuartzCrystalBlue        string
	QuartzCtystalPink        string
	QuickwireDarkBlue        string
	RawQuartzPink            string
	RotorBlue                string
	ScrewBlue                string
	SilicaTeal               string
	UraniumFuelRodGreen      string
	UraniumGreen             string
	UraniumWasteGreen        string
	WaterBlue                string
}

var Colors = ColorSet{
	// Blacks
	AILimiterBlack:            "#060606",
	BatteryBlack:              "#1B1C30",
	CoalBlack:                 "#0B0B19",
	ComputerBlack:             "#1C1C1C",
	CokeBlack:                 "#030309",
	EncasedPlutoniumCellBlack: "#191919",
	IronRodBlack:              "#0D0D0F",
	OilBlack:                  "#0A090F",
	RubberBlack:               "#202020",
	SteelBlack:                "#0A090F",
	SteelBeamBlack:            "#0C0909",
	SteelPipeBlack:            "#222020",

	// Whites & Grays
	AluminumCasingGray:      "#C7C9CB",
	AluminumIngotGray:       "#D2D3D4",
	AluminumScrapGray:       "#BCC0C9",
	AluminumSheetGray:       "#9EA0A2",
	AluminaSolutionWhite:    "#DDDEDF",
	BatteryGray:             "#BCBBC4",
	CateriumOreGray:         "#A7AAB7",
	ConcreteWhite:           "#EEEBEA",
	GaseliskGray:            "#B3B8BC",
	HeatSinkGray:            "#B0B3B0",
	HeavyModularFrameBlack:  "#191919",
	HeavyModularFrameGray:   "#989FA9",
	IronGray:                "#69717C",
	IronIngotGray:           "#989A9D",
	IronOreGray:             "#989FA9",
	IronPlateGray:           "#BCBEC1",
	NitricAcidWhite:         "#F7FAD7",
	PlutoniumPelletGray:     "#5A5552",
	RawQuartzGray:           "#767676",
	ReinforcedIronPlateGray: "#697082",
	SilicaWhite:             "#D8DDE7",
	StatorGray:              "#505054",
	SulphurGray:             "#867F7D",

	// Reds, Yellows & Browns
	AluminumScrapOrange:        "#CA852C",
	BauxiteReddish:             "#CD7660",
	CateriumOreYellow:          "#E2B148",
	CateriumIngotYellow:        "#CCA566",
	ClusteriskOrange:           "#CC7C00",
	ClusteriskRed:              "#4B3131",
	ComputerOrange:             "#89552F",
	ConcreteBrownish:           "#D3BEA4",
	CopperIngotReddish:         "#A56355",
	CopperOreOrange:            "#BD4C39",
	CopperSheetReddish:         "#8C554D",
	CrystalOscillatorYellow:    "#E5B82A",
	EncasedIndustrialBeamRed:   "#893F3F",
	EncasedPlutoniumCellYellow: "#DDD839",
	FicsitOrange:               "#FA9549",
	FuelOrange:                 "#D47615",
	FusedModularFrameOrange:    "#FDB164",
	HeatSinkReddish:            "#76402C",
	IronOreReddish:             "#8E5C5C",
	LimestoneBrownish:          "#C8BFA7",
	NobeliskRed:                "#901400",
	NobeliskYellow:             "#FFE500",
	NonFissileUraniumYellow:    "#E1D529",
	NukeliskYellow:             "#DAC500",
	PlutoniumFuelRodYellow:     "#E5D04A",
	QuickwireYellow:            "#CDBD8B",
	RubberYellow:               "#D1B84F",
	SteelBeamRed:               "#4B2323",
	SteelPipeReddish:           "#6C3936",
	SulphurYellow:              "#FCDC48",
	SulphuricAcidYellow:        "#FFF03A",
	SupercomputerOrange:        "#EEBA72",
	TurbofuelRed:               "#A10000",
	UraniumCellYellow:          "#DAD02B",
	UraniumFuelRodYellow:       "#BCA626",
	UraniumWasteYellow:         "#F9F500",
	WireBrownish:               "#A69481",
	WireRed:                    "#A95236",

	// Greens, Blues & Purples
	AILimiterTeal:            "#76A69F",
	CircuitBoardGreen:        "#588B55",
	CopperOreGreenish:        "#70978A",
	CopperSheetBlue:          "#13239B",
	CrystalOscillatorPinkish: "#B9ADBC",
	GaseliskGreen:            "#4F6F27",
	HeavyOilResidueHORPurple: "#AE1CD7",
	LiquidBiofuelGreen:       "#133405",
	NonFissileUraniumGreen:   "#255D2E",
	NukeliskGreenish:         "#4D5954",
	PlasticBlue:              "#3091E6",
	PlutoniumFeulRodBlue:     "#B5EAF3",
	PlutoniumPelletBlue:      "#00B9FB",
	PolymerResinBlue:         "#0D0087",
	PulseliskTeal:            "#7E9EAA",
	QuartzCrystalBlue:        "#3948BF",
	QuartzCtystalPink:        "#D97FEE",
	QuickwireDarkBlue:        "#181820",
	RawQuartzPink:            "#F177B5",
	RotorBlue:                "#2F3354",
	ScrewBlue:                "#2443A0",
	SilicaTeal:               "#9BC0CE",
	UraniumFuelRodGreen:      "#5BD37A",
	UraniumGreen:             "#88D288",
	UraniumWasteGreen:        "#3CC142",
	WaterBlue:                "#1662AD",
}

// generic styles
var (
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
	SubTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Colors.FicsitOrange))
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color(Colors.BauxiteReddish)).Render(" • ")
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(Colors.WaterBlue))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Colors.RotorBlue))
)
