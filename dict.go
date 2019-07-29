package main

type Harf struct {
	Isolated, Initial, Median, Final rune
	ConnectNext, Tashkeel            bool
}

var space = '\u0020'

var Haruf = map[string]Harf{

	"aN": Harf{
		Isolated: '\u064b',
		Initial:  '\u064b',
		Median:   '\u064b',
		Final:    '\u064b',
		Tashkeel: true},

	"uN": Harf{
		Isolated: '\u064c',
		Initial:  '\u064c',
		Median:   '\u064c',
		Final:    '\u064c',
		Tashkeel: true},

	"iN": Harf{
		Isolated: '\u064d',
		Initial:  '\u064d',
		Median:   '\u064d',
		Final:    '\u064d',
		Tashkeel: true},

	"a": Harf{
		Isolated: '\u064e',
		Initial:  '\u064e',
		Median:   '\u064e',
		Final:    '\u064e',
		Tashkeel: true},

	"u": Harf{
		Isolated: '\u064f',
		Initial:  '\u064f',
		Median:   '\u064f',
		Final:    '\u064f',
		Tashkeel: true},

	"i": Harf{
		Isolated: '\u0650',
		Initial:  '\u0650',
		Median:   '\u0650',
		Final:    '\u0650',
		Tashkeel: true},

	"-": Harf{
		Isolated: '\u0652',
		Initial:  '\u0652',
		Median:   '\u0652',
		Final:    '\u0652',
		Tashkeel: true},

	"~aa": Harf{
		Isolated:    '\ufe81',
		Initial:     '\u0622',
		Median:      '\u0622',
		Final:       '\ufe82',
		Tashkeel:    false,
		ConnectNext: false},
	"'": Harf{
		Isolated:    '\ufe80',
		Initial:     '\u0621',
		Median:      '\u0621',
		Final:       '\u0621',
		ConnectNext: false},

	"'a": Harf{
		Isolated: '\u0623',
		Initial:  '\u0623',
		Median:   '\u0623',
		Final:    '\u0623',
		Tashkeel: false},

	"'i": Harf{
		Isolated: '\ufe89',
		Initial:  '\ufe8b',
		Median:   '\ufe8c',
		Final:    '\ufe8a',
		Tashkeel: false},

	"'w": Harf{
		Isolated: '\ufe85',
		Initial:  '\ufe85',
		Median:   '\ufe85',
		Final:    '\ufe86',
		Tashkeel: false},

	"aa": Harf{
		Isolated:    '\ufe8d',
		Initial:     '\u0627',
		Median:      '\ufe8e',
		Final:       '\ufe8e',
		ConnectNext: false},

	"b": Harf{
		Isolated:    '\ufe8f',
		Initial:     '\ufe91',
		Median:      '\ufe92',
		Final:       '\ufe90',
		ConnectNext: true},

	"t": Harf{
		Isolated:    '\ufe95',
		Initial:     '\ufe97',
		Median:      '\ufe98',
		Final:       '\ufe96',
		ConnectNext: true},

	"th": Harf{
		Isolated:    '\ufe99',
		Initial:     '\ufe9b',
		Median:      '\ufe9c',
		Final:       '\ufe9a',
		ConnectNext: true},

	"j": Harf{
		Isolated:    '\ufe9d',
		Initial:     '\ufe9f',
		Median:      '\ufea0',
		Final:       '\ufe9e',
		ConnectNext: true},

	"H": Harf{
		Isolated:    '\ufea1',
		Initial:     '\ufea3',
		Median:      '\ufea4',
		Final:       '\ufea2',
		ConnectNext: true},

	"kh": Harf{
		Isolated:    '\ufea5',
		Initial:     '\ufea7',
		Median:      '\ufea8',
		Final:       '\ufea6',
		ConnectNext: true},

	"d": Harf{
		Isolated:    '\ufea9',
		Initial:     '\u062f',
		Median:      '\ufeaa',
		Final:       '\ufeaa',
		ConnectNext: false},

	"dh": Harf{
		Isolated:    '\ufeab',
		Initial:     '\u0630',
		Median:      '\ufeac',
		Final:       '\ufeac',
		ConnectNext: false},

	"r": Harf{
		Isolated:    '\ufead',
		Initial:     '\u0631',
		Median:      '\ufeae',
		Final:       '\ufeae',
		ConnectNext: false},

	"z": Harf{
		Isolated:    '\ufeaf',
		Initial:     '\u0632',
		Median:      '\ufeb0',
		Final:       '\ufeb0',
		ConnectNext: false},

	"s": Harf{
		Isolated:    '\ufeb1',
		Initial:     '\ufeb3',
		Median:      '\ufeb4',
		Final:       '\ufeb2',
		ConnectNext: true},

	"sh": Harf{
		Isolated:    '\ufeb5',
		Initial:     '\ufeb7',
		Median:      '\ufeb8',
		Final:       '\ufeb6',
		ConnectNext: true},

	"S": Harf{
		Isolated:    '\ufeb9',
		Initial:     '\ufebb',
		Median:      '\ufebc',
		Final:       '\ufeba',
		ConnectNext: true},

	"D": Harf{
		Isolated:    '\ufebd',
		Initial:     '\ufebf',
		Median:      '\ufec0',
		Final:       '\ufebe',
		ConnectNext: true},

	"T": Harf{
		Isolated:    '\ufec1',
		Initial:     '\ufec3',
		Median:      '\ufec4',
		Final:       '\ufec2',
		ConnectNext: true},

	"Z": Harf{
		Isolated:    '\ufec5',
		Initial:     '\ufec7',
		Median:      '\ufec8',
		Final:       '\ufec6',
		ConnectNext: true},

	"`": Harf{
		Isolated:    '\ufec9',
		Initial:     '\ufecb',
		Median:      '\ufecc',
		Final:       '\ufeca',
		ConnectNext: true},

	"gh": Harf{
		Isolated:    '\ufecd',
		Initial:     '\ufecf',
		Median:      '\ufed0',
		Final:       '\ufece',
		ConnectNext: true},

	"f": Harf{
		Isolated:    '\ufed1',
		Initial:     '\ufed3',
		Median:      '\ufed4',
		Final:       '\ufed2',
		ConnectNext: true},

	"q": Harf{
		Isolated:    '\ufed5',
		Initial:     '\ufed7',
		Median:      '\ufed8',
		Final:       '\ufed6',
		ConnectNext: true},

	"k": Harf{
		Isolated:    '\ufed9',
		Initial:     '\ufedb',
		Median:      '\ufedc',
		Final:       '\ufeda',
		ConnectNext: true},

	"l": Harf{
		Isolated:    '\ufedd',
		Initial:     '\ufedf',
		Median:      '\ufee0',
		Final:       '\ufede',
		ConnectNext: true},

	"m": Harf{
		Isolated:    '\ufee1',
		Initial:     '\ufee3',
		Median:      '\ufee4',
		Final:       '\ufee2',
		ConnectNext: true},

	"n": Harf{
		Isolated:    '\ufee5',
		Initial:     '\ufee7',
		Median:      '\ufee8',
		Final:       '\ufee6',
		ConnectNext: true},

	"h": Harf{
		Isolated:    '\ufee9',
		Initial:     '\ufeeb',
		Median:      '\ufeec',
		Final:       '\ufeea',
		ConnectNext: true},

	"w": Harf{
		Isolated:    '\ufeed',
		Initial:     '\u0648',
		Median:      '\ufeee',
		Final:       '\ufeee',
		ConnectNext: false},

	"y": Harf{
		Isolated:    '\ufbfc',
		Initial:     '\ufbfe',
		Median:      '\ufbff',
		Final:       '\ufbfd',
		ConnectNext: true},

	"ta": Harf{
		Isolated:    '\ufe93',
		Initial:     '\u0629',
		Median:      '\u0629',
		Final:       '\ufe94',
		ConnectNext: false},

	"ae": Harf{
		Isolated:    '\ufeef',
		Initial:     '\u0649',
		Median:      '\ufef0',
		Final:       '\ufef0',
		ConnectNext: false},

	"e": Harf{
		Isolated:    '\ufe8d',
		Initial:     '\u0627',
		Median:      '\ufe8e',
		Final:       '\ufe8e',
		ConnectNext: false},
}
