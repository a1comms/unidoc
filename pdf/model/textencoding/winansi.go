/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package textencoding

import (
	"github.com/unidoc/unidoc/common"
	"github.com/unidoc/unidoc/pdf/core"
)

func splitWords(raw string, encoder TextEncoder) []string {
	runes := []rune(raw)

	words := []string{}

	startsAt := 0
	for idx, code := range runes {
		glyph, found := encoder.RuneToGlyphName(code)
		if !found {
			common.Log.Debug("Glyph not found for code: %s\n", string(code))
			continue
		}

		if glyph == "space" {
			word := runes[startsAt:idx]
			words = append(words, string(word))
			startsAt = idx + 1
		}
	}

	word := runes[startsAt:]
	if len(word) > 0 {
		words = append(words, string(word))
	}

	return words
}

type WinAnsiEncoder struct {
}

func NewWinAnsiTextEncoder() WinAnsiEncoder {
	encoder := WinAnsiEncoder{}
	return encoder
}

func (winenc WinAnsiEncoder) ToPdfObject() core.PdfObject {
	return core.MakeName("WinAnsiEncoding")
}

// Convert utf8 runes to WinAnsiEncoded encoded string (series of char codes).
func (winenc WinAnsiEncoder) Encode(raw string) string {
	encoded := []byte{}
	for _, rune := range raw {
		if code, has := utf8ToWinAnsiEncodingMap[rune]; has {
			encoded = append(encoded, code)
		}
	}

	return string(encoded)
}

func (winenc WinAnsiEncoder) RuneToGlyphName(val rune) (string, bool) {
	code, found := winenc.RuneToCharcode(val)
	if !found {
		return "", false
	}

	glyph, found := winenc.CharcodeToGlyphName(code)
	if !found {
		return "", false
	}

	return glyph, true
}

func (winenc WinAnsiEncoder) CharcodeToGlyphName(code byte) (string, bool) {
	glyph, has := winAnsiEncodingGlyphMap[code]
	if !has {
		return "", false
	}
	return glyph, true
}

func (winenc WinAnsiEncoder) GlyphNameToCharcode(glyph string) (byte, bool) {
	for code, name := range winAnsiEncodingGlyphMap {
		if name == glyph {
			return code, true
		}
	}

	// Not found.
	return 0, false
}

// Convert UTF-8 rune to character code.  If applicable.
func (winenc WinAnsiEncoder) RuneToCharcode(val rune) (byte, bool) {
	code, has := utf8ToWinAnsiEncodingMap[val]
	if !has {
		return 0, false
	}
	return code, true
}

func (winenc WinAnsiEncoder) CharcodeToRune(charcode byte) (rune, bool) {
	val, has := winAnsiEncodingToUtf8Map[charcode]
	if !has {
		return 0, false
	}

	return val, true
}

// WinAnsiEncoding.

// Convert a UTF8 string to WinAnsiEncoding byte string.
func utf8ToWinAnsiEncoding(strUtf8 string) string {
	encoded := []byte{}
	for _, rune := range strUtf8 {
		if code, has := utf8ToWinAnsiEncodingMap[rune]; has {
			encoded = append(encoded, code)
		}
	}
	return string(encoded)
}

// Maps to enable conversion of WinAnsiEncoding character codes to glyphs, utf8 and vice versa.
var winAnsiEncodingGlyphMap = map[byte]string{
	32:  "space",
	33:  "exclam",
	34:  "quotedbl",
	35:  "numbersign",
	36:  "dollar",
	37:  "percent",
	38:  "ampersand",
	39:  "quotesingle",
	40:  "parenleft",
	41:  "parenright",
	42:  "asterisk",
	43:  "plus",
	44:  "comma",
	45:  "hyphen",
	46:  "period",
	47:  "slash",
	48:  "zero",
	49:  "one",
	50:  "two",
	51:  "three",
	52:  "four",
	53:  "five",
	54:  "six",
	55:  "seven",
	56:  "eight",
	57:  "nine",
	58:  "colon",
	59:  "semicolon",
	60:  "less",
	61:  "equal",
	62:  "greater",
	63:  "question",
	64:  "at",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "Q",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "bracketleft",
	92:  "backslash",
	93:  "bracketright",
	94:  "asciicircum",
	95:  "underscore",
	96:  "grave",
	97:  "a",
	98:  "b",
	99:  "c",
	100: "d",
	101: "e",
	102: "f",
	103: "g",
	104: "h",
	105: "i",
	106: "j",
	107: "k",
	108: "l",
	109: "m",
	110: "n",
	111: "o",
	112: "p",
	113: "q",
	114: "r",
	115: "s",
	116: "t",
	117: "u",
	118: "v",
	119: "w",
	120: "x",
	121: "y",
	122: "z",
	123: "braceleft",
	124: "bar",
	125: "braceright",
	126: "asciitilde",
	127: "bullet",
	128: "Euro",
	129: "bullet",
	130: "quotesinglbase",
	131: "florin",
	132: "quotedblbase",
	133: "ellipsis",
	134: "dagger",
	135: "daggerdbl",
	136: "circumflex",
	137: "perthousand",
	138: "Scaron",
	139: "guilsinglleft",
	140: "OE",
	141: "bullet",
	142: "Zcaron",
	143: "bullet",
	144: "bullet",
	145: "quoteleft",
	146: "quoteright",
	147: "quotedblleft",
	148: "quotedblright",
	149: "bullet",
	150: "endash",
	151: "emdash",
	152: "tilde",
	153: "trademark",
	154: "scaron",
	155: "guilsinglright",
	156: "oe",
	157: "bullet",
	158: "zcaron",
	159: "Ydieresis",
	160: "space",
	161: "exclamdown",
	162: "cent",
	163: "sterling",
	164: "currency",
	165: "yen",
	166: "brokenbar",
	167: "section",
	168: "dieresis",
	169: "copyright",
	170: "ordfeminine",
	171: "guillemotleft",
	172: "logicalnot",
	173: "hyphen",
	174: "registered",
	175: "macron",
	176: "degree",
	177: "plusminus",
	178: "twosuperior",
	179: "threesuperior",
	180: "acute",
	181: "mu",
	182: "paragraph",
	183: "periodcentered",
	184: "cedilla",
	185: "onesuperior",
	186: "ordmasculine",
	187: "guillemotright",
	188: "onequarter",
	189: "onehalf",
	190: "threequarters",
	191: "questiondown",
	192: "Agrave",
	193: "Aacute",
	194: "Acircumflex",
	195: "Atilde",
	196: "Adieresis",
	197: "Aring",
	198: "AE",
	199: "Ccedilla",
	200: "Egrave",
	201: "Eacute",
	202: "Ecircumflex",
	203: "Edieresis",
	204: "Igrave",
	205: "Iacute",
	206: "Icircumflex",
	207: "Idieresis",
	208: "Eth",
	209: "Ntilde",
	210: "Ograve",
	211: "Oacute",
	212: "Ocircumflex",
	213: "Otilde",
	214: "Odieresis",
	215: "multiply",
	216: "Oslash",
	217: "Ugrave",
	218: "Uacute",
	219: "Ucircumflex",
	220: "Udieresis",
	221: "Yacute",
	222: "Thorn",
	223: "germandbls",
	224: "agrave",
	225: "aacute",
	226: "acircumflex",
	227: "atilde",
	228: "adieresis",
	229: "aring",
	230: "ae",
	231: "ccedilla",
	232: "egrave",
	233: "eacute",
	234: "ecircumflex",
	235: "edieresis",
	236: "igrave",
	237: "iacute",
	238: "icircumflex",
	239: "idieresis",
	240: "eth",
	241: "ntilde",
	242: "ograve",
	243: "oacute",
	244: "ocircumflex",
	245: "otilde",
	246: "odieresis",
	247: "divide",
	248: "oslash",
	249: "ugrave",
	250: "uacute",
	251: "ucircumflex",
	252: "udieresis",
	253: "yacute",
	254: "thorn",
	255: "ydieresis",
}

var winAnsiEncodingToUtf8Map = map[byte]rune{
	32:  '\u0020',
	33:  '\u0021',
	34:  '\u0022',
	35:  '\u0023',
	36:  '\u0024',
	37:  '\u0025',
	38:  '\u0026',
	39:  '\u0027',
	40:  '\u0028',
	41:  '\u0029',
	42:  '\u002a',
	43:  '\u002b',
	44:  '\u002c',
	45:  '\u002d',
	46:  '\u002e',
	47:  '\u002f',
	48:  '\u0030',
	49:  '\u0031',
	50:  '\u0032',
	51:  '\u0033',
	52:  '\u0034',
	53:  '\u0035',
	54:  '\u0036',
	55:  '\u0037',
	56:  '\u0038',
	57:  '\u0039',
	58:  '\u003a',
	59:  '\u003b',
	60:  '\u003c',
	61:  '\u003d',
	62:  '\u003e',
	63:  '\u003f',
	64:  '\u0040',
	65:  '\u0041',
	66:  '\u0042',
	67:  '\u0043',
	68:  '\u0044',
	69:  '\u0045',
	70:  '\u0046',
	71:  '\u0047',
	72:  '\u0048',
	73:  '\u0049',
	74:  '\u004a',
	75:  '\u004b',
	76:  '\u004c',
	77:  '\u004d',
	78:  '\u004e',
	79:  '\u004f',
	80:  '\u0050',
	81:  '\u0051',
	82:  '\u0052',
	83:  '\u0053',
	84:  '\u0054',
	85:  '\u0055',
	86:  '\u0056',
	87:  '\u0057',
	88:  '\u0058',
	89:  '\u0059',
	90:  '\u005a',
	91:  '\u005b',
	92:  '\u005c',
	93:  '\u005d',
	94:  '\u005e',
	95:  '\u005f',
	96:  '\u0060',
	97:  '\u0061',
	98:  '\u0062',
	99:  '\u0063',
	100: '\u0064',
	101: '\u0065',
	102: '\u0066',
	103: '\u0067',
	104: '\u0068',
	105: '\u0069',
	106: '\u006a',
	107: '\u006b',
	108: '\u006c',
	109: '\u006d',
	110: '\u006e',
	111: '\u006f',
	112: '\u0070',
	113: '\u0071',
	114: '\u0072',
	115: '\u0073',
	116: '\u0074',
	117: '\u0075',
	118: '\u0076',
	119: '\u0077',
	120: '\u0078',
	121: '\u0079',
	122: '\u007a',
	123: '\u007b',
	124: '\u007c',
	125: '\u007d',
	126: '\u007e',
	127: '\u2022',
	128: '\u20ac',
	129: '\u2022',
	130: '\u201a',
	131: '\u0192',
	132: '\u201e',
	133: '\u2026',
	134: '\u2020',
	135: '\u2021',
	136: '\u02c6',
	137: '\u2030',
	138: '\u0160',
	139: '\u2039',
	140: '\u0152',
	141: '\u2022',
	142: '\u017d',
	143: '\u2022',
	144: '\u2022',
	145: '\u2018',
	146: '\u2019',
	147: '\u201c',
	148: '\u201d',
	149: '\u2022',
	150: '\u2013',
	151: '\u2014',
	152: '\u02dc',
	153: '\u2122',
	154: '\u0161',
	155: '\u203a',
	156: '\u0153',
	157: '\u2022',
	158: '\u017e',
	159: '\u0178',
	160: '\u0020',
	161: '\u00a1',
	162: '\u00a2',
	163: '\u00a3',
	164: '\u00a4',
	165: '\u00a5',
	166: '\u00a6',
	167: '\u00a7',
	168: '\u00a8',
	169: '\u00a9',
	170: '\u00aa',
	171: '\u00ab',
	172: '\u00ac',
	173: '\u002d',
	174: '\u00ae',
	175: '\u00af',
	176: '\u00b0',
	177: '\u00b1',
	178: '\u00b2',
	179: '\u00b3',
	180: '\u00b4',
	181: '\u00b5',
	182: '\u00b6',
	183: '\u00b7',
	184: '\u00b8',
	185: '\u00b9',
	186: '\u00ba',
	187: '\u00bb',
	188: '\u00bc',
	189: '\u00bd',
	190: '\u00be',
	191: '\u00bf',
	192: '\u00c0',
	193: '\u00c1',
	194: '\u00c2',
	195: '\u00c3',
	196: '\u00c4',
	197: '\u00c5',
	198: '\u00c6',
	199: '\u00c7',
	200: '\u00c8',
	201: '\u00c9',
	202: '\u00ca',
	203: '\u00cb',
	204: '\u00cc',
	205: '\u00cd',
	206: '\u00ce',
	207: '\u00cf',
	208: '\u00d0',
	209: '\u00d1',
	210: '\u00d2',
	211: '\u00d3',
	212: '\u00d4',
	213: '\u00d5',
	214: '\u00d6',
	215: '\u00d7',
	216: '\u00d8',
	217: '\u00d9',
	218: '\u00da',
	219: '\u00db',
	220: '\u00dc',
	221: '\u00dd',
	222: '\u00de',
	223: '\u00df',
	224: '\u00e0',
	225: '\u00e1',
	226: '\u00e2',
	227: '\u00e3',
	228: '\u00e4',
	229: '\u00e5',
	230: '\u00e6',
	231: '\u00e7',
	232: '\u00e8',
	233: '\u00e9',
	234: '\u00ea',
	235: '\u00eb',
	236: '\u00ec',
	237: '\u00ed',
	238: '\u00ee',
	239: '\u00ef',
	240: '\u00f0',
	241: '\u00f1',
	242: '\u00f2',
	243: '\u00f3',
	244: '\u00f4',
	245: '\u00f5',
	246: '\u00f6',
	247: '\u00f7',
	248: '\u00f8',
	249: '\u00f9',
	250: '\u00fa',
	251: '\u00fb',
	252: '\u00fc',
	253: '\u00fd',
	254: '\u00fe',
	255: '\u00ff',
}

var utf8ToWinAnsiEncodingMap = map[rune]byte{
	'\u0020': 32,
	'\u0021': 33,
	'\u0022': 34,
	'\u0023': 35,
	'\u0024': 36,
	'\u0025': 37,
	'\u0026': 38,
	'\u0027': 39,
	'\u0028': 40,
	'\u0029': 41,
	'\u002a': 42,
	'\u002b': 43,
	'\u002c': 44,
	'\u002d': 45,
	'\u002e': 46,
	'\u002f': 47,
	'\u0030': 48,
	'\u0031': 49,
	'\u0032': 50,
	'\u0033': 51,
	'\u0034': 52,
	'\u0035': 53,
	'\u0036': 54,
	'\u0037': 55,
	'\u0038': 56,
	'\u0039': 57,
	'\u003a': 58,
	'\u003b': 59,
	'\u003c': 60,
	'\u003d': 61,
	'\u003e': 62,
	'\u003f': 63,
	'\u0040': 64,
	'\u0041': 65,
	'\u0042': 66,
	'\u0043': 67,
	'\u0044': 68,
	'\u0045': 69,
	'\u0046': 70,
	'\u0047': 71,
	'\u0048': 72,
	'\u0049': 73,
	'\u004a': 74,
	'\u004b': 75,
	'\u004c': 76,
	'\u004d': 77,
	'\u004e': 78,
	'\u004f': 79,
	'\u0050': 80,
	'\u0051': 81,
	'\u0052': 82,
	'\u0053': 83,
	'\u0054': 84,
	'\u0055': 85,
	'\u0056': 86,
	'\u0057': 87,
	'\u0058': 88,
	'\u0059': 89,
	'\u005a': 90,
	'\u005b': 91,
	'\u005c': 92,
	'\u005d': 93,
	'\u005e': 94,
	'\u005f': 95,
	'\u0060': 96,
	'\u0061': 97,
	'\u0062': 98,
	'\u0063': 99,
	'\u0064': 100,
	'\u0065': 101,
	'\u0066': 102,
	'\u0067': 103,
	'\u0068': 104,
	'\u0069': 105,
	'\u006a': 106,
	'\u006b': 107,
	'\u006c': 108,
	'\u006d': 109,
	'\u006e': 110,
	'\u006f': 111,
	'\u0070': 112,
	'\u0071': 113,
	'\u0072': 114,
	'\u0073': 115,
	'\u0074': 116,
	'\u0075': 117,
	'\u0076': 118,
	'\u0077': 119,
	'\u0078': 120,
	'\u0079': 121,
	'\u007a': 122,
	'\u007b': 123,
	'\u007c': 124,
	'\u007d': 125,
	'\u007e': 126,
	'\u2022': 127,
	'\u20ac': 128,
	// '\u2022': 129, // duplicate
	'\u201a': 130,
	'\u0192': 131,
	'\u201e': 132,
	'\u2026': 133,
	'\u2020': 134,
	'\u2021': 135,
	'\u02c6': 136,
	'\u2030': 137,
	'\u0160': 138,
	'\u2039': 139,
	'\u0152': 140,
	//'\u2022': 141, // duplicate
	'\u017d': 142,
	//'\u2022': 143, // duplicate
	// '\u2022': 144, // duplicate
	'\u2018': 145,
	'\u2019': 146,
	'\u201c': 147,
	'\u201d': 148,
	//'\u2022': 149, // duplicate
	'\u2013': 150,
	'\u2014': 151,
	'\u02dc': 152,
	'\u2122': 153,
	'\u0161': 154,
	'\u203a': 155,
	'\u0153': 156,
	//'\u2022': 157, // duplicate
	'\u017e': 158,
	'\u0178': 159,
	//'\u0020': 160, // duplicate
	'\u00a1': 161,
	'\u00a2': 162,
	'\u00a3': 163,
	'\u00a4': 164,
	'\u00a5': 165,
	'\u00a6': 166,
	'\u00a7': 167,
	'\u00a8': 168,
	'\u00a9': 169,
	'\u00aa': 170,
	'\u00ab': 171,
	'\u00ac': 172,
	//'\u002d': 173, // duplicate
	'\u00ae': 174,
	'\u00af': 175,
	'\u00b0': 176,
	'\u00b1': 177,
	'\u00b2': 178,
	'\u00b3': 179,
	'\u00b4': 180,
	'\u00b5': 181,
	'\u00b6': 182,
	'\u00b7': 183,
	'\u00b8': 184,
	'\u00b9': 185,
	'\u00ba': 186,
	'\u00bb': 187,
	'\u00bc': 188,
	'\u00bd': 189,
	'\u00be': 190,
	'\u00bf': 191,
	'\u00c0': 192,
	'\u00c1': 193,
	'\u00c2': 194,
	'\u00c3': 195,
	'\u00c4': 196,
	'\u00c5': 197,
	'\u00c6': 198,
	'\u00c7': 199,
	'\u00c8': 200,
	'\u00c9': 201,
	'\u00ca': 202,
	'\u00cb': 203,
	'\u00cc': 204,
	'\u00cd': 205,
	'\u00ce': 206,
	'\u00cf': 207,
	'\u00d0': 208,
	'\u00d1': 209,
	'\u00d2': 210,
	'\u00d3': 211,
	'\u00d4': 212,
	'\u00d5': 213,
	'\u00d6': 214,
	'\u00d7': 215,
	'\u00d8': 216,
	'\u00d9': 217,
	'\u00da': 218,
	'\u00db': 219,
	'\u00dc': 220,
	'\u00dd': 221,
	'\u00de': 222,
	'\u00df': 223,
	'\u00e0': 224,
	'\u00e1': 225,
	'\u00e2': 226,
	'\u00e3': 227,
	'\u00e4': 228,
	'\u00e5': 229,
	'\u00e6': 230,
	'\u00e7': 231,
	'\u00e8': 232,
	'\u00e9': 233,
	'\u00ea': 234,
	'\u00eb': 235,
	'\u00ec': 236,
	'\u00ed': 237,
	'\u00ee': 238,
	'\u00ef': 239,
	'\u00f0': 240,
	'\u00f1': 241,
	'\u00f2': 242,
	'\u00f3': 243,
	'\u00f4': 244,
	'\u00f5': 245,
	'\u00f6': 246,
	'\u00f7': 247,
	'\u00f8': 248,
	'\u00f9': 249,
	'\u00fa': 250,
	'\u00fb': 251,
	'\u00fc': 252,
	'\u00fd': 253,
	'\u00fe': 254,
	'\u00ff': 255,
}