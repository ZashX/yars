// Code generated from antlr/Yars.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type YarsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var yarslexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yarslexerLexerInit() {
	staticData := &yarslexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'int32'", "'uint32'", "'float32'", "'bool'", "'string'", "'not'",
		"'and'", "'or'", "'in'", "'+'", "'-'", "'*'", "'unit'", "'action'",
		"'project'", "'recommend'", "", "'{'", "'}'", "'['", "']'", "'('", "')'",
		"';'", "':'", "'@'", "','", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "INT32", "UINT32", "FLOAT32", "BOOL", "STRING", "NOT", "AND", "OR",
		"IN", "PLUS", "MINUS", "MUL", "UNIT", "ACTION", "PROJECT", "RECOMMEND",
		"IDENTIFIER", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "L_RBRACKET",
		"R_RBRACKET", "SEMI", "COLON", "AT", "COMMA", "DOT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"INT32", "UINT32", "FLOAT32", "BOOL", "STRING", "NOT", "AND", "OR",
		"IN", "PLUS", "MINUS", "MUL", "UNIT", "ACTION", "PROJECT", "RECOMMEND",
		"IDENTIFIER", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "L_RBRACKET",
		"R_RBRACKET", "SEMI", "COLON", "AT", "COMMA", "DOT", "WHITESPACE", "LETTER",
		"UNICODE_DIGIT", "UNICODE_LETTER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 193, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 152, 8,
		16, 10, 16, 12, 16, 155, 9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 4, 28, 180, 8, 28, 11, 28,
		12, 28, 181, 1, 28, 1, 28, 1, 29, 1, 29, 3, 29, 188, 8, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 0, 61, 0, 63, 0, 1, 0, 3, 3, 0, 9, 10, 13,
		13, 32, 32, 61, 0, 48, 57, 1632, 1641, 1776, 1785, 1984, 1993, 2406, 2415,
		2534, 2543, 2662, 2671, 2790, 2799, 2918, 2927, 3046, 3055, 3174, 3183,
		3302, 3311, 3430, 3439, 3558, 3567, 3664, 3673, 3792, 3801, 3872, 3881,
		4160, 4169, 4240, 4249, 6112, 6121, 6160, 6169, 6470, 6479, 6608, 6617,
		6784, 6793, 6800, 6809, 6992, 7001, 7088, 7097, 7232, 7241, 7248, 7257,
		42528, 42537, 43216, 43225, 43264, 43273, 43472, 43481, 43504, 43513, 43600,
		43609, 44016, 44025, 65296, 65305, 66720, 66729, 68912, 68921, 69734, 69743,
		69872, 69881, 69942, 69951, 70096, 70105, 70384, 70393, 70736, 70745, 70864,
		70873, 71248, 71257, 71360, 71369, 71472, 71481, 71904, 71913, 72016, 72025,
		72784, 72793, 73040, 73049, 73120, 73129, 92768, 92777, 93008, 93017, 120782,
		120831, 123200, 123209, 123632, 123641, 125264, 125273, 130032, 130041,
		622, 0, 65, 90, 97, 122, 170, 170, 181, 181, 186, 186, 192, 214, 216, 246,
		248, 705, 710, 721, 736, 740, 748, 748, 750, 750, 880, 884, 886, 887, 890,
		893, 895, 895, 902, 902, 904, 906, 908, 908, 910, 929, 931, 1013, 1015,
		1153, 1162, 1327, 1329, 1366, 1369, 1369, 1376, 1416, 1488, 1514, 1519,
		1522, 1568, 1610, 1646, 1647, 1649, 1747, 1749, 1749, 1765, 1766, 1774,
		1775, 1786, 1788, 1791, 1791, 1808, 1808, 1810, 1839, 1869, 1957, 1969,
		1969, 1994, 2026, 2036, 2037, 2042, 2042, 2048, 2069, 2074, 2074, 2084,
		2084, 2088, 2088, 2112, 2136, 2144, 2154, 2208, 2228, 2230, 2247, 2308,
		2361, 2365, 2365, 2384, 2384, 2392, 2401, 2417, 2432, 2437, 2444, 2447,
		2448, 2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489, 2493, 2493, 2510,
		2510, 2524, 2525, 2527, 2529, 2544, 2545, 2556, 2556, 2565, 2570, 2575,
		2576, 2579, 2600, 2602, 2608, 2610, 2611, 2613, 2614, 2616, 2617, 2649,
		2652, 2654, 2654, 2674, 2676, 2693, 2701, 2703, 2705, 2707, 2728, 2730,
		2736, 2738, 2739, 2741, 2745, 2749, 2749, 2768, 2768, 2784, 2785, 2809,
		2809, 2821, 2828, 2831, 2832, 2835, 2856, 2858, 2864, 2866, 2867, 2869,
		2873, 2877, 2877, 2908, 2909, 2911, 2913, 2929, 2929, 2947, 2947, 2949,
		2954, 2958, 2960, 2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975, 2979,
		2980, 2984, 2986, 2990, 3001, 3024, 3024, 3077, 3084, 3086, 3088, 3090,
		3112, 3114, 3129, 3133, 3133, 3160, 3162, 3168, 3169, 3200, 3200, 3205,
		3212, 3214, 3216, 3218, 3240, 3242, 3251, 3253, 3257, 3261, 3261, 3294,
		3294, 3296, 3297, 3313, 3314, 3332, 3340, 3342, 3344, 3346, 3386, 3389,
		3389, 3406, 3406, 3412, 3414, 3423, 3425, 3450, 3455, 3461, 3478, 3482,
		3505, 3507, 3515, 3517, 3517, 3520, 3526, 3585, 3632, 3634, 3635, 3648,
		3654, 3713, 3714, 3716, 3716, 3718, 3722, 3724, 3747, 3749, 3749, 3751,
		3760, 3762, 3763, 3773, 3773, 3776, 3780, 3782, 3782, 3804, 3807, 3840,
		3840, 3904, 3911, 3913, 3948, 3976, 3980, 4096, 4138, 4159, 4159, 4176,
		4181, 4186, 4189, 4193, 4193, 4197, 4198, 4206, 4208, 4213, 4225, 4238,
		4238, 4256, 4293, 4295, 4295, 4301, 4301, 4304, 4346, 4348, 4680, 4682,
		4685, 4688, 4694, 4696, 4696, 4698, 4701, 4704, 4744, 4746, 4749, 4752,
		4784, 4786, 4789, 4792, 4798, 4800, 4800, 4802, 4805, 4808, 4822, 4824,
		4880, 4882, 4885, 4888, 4954, 4992, 5007, 5024, 5109, 5112, 5117, 5121,
		5740, 5743, 5759, 5761, 5786, 5792, 5866, 5873, 5880, 5888, 5900, 5902,
		5905, 5920, 5937, 5952, 5969, 5984, 5996, 5998, 6000, 6016, 6067, 6103,
		6103, 6108, 6108, 6176, 6264, 6272, 6276, 6279, 6312, 6314, 6314, 6320,
		6389, 6400, 6430, 6480, 6509, 6512, 6516, 6528, 6571, 6576, 6601, 6656,
		6678, 6688, 6740, 6823, 6823, 6917, 6963, 6981, 6987, 7043, 7072, 7086,
		7087, 7098, 7141, 7168, 7203, 7245, 7247, 7258, 7293, 7296, 7304, 7312,
		7354, 7357, 7359, 7401, 7404, 7406, 7411, 7413, 7414, 7418, 7418, 7424,
		7615, 7680, 7957, 7960, 7965, 7968, 8005, 8008, 8013, 8016, 8023, 8025,
		8025, 8027, 8027, 8029, 8029, 8031, 8061, 8064, 8116, 8118, 8124, 8126,
		8126, 8130, 8132, 8134, 8140, 8144, 8147, 8150, 8155, 8160, 8172, 8178,
		8180, 8182, 8188, 8305, 8305, 8319, 8319, 8336, 8348, 8450, 8450, 8455,
		8455, 8458, 8467, 8469, 8469, 8473, 8477, 8484, 8484, 8486, 8486, 8488,
		8488, 8490, 8493, 8495, 8505, 8508, 8511, 8517, 8521, 8526, 8526, 8579,
		8580, 11264, 11310, 11312, 11358, 11360, 11492, 11499, 11502, 11506, 11507,
		11520, 11557, 11559, 11559, 11565, 11565, 11568, 11623, 11631, 11631, 11648,
		11670, 11680, 11686, 11688, 11694, 11696, 11702, 11704, 11710, 11712, 11718,
		11720, 11726, 11728, 11734, 11736, 11742, 11823, 11823, 12293, 12294, 12337,
		12341, 12347, 12348, 12353, 12438, 12445, 12447, 12449, 12538, 12540, 12543,
		12549, 12591, 12593, 12686, 12704, 12735, 12784, 12799, 13312, 19903, 19968,
		40956, 40960, 42124, 42192, 42237, 42240, 42508, 42512, 42527, 42538, 42539,
		42560, 42606, 42623, 42653, 42656, 42725, 42775, 42783, 42786, 42888, 42891,
		42943, 42946, 42954, 42997, 43009, 43011, 43013, 43015, 43018, 43020, 43042,
		43072, 43123, 43138, 43187, 43250, 43255, 43259, 43259, 43261, 43262, 43274,
		43301, 43312, 43334, 43360, 43388, 43396, 43442, 43471, 43471, 43488, 43492,
		43494, 43503, 43514, 43518, 43520, 43560, 43584, 43586, 43588, 43595, 43616,
		43638, 43642, 43642, 43646, 43695, 43697, 43697, 43701, 43702, 43705, 43709,
		43712, 43712, 43714, 43714, 43739, 43741, 43744, 43754, 43762, 43764, 43777,
		43782, 43785, 43790, 43793, 43798, 43808, 43814, 43816, 43822, 43824, 43866,
		43868, 43881, 43888, 44002, 44032, 55203, 55216, 55238, 55243, 55291, 63744,
		64109, 64112, 64217, 64256, 64262, 64275, 64279, 64285, 64285, 64287, 64296,
		64298, 64310, 64312, 64316, 64318, 64318, 64320, 64321, 64323, 64324, 64326,
		64433, 64467, 64829, 64848, 64911, 64914, 64967, 65008, 65019, 65136, 65140,
		65142, 65276, 65313, 65338, 65345, 65370, 65382, 65470, 65474, 65479, 65482,
		65487, 65490, 65495, 65498, 65500, 65536, 65547, 65549, 65574, 65576, 65594,
		65596, 65597, 65599, 65613, 65616, 65629, 65664, 65786, 66176, 66204, 66208,
		66256, 66304, 66335, 66349, 66368, 66370, 66377, 66384, 66421, 66432, 66461,
		66464, 66499, 66504, 66511, 66560, 66717, 66736, 66771, 66776, 66811, 66816,
		66855, 66864, 66915, 67072, 67382, 67392, 67413, 67424, 67431, 67584, 67589,
		67592, 67592, 67594, 67637, 67639, 67640, 67644, 67644, 67647, 67669, 67680,
		67702, 67712, 67742, 67808, 67826, 67828, 67829, 67840, 67861, 67872, 67897,
		67968, 68023, 68030, 68031, 68096, 68096, 68112, 68115, 68117, 68119, 68121,
		68149, 68192, 68220, 68224, 68252, 68288, 68295, 68297, 68324, 68352, 68405,
		68416, 68437, 68448, 68466, 68480, 68497, 68608, 68680, 68736, 68786, 68800,
		68850, 68864, 68899, 69248, 69289, 69296, 69297, 69376, 69404, 69415, 69415,
		69424, 69445, 69552, 69572, 69600, 69622, 69635, 69687, 69763, 69807, 69840,
		69864, 69891, 69926, 69956, 69956, 69959, 69959, 69968, 70002, 70006, 70006,
		70019, 70066, 70081, 70084, 70106, 70106, 70108, 70108, 70144, 70161, 70163,
		70187, 70272, 70278, 70280, 70280, 70282, 70285, 70287, 70301, 70303, 70312,
		70320, 70366, 70405, 70412, 70415, 70416, 70419, 70440, 70442, 70448, 70450,
		70451, 70453, 70457, 70461, 70461, 70480, 70480, 70493, 70497, 70656, 70708,
		70727, 70730, 70751, 70753, 70784, 70831, 70852, 70853, 70855, 70855, 71040,
		71086, 71128, 71131, 71168, 71215, 71236, 71236, 71296, 71338, 71352, 71352,
		71424, 71450, 71680, 71723, 71840, 71903, 71935, 71942, 71945, 71945, 71948,
		71955, 71957, 71958, 71960, 71983, 71999, 71999, 72001, 72001, 72096, 72103,
		72106, 72144, 72161, 72161, 72163, 72163, 72192, 72192, 72203, 72242, 72250,
		72250, 72272, 72272, 72284, 72329, 72349, 72349, 72384, 72440, 72704, 72712,
		72714, 72750, 72768, 72768, 72818, 72847, 72960, 72966, 72968, 72969, 72971,
		73008, 73030, 73030, 73056, 73061, 73063, 73064, 73066, 73097, 73112, 73112,
		73440, 73458, 73648, 73648, 73728, 74649, 74880, 75075, 77824, 78894, 82944,
		83526, 92160, 92728, 92736, 92766, 92880, 92909, 92928, 92975, 92992, 92995,
		93027, 93047, 93053, 93071, 93760, 93823, 93952, 94026, 94032, 94032, 94099,
		94111, 94176, 94177, 94179, 94179, 94208, 100343, 100352, 101589, 101632,
		101640, 110592, 110878, 110928, 110930, 110948, 110951, 110960, 111355,
		113664, 113770, 113776, 113788, 113792, 113800, 113808, 113817, 119808,
		119892, 119894, 119964, 119966, 119967, 119970, 119970, 119973, 119974,
		119977, 119980, 119982, 119993, 119995, 119995, 119997, 120003, 120005,
		120069, 120071, 120074, 120077, 120084, 120086, 120092, 120094, 120121,
		120123, 120126, 120128, 120132, 120134, 120134, 120138, 120144, 120146,
		120485, 120488, 120512, 120514, 120538, 120540, 120570, 120572, 120596,
		120598, 120628, 120630, 120654, 120656, 120686, 120688, 120712, 120714,
		120744, 120746, 120770, 120772, 120779, 123136, 123180, 123191, 123197,
		123214, 123214, 123584, 123627, 124928, 125124, 125184, 125251, 125259,
		125259, 126464, 126467, 126469, 126495, 126497, 126498, 126500, 126500,
		126503, 126503, 126505, 126514, 126516, 126519, 126521, 126521, 126523,
		126523, 126530, 126530, 126535, 126535, 126537, 126537, 126539, 126539,
		126541, 126543, 126545, 126546, 126548, 126548, 126551, 126551, 126553,
		126553, 126555, 126555, 126557, 126557, 126559, 126559, 126561, 126562,
		126564, 126564, 126567, 126570, 126572, 126578, 126580, 126583, 126585,
		126588, 126590, 126590, 126592, 126601, 126603, 126619, 126625, 126627,
		126629, 126633, 126635, 126651, 131072, 173789, 173824, 177972, 177984,
		178205, 178208, 183969, 183984, 191456, 194560, 195101, 196608, 201546,
		193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 71, 1,
		0, 0, 0, 5, 78, 1, 0, 0, 0, 7, 86, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 98,
		1, 0, 0, 0, 13, 102, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 109, 1, 0, 0,
		0, 19, 112, 1, 0, 0, 0, 21, 114, 1, 0, 0, 0, 23, 116, 1, 0, 0, 0, 25, 118,
		1, 0, 0, 0, 27, 123, 1, 0, 0, 0, 29, 130, 1, 0, 0, 0, 31, 138, 1, 0, 0,
		0, 33, 148, 1, 0, 0, 0, 35, 156, 1, 0, 0, 0, 37, 158, 1, 0, 0, 0, 39, 160,
		1, 0, 0, 0, 41, 162, 1, 0, 0, 0, 43, 164, 1, 0, 0, 0, 45, 166, 1, 0, 0,
		0, 47, 168, 1, 0, 0, 0, 49, 170, 1, 0, 0, 0, 51, 172, 1, 0, 0, 0, 53, 174,
		1, 0, 0, 0, 55, 176, 1, 0, 0, 0, 57, 179, 1, 0, 0, 0, 59, 187, 1, 0, 0,
		0, 61, 189, 1, 0, 0, 0, 63, 191, 1, 0, 0, 0, 65, 66, 5, 105, 0, 0, 66,
		67, 5, 110, 0, 0, 67, 68, 5, 116, 0, 0, 68, 69, 5, 51, 0, 0, 69, 70, 5,
		50, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 117, 0, 0, 72, 73, 5, 105, 0, 0,
		73, 74, 5, 110, 0, 0, 74, 75, 5, 116, 0, 0, 75, 76, 5, 51, 0, 0, 76, 77,
		5, 50, 0, 0, 77, 4, 1, 0, 0, 0, 78, 79, 5, 102, 0, 0, 79, 80, 5, 108, 0,
		0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 116, 0, 0, 83,
		84, 5, 51, 0, 0, 84, 85, 5, 50, 0, 0, 85, 6, 1, 0, 0, 0, 86, 87, 5, 98,
		0, 0, 87, 88, 5, 111, 0, 0, 88, 89, 5, 111, 0, 0, 89, 90, 5, 108, 0, 0,
		90, 8, 1, 0, 0, 0, 91, 92, 5, 115, 0, 0, 92, 93, 5, 116, 0, 0, 93, 94,
		5, 114, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 97, 5, 103,
		0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 5, 110, 0, 0, 99, 100, 5, 111, 0, 0,
		100, 101, 5, 116, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103, 5, 97, 0, 0, 103,
		104, 5, 110, 0, 0, 104, 105, 5, 100, 0, 0, 105, 14, 1, 0, 0, 0, 106, 107,
		5, 111, 0, 0, 107, 108, 5, 114, 0, 0, 108, 16, 1, 0, 0, 0, 109, 110, 5,
		105, 0, 0, 110, 111, 5, 110, 0, 0, 111, 18, 1, 0, 0, 0, 112, 113, 5, 43,
		0, 0, 113, 20, 1, 0, 0, 0, 114, 115, 5, 45, 0, 0, 115, 22, 1, 0, 0, 0,
		116, 117, 5, 42, 0, 0, 117, 24, 1, 0, 0, 0, 118, 119, 5, 117, 0, 0, 119,
		120, 5, 110, 0, 0, 120, 121, 5, 105, 0, 0, 121, 122, 5, 116, 0, 0, 122,
		26, 1, 0, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5, 99, 0, 0, 125, 126,
		5, 116, 0, 0, 126, 127, 5, 105, 0, 0, 127, 128, 5, 111, 0, 0, 128, 129,
		5, 110, 0, 0, 129, 28, 1, 0, 0, 0, 130, 131, 5, 112, 0, 0, 131, 132, 5,
		114, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 106, 0, 0, 134, 135, 5,
		101, 0, 0, 135, 136, 5, 99, 0, 0, 136, 137, 5, 116, 0, 0, 137, 30, 1, 0,
		0, 0, 138, 139, 5, 114, 0, 0, 139, 140, 5, 101, 0, 0, 140, 141, 5, 99,
		0, 0, 141, 142, 5, 111, 0, 0, 142, 143, 5, 109, 0, 0, 143, 144, 5, 109,
		0, 0, 144, 145, 5, 101, 0, 0, 145, 146, 5, 110, 0, 0, 146, 147, 5, 100,
		0, 0, 147, 32, 1, 0, 0, 0, 148, 153, 3, 59, 29, 0, 149, 152, 3, 59, 29,
		0, 150, 152, 3, 61, 30, 0, 151, 149, 1, 0, 0, 0, 151, 150, 1, 0, 0, 0,
		152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		34, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5, 123, 0, 0, 157, 36,
		1, 0, 0, 0, 158, 159, 5, 125, 0, 0, 159, 38, 1, 0, 0, 0, 160, 161, 5, 91,
		0, 0, 161, 40, 1, 0, 0, 0, 162, 163, 5, 93, 0, 0, 163, 42, 1, 0, 0, 0,
		164, 165, 5, 40, 0, 0, 165, 44, 1, 0, 0, 0, 166, 167, 5, 41, 0, 0, 167,
		46, 1, 0, 0, 0, 168, 169, 5, 59, 0, 0, 169, 48, 1, 0, 0, 0, 170, 171, 5,
		58, 0, 0, 171, 50, 1, 0, 0, 0, 172, 173, 5, 64, 0, 0, 173, 52, 1, 0, 0,
		0, 174, 175, 5, 44, 0, 0, 175, 54, 1, 0, 0, 0, 176, 177, 5, 46, 0, 0, 177,
		56, 1, 0, 0, 0, 178, 180, 7, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 181, 1,
		0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0,
		0, 183, 184, 6, 28, 0, 0, 184, 58, 1, 0, 0, 0, 185, 188, 3, 63, 31, 0,
		186, 188, 5, 95, 0, 0, 187, 185, 1, 0, 0, 0, 187, 186, 1, 0, 0, 0, 188,
		60, 1, 0, 0, 0, 189, 190, 7, 1, 0, 0, 190, 62, 1, 0, 0, 0, 191, 192, 7,
		2, 0, 0, 192, 64, 1, 0, 0, 0, 5, 0, 151, 153, 181, 187, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// YarsLexerInit initializes any static state used to implement YarsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewYarsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func YarsLexerInit() {
	staticData := &yarslexerLexerStaticData
	staticData.once.Do(yarslexerLexerInit)
}

// NewYarsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewYarsLexer(input antlr.CharStream) *YarsLexer {
	YarsLexerInit()
	l := new(YarsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &yarslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Yars.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// YarsLexer tokens.
const (
	YarsLexerINT32      = 1
	YarsLexerUINT32     = 2
	YarsLexerFLOAT32    = 3
	YarsLexerBOOL       = 4
	YarsLexerSTRING     = 5
	YarsLexerNOT        = 6
	YarsLexerAND        = 7
	YarsLexerOR         = 8
	YarsLexerIN         = 9
	YarsLexerPLUS       = 10
	YarsLexerMINUS      = 11
	YarsLexerMUL        = 12
	YarsLexerUNIT       = 13
	YarsLexerACTION     = 14
	YarsLexerPROJECT    = 15
	YarsLexerRECOMMEND  = 16
	YarsLexerIDENTIFIER = 17
	YarsLexerL_CURLY    = 18
	YarsLexerR_CURLY    = 19
	YarsLexerL_BRACKET  = 20
	YarsLexerR_BRACKET  = 21
	YarsLexerL_RBRACKET = 22
	YarsLexerR_RBRACKET = 23
	YarsLexerSEMI       = 24
	YarsLexerCOLON      = 25
	YarsLexerAT         = 26
	YarsLexerCOMMA      = 27
	YarsLexerDOT        = 28
	YarsLexerWHITESPACE = 29
)