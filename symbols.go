package termui

const (
	DOT      = '•'
	ELLIPSES = '…'

	UP_ARROW    = '▲'
	DOWN_ARROW  = '▼'
	LEFT_ARROW  = '◀'
	RIGHT_ARROW = '▶'

	COLLAPSED        = '+'
	EXPANDED         = '−'
	EMPTY      int32 = ' '
	RUR        int32 = '₽'
	SHEKEL     int32 = '₪'
	LAOS_KIP   int32 = '₭'
	LIGHTCOIN  int32 = 'Ł'
	SYMBOL_YAT int32 = 'Ѣ'

	CHECK_DEFAULT   int32 = '✓'
	DISABLE_DEFAULT int32 = '✖'

	// греческие буквы

	GR_TETA  int32 = 'Θ'
	GR_DELTA int32 = 'Δ'
	GR_KSI   int32 = 'Ξ'
	GR_PSI   int32 = 'Ψ'
	GR_OMEGA int32 = 'Ω'
	GR_SIGMA int32 = 'Σ'

	// римские цифры

	ROME_1  int32 = 'Ⅰ'
	ROME_2  int32 = 'Ⅱ'
	ROME_3  int32 = 'Ⅲ'
	ROME_4  int32 = 'Ⅳ'
	ROME_5  int32 = 'Ⅴ'
	ROME_6  int32 = 'Ⅵ'
	ROME_7  int32 = 'Ⅶ'
	ROME_8  int32 = 'Ⅷ'
	ROME_9  int32 = 'Ⅸ'
	ROME_10 int32 = 'Ⅹ'
	ROME_11 int32 = 'Ⅺ'
	ROME_12 int32 = 'Ⅻ'

	ROME_50   int32 = 'Ⅼ'
	ROME_100  int32 = 'Ⅽ'
	ROME_500  int32 = 'Ⅾ'
	ROME_1000 int32 = 'Ⅿ'

	NUMBER_BOLD_0 int32 = '𝟬'
	NUMBER_BOLD_1 int32 = '𝟭'
	NUMBER_BOLD_2 int32 = '𝟮'
	NUMBER_BOLD_3 int32 = '𝟯'
	NUMBER_BOLD_4 int32 = '𝟰'
	NUMBER_BOLD_5 int32 = '𝟱'
	NUMBER_BOLD_6 int32 = '𝟲'
	NUMBER_BOLD_7 int32 = '𝟳'
	NUMBER_BOLD_8 int32 = '𝟴'
	NUMBER_BOLD_9 int32 = '𝟵'

	/*
		color symbols
	*/

	CS_FOOTPRINT        int32 = '👣'
	CS_FOOTPRINT_ANIMAL int32 = '🐾'
	CS_CLOVER_4_LIST    int32 = '🍀'
	CS_ANGER            int32 = '💢'
	CS_EYES             int32 = '👀'
	CS_HOURGLASS        int32 = '⌛'
	CS_COFFEE           int32 = '☕'
	CS_SHOWING          int32 = '🌨'
	CS_FIRE             int32 = '🔥'
	CS_FIRE_ENGINE      int32 = '🚒'
	CS_ROCKET           int32 = '🚀'
	CS_DIZZINESS        int32 = '💫'
	CS_BUMP             int32 = '💥'
	CS_NOTES            int32 = '🎶'
	CS_BEEPER           int32 = '📯'

	/*
		triangles
	*/
	TRIA_TOP_L    int32 = '◸'
	TRIA_TOP_R    int32 = '◹'
	TRIA_BOTTOM_L int32 = '◺'
	TRIA_BOTTOM_R int32 = '◿'

	/*
		i-chin
	*/
	I_CHIN_1 int32 = '☰'
	I_CHIN_2 int32 = '☱'
	I_CHIN_3 int32 = '☲'
	I_CHIN_4 int32 = '☳'
	I_CHIN_5 int32 = '☴'
	I_CHIN_6 int32 = '☵'
	I_CHIN_7 int32 = '☶'
	I_CHIN_8 int32 = '☷'
)

var (
	BARS = [...]rune{' ', '▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

	SHADED_BLOCKS = [...]rune{' ', '░', '▒', '▓', '█'}

	IRREGULAR_BLOCKS = [...]rune{
		' ', '▘', '▝', '▀', '▖', '▌', '▞', '▛',
		'▗', '▚', '▐', '▜', '▄', '▙', '▟', '█',
	}

	BRAILLE_OFFSET = '\u2800'
	BRAILLE        = [4][2]rune{
		{'\u0001', '\u0008'},
		{'\u0002', '\u0010'},
		{'\u0004', '\u0020'},
		{'\u0040', '\u0080'},
	}

	DOUBLE_BRAILLE = map[[2]int]rune{
		[2]int{0, 0}: '⣀',
		[2]int{0, 1}: '⡠',
		[2]int{0, 2}: '⡐',
		[2]int{0, 3}: '⡈',

		[2]int{1, 0}: '⢄',
		[2]int{1, 1}: '⠤',
		[2]int{1, 2}: '⠔',
		[2]int{1, 3}: '⠌',

		[2]int{2, 0}: '⢂',
		[2]int{2, 1}: '⠢',
		[2]int{2, 2}: '⠒',
		[2]int{2, 3}: '⠊',

		[2]int{3, 0}: '⢁',
		[2]int{3, 1}: '⠡',
		[2]int{3, 2}: '⠑',
		[2]int{3, 3}: '⠉',
	}

	SINGLE_BRAILLE_LEFT  = [4]rune{'\u2840', '⠄', '⠂', '⠁'}
	SINGLE_BRAILLE_RIGHT = [4]rune{'\u2880', '⠠', '⠐', '⠈'}
)
