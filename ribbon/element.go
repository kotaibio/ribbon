package ribbon

type Element struct {
	AtomicNumber      int
	Radius            float64
	IonicRadius       float64
	CovalentRadius    float64
	VanDerWaalsRadius float64
	CrystalRadius     float64
	HexColor          string
	Symbol            string
	Name              string
}

var ElementsBySymbol map[string]Element

func init() {
	ElementsBySymbol = make(map[string]Element)
	for _, e := range Elements {
		ElementsBySymbol[e.Symbol] = e
	}
}

var Elements = []Element{
	{},
	{1, 0.53, 0.25, 0.37, 1.20, 0.10, "#ffffff", "H", "hydrogen"},
	{2, 0.31, 0.31, 0.32, 1.40, 0.00, "#d9ffff", "He", "helium"},
	{3, 1.67, 1.45, 1.34, 1.82, 0.90, "#cc80ff", "Li", "lithium"},
	{4, 1.12, 1.05, 0.90, 1.53, 0.41, "#c2ff00", "Be", "beryllium"},
	{5, 0.87, 0.85, 0.82, 1.92, 0.25, "#ffb5b5", "B", "boron"},
	{6, 0.67, 0.70, 0.77, 1.70, 0.29, "#909090", "C", "carbon"},
	{7, 0.56, 0.65, 0.75, 1.55, 0.30, "#3050f8", "N", "nitrogen"},
	{8, 0.48, 0.60, 0.73, 1.52, 1.21, "#ff0d0d", "O", "oxygen"},
	{9, 0.42, 0.50, 0.71, 1.47, 1.19, "#90e050", "F", "fluorine"},
	{10, 0.38, 0.38, 0.69, 1.54, 0.00, "#b3e3f5", "Ne", "neon"},
	{11, 1.90, 1.80, 1.54, 2.27, 1.16, "#ab5cf2", "Na", "sodium"},
	{12, 1.45, 1.50, 1.30, 1.73, 0.86, "#8aff00", "Mg", "magnesium"},
	{13, 1.18, 1.25, 1.18, 1.84, 0.53, "#bfa6a6", "Al", "aluminium"},
	{14, 1.11, 1.10, 1.11, 2.10, 0.40, "#f0c8a0", "Si", "silicon"},
	{15, 0.98, 1.00, 1.06, 1.80, 0.31, "#ff8000", "P", "phosphorus"},
	{16, 0.88, 1.00, 1.02, 1.80, 0.43, "#ffff30", "S", "sulfur"},
	{17, 0.79, 1.00, 0.99, 1.75, 1.67, "#1ff01f", "Cl", "chlorine"},
	{18, 0.71, 0.71, 0.97, 1.88, 0.00, "#80d1e3", "Ar", "argon"},
	{19, 2.43, 2.20, 1.96, 2.75, 1.52, "#8f40d4", "K", "potassium"},
	{20, 1.94, 1.80, 1.74, 2.31, 1.14, "#3dff00", "Ca", "calcium"},
	{21, 1.84, 1.60, 1.44, 2.11, 0.89, "#e6e6e6", "Sc", "scandium"},
	{22, 1.76, 1.40, 1.36, 0.00, 0.75, "#bfc2c7", "Ti", "titanium"},
	{23, 1.71, 1.35, 1.25, 0.00, 0.68, "#a6a6ab", "V", "vanadium"},
	{24, 1.66, 1.40, 1.27, 0.00, 0.76, "#8a99c7", "Cr", "chromium"},
	{25, 1.61, 1.40, 1.39, 0.00, 0.81, "#9c7ac7", "Mn", "manganese"},
	{26, 1.56, 1.40, 1.25, 0.00, 0.69, "#e06633", "Fe", "iron"},
	{27, 1.52, 1.35, 1.26, 0.00, 0.54, "#f090a0", "Co", "cobalt"},
	{28, 1.49, 1.35, 1.21, 1.63, 0.70, "#50d050", "Ni", "nickel"},
	{29, 1.45, 1.35, 1.38, 1.40, 0.71, "#c88033", "Cu", "copper"},
	{30, 1.42, 1.35, 1.31, 1.39, 0.74, "#7d80b0", "Zn", "zinc"},
	{31, 1.36, 1.30, 1.26, 1.87, 0.76, "#c28f8f", "Ga", "gallium"},
	{32, 1.25, 1.25, 1.22, 2.11, 0.53, "#668f8f", "Ge", "germanium"},
	{33, 1.14, 1.15, 1.19, 1.85, 0.72, "#bd80e3", "As", "arsenic"},
	{34, 1.03, 1.15, 1.16, 1.90, 0.56, "#ffa100", "Se", "selenium"},
	{35, 0.94, 1.15, 1.14, 1.85, 1.82, "#a62929", "Br", "bromine"},
	{36, 0.88, 0.88, 1.10, 2.02, 0.00, "#5cb8d1", "Kr", "krypton"},
	{37, 2.65, 2.35, 2.11, 3.03, 1.66, "#702eb0", "Rb", "rubidium"},
	{38, 2.19, 2.00, 1.92, 2.49, 1.32, "#00ff00", "Sr", "strontium"},
	{39, 2.12, 1.85, 1.62, 0.00, 1.04, "#94ffff", "Y", "yttrium"},
	{40, 2.06, 1.55, 1.48, 0.00, 0.86, "#94e0e0", "Zr", "zirconium"},
	{41, 1.98, 1.45, 1.37, 0.00, 0.78, "#73c2c9", "Nb", "niobium"},
	{42, 1.90, 1.45, 1.45, 0.00, 0.79, "#54b5b5", "Mo", "molybdenum"},
	{43, 1.83, 1.35, 1.56, 0.00, 0.79, "#3b9e9e", "Tc", "technetium"},
	{44, 1.78, 1.30, 1.26, 0.00, 0.82, "#248f8f", "Ru", "ruthenium"},
	{45, 1.73, 1.35, 1.35, 0.00, 0.81, "#0a7d8c", "Rh", "rhodium"},
	{46, 1.69, 1.40, 1.31, 1.63, 0.78, "#006985", "Pd", "palladium"},
	{47, 1.65, 1.60, 1.53, 1.72, 1.29, "#c0c0c0", "Ag", "silver"},
	{48, 1.61, 1.55, 1.48, 1.58, 0.92, "#ffd98f", "Cd", "cadmium"},
	{49, 1.56, 1.55, 1.44, 1.93, 0.94, "#a67573", "In", "indium"},
	{50, 1.45, 1.45, 1.41, 2.17, 0.69, "#668080", "Sn", "tin"},
	{51, 1.33, 1.45, 1.38, 2.06, 0.90, "#9e63b5", "Sb", "antimony"},
	{52, 1.23, 1.40, 1.35, 2.06, 1.11, "#d47a00", "Te", "tellurium"},
	{53, 1.15, 1.40, 1.33, 1.98, 2.06, "#940094", "I", "iodine"},
	{54, 1.08, 1.08, 1.30, 2.16, 0.62, "#429eb0", "Xe", "xenon"},
	{55, 2.98, 2.60, 2.25, 3.43, 1.81, "#57178f", "Cs", "caesium"},
	{56, 2.53, 2.15, 1.98, 2.68, 1.49, "#00c900", "Ba", "barium"},
	{57, 1.95, 1.95, 1.69, 0.00, 1.36, "#70d4ff", "La", "lanthanum"},
	{58, 1.85, 1.85, 0.00, 0.00, 1.15, "#ffffc7", "Ce", "cerium"},
	{59, 2.47, 1.85, 0.00, 0.00, 1.32, "#d9ffc7", "Pr", "praseodymium"},
	{60, 2.06, 1.85, 0.00, 0.00, 1.30, "#c7ffc7", "Nd", "neodymium"},
	{61, 2.05, 1.85, 0.00, 0.00, 1.28, "#a3ffc7", "Pm", "promethium"},
	{62, 2.38, 1.85, 0.00, 0.00, 1.10, "#8fffc7", "Sm", "samarium"},
	{63, 2.31, 1.85, 0.00, 0.00, 1.31, "#61ffc7", "Eu", "europium"},
	{64, 2.33, 1.80, 0.00, 0.00, 1.08, "#45ffc7", "Gd", "gadolinium"},
	{65, 2.25, 1.75, 0.00, 0.00, 1.18, "#30ffc7", "Tb", "terbium"},
	{66, 2.28, 1.75, 0.00, 0.00, 1.05, "#1fffc7", "Dy", "dysprosium"},
	{67, 2.26, 1.75, 0.00, 0.00, 1.04, "#00ff9c", "Ho", "holmium"},
	{68, 2.26, 1.75, 0.00, 0.00, 1.03, "#00e675", "Er", "erbium"},
	{69, 2.22, 1.75, 0.00, 0.00, 1.02, "#00d452", "Tm", "thulium"},
	{70, 2.22, 1.75, 0.00, 0.00, 1.13, "#00bf38", "Yb", "ytterbium"},
	{71, 2.17, 1.75, 1.60, 0.00, 1.00, "#00ab24", "Lu", "lutetium"},
	{72, 2.08, 1.55, 1.50, 0.00, 0.85, "#4dc2ff", "Hf", "hafnium"},
	{73, 2.00, 1.45, 1.38, 0.00, 0.78, "#4da6ff", "Ta", "tantalum"},
	{74, 1.93, 1.35, 1.46, 0.00, 0.74, "#2194d6", "W", "tungsten"},
	{75, 1.88, 1.35, 1.59, 0.00, 0.77, "#267dab", "Re", "rhenium"},
	{76, 1.85, 1.30, 1.28, 0.00, 0.77, "#266696", "Os", "osmium"},
	{77, 1.80, 1.35, 1.37, 0.00, 0.77, "#175487", "Ir", "iridium"},
	{78, 1.77, 1.35, 1.28, 1.75, 0.74, "#d0d0e0", "Pt", "platinum"},
	{79, 1.74, 1.35, 1.44, 1.66, 1.51, "#ffd123", "Au", "gold"},
	{80, 1.71, 1.50, 1.49, 1.55, 0.83, "#b8b8d0", "Hg", "mercury"},
	{81, 1.56, 1.90, 1.48, 1.96, 1.03, "#a6544d", "Tl", "thallium"},
	{82, 1.54, 1.80, 1.47, 2.02, 1.49, "#575961", "Pb", "lead"},
	{83, 1.43, 1.60, 1.46, 2.07, 1.17, "#9e4fb5", "Bi", "bismuth"},
	{84, 1.35, 1.90, 0.00, 1.97, 1.08, "#ab5c00", "Po", "polonium"},
	{85, 1.27, 1.27, 0.00, 2.02, 0.76, "#754f45", "At", "astatine"},
	{86, 1.20, 1.20, 1.45, 0.00, 0.00, "#428296", "Rn", "radon"},
	{87, 0.00, 0.00, 0.00, 3.48, 1.94, "#420066", "Fr", "francium"},
	{88, 0.00, 2.15, 0.00, 2.83, 1.62, "#007d00", "Ra", "radium"},
	{89, 1.95, 1.95, 0.00, 0.00, 1.26, "#70abfa", "Ac", "actinium"},
	{90, 1.80, 1.80, 0.00, 0.00, 1.19, "#00baff", "Th", "thorium"},
	{91, 1.80, 1.80, 0.00, 0.00, 1.09, "#00a1ff", "Pa", "protactinium"},
	{92, 1.75, 1.75, 0.00, 1.86, 0.87, "#008fff", "U", "uranium"},
	{93, 1.75, 1.75, 0.00, 0.00, 0.00, "#0080ff", "Np", "neptunium"},
	{94, 1.75, 1.75, 0.00, 0.00, 1.00, "#006bff", "Pu", "plutonium"},
	{95, 1.75, 1.75, 0.00, 0.00, 1.12, "#545cf2", "Am", "americium"},
	{96, 0.00, 0.00, 0.00, 0.00, 1.11, "#785ce3", "Cm", "curium"},
}
