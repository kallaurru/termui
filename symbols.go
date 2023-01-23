package termui

const (
	DOT      = '•'
	ELLIPSES = '…'

	UP_ARROW    = '▲'
	DOWN_ARROW  = '▼'
	LEFT_ARROW  = '◀'
	RIGHT_ARROW = '▶'

	COLLAPSED       = '+'
	EXPANDED        = '−'
	EMPTY     int32 = ' '
	RUR       int32 = '\u20b6'
	SHEKEL    int32 = '\u20aa'

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
