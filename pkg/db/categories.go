package db

// // type Category struct {
// // 	Category string `json:"category"`
// // 	Section  string `json:"section"`
// // }

// // TODO Make this dynamic as well

// var CategoryMap = map[int]Category{
// 	0: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	1: {

// 		Category: "Seventh Umbral Era Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	2: {

// 		Category: "Seventh Astral Era Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	3: {

// 		Category: "Heavensward Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	4: {

// 		Category: "Dragonsong Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	5: {

// 		Category: "Post-Dragonsong Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	6: {

// 		Category: "Stormblood Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	7: {

// 		Category: "Post-Stormblood Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	8: {

// 		Category: "Shadowbringers Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	9: {

// 		Category: "Post-Shadowbringers Main Scenario Quests",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	10: {

// 		Category: "Post-Shadowbringers Main Scenario Quests II",
// 		Section:  "Main Scenario (A Realm Reborn through Shadowbringers)",
// 	},
// 	11: {

// 		Category: "Endwalker Main Scenario Quests",
// 		Section:  "Main Scenario (Endwalker)",
// 	},
// 	12: {

// 		Category: "Post-Endwalker Main Scenario Quests",
// 		Section:  "Main Scenario (Endwalker)",
// 	},
// 	13: {

// 		Category: "Chronicles of a New Era - Primals",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	14: {

// 		Category: "Chronicles of a New Era - Bahamut",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	15: {

// 		Category: "Chronicles of a New Era - The Crystal Tower",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	16: {

// 		Category: "Chronicles of a New Era - Alexander",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	17: {

// 		Category: "Chronicles of a New Era - The Warring Triad",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	18: {

// 		Category: "Chronicles of a New Era - The Shadow of Mhach",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	19: {

// 		Category: "Chronicles of a New Era - Omega",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	20: {

// 		Category: "Chronicles of a New Era - Omega",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	21: {

// 		Category: "Chronicles of a New Era - Return to Ivalice",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	22: {

// 		Category: "Chronicles of a New Era - The Four Lords",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	23: {

// 		Category: "Chronicles of a New Era - Eden",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	24: {

// 		Category: "YoRHa: Dark Apocalypse",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	25: {

// 		Category: "Chronicles of a New Era - The Sorrow of Werlyt",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	26: {

// 		Category: "Chronicles of a New Era - Pandæmonium",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	27: {

// 		Category: "Chronicles of a New Era - Myths of the Realm",
// 		Section:  "Chronicles of a New Era",
// 	},
// 	28: {

// 		Category: "Loporrit Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	29: {

// 		Category: "Loporrit Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	30: {

// 		Category: "Omicron Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	31: {

// 		Category: "Omicron Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	32: {

// 		Category: "Arkasodara Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	33: {

// 		Category: "Arkasodara Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	34: {

// 		Category: "Amalj'aa Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	35: {

// 		Category: "Amalj'aa Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	36: {

// 		Category: "Sylph Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	37: {

// 		Category: "Sylph Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	38: {

// 		Category: "Kobold Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	39: {

// 		Category: "Kobold Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	40: {

// 		Category: "Sahagin Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	41: {

// 		Category: "Sahagin Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	42: {

// 		Category: "Ixal Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	43: {

// 		Category: "Ixal Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	44: {

// 		Category: "Vanu Vanu Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	45: {

// 		Category: "Vanu Vanu Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	46: {

// 		Category: "Vath Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	47: {

// 		Category: "Vath Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	48: {

// 		Category: "Moogle Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	49: {

// 		Category: "Moogle Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	50: {

// 		Category: "Kojin Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	51: {

// 		Category: "Kojin Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	52: {

// 		Category: "Ananta Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	53: {

// 		Category: "Ananta Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	54: {

// 		Category: "Namazu Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	55: {

// 		Category: "Namazu Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	56: {

// 		Category: "Pixie Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	57: {

// 		Category: "Pixie Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	58: {

// 		Category: "Qitari Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	59: {

// 		Category: "Qitari Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	60: {

// 		Category: "Dwarf Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	61: {

// 		Category: "Dwarf Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	62: {

// 		Category: "Tribal Alliance Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	63: {

// 		Category: "Tribal Alliance Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	64: {

// 		Category: "Tribal Alliance Quests",
// 		Section:  "Tribal Quests (A Realm Reborn through Shadowbringers)",
// 	},
// 	65: {

// 		Category: "Tribal Alliance Quests",
// 		Section:  "Tribal Quests (Endwalker)",
// 	},
// 	66: {

// 		Category: "Chronicles of Light",
// 		Section:  "Sidequests",
// 	},
// 	67: {

// 		Category: "Chronicles of Light",
// 		Section:  "Sidequests",
// 	},
// 	68: {

// 		Category: "Chronicles of Light",
// 		Section:  "Sidequests",
// 	},
// 	69: {

// 		Category: "Hildibrand Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	70: {

// 		Category: "Hildibrand Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	71: {

// 		Category: "Hildibrand Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	72: {

// 		Category: "Hildibrand Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	73: {

// 		Category: "Hildibrand Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	74: {

// 		Category: "Weapon Enhancement Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	75: {

// 		Category: "Weapon Enhancement Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	76: {

// 		Category: "Weapon Enhancement Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	77: {

// 		Category: "Weapon Enhancement Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	78: {

// 		Category: "Side Story Quests",
// 		Section:  "Sidequests",
// 	},
// 	79: {

// 		Category: "Side Story Quests",
// 		Section:  "Sidequests",
// 	},
// 	80: {

// 		Category: "Side Story Quests",
// 		Section:  "Sidequests",
// 	},
// 	81: {

// 		Category: "Side Story Quests",
// 		Section:  "Sidequests",
// 	},
// 	82: {

// 		Category: "Side Story Quests",
// 		Section:  "Sidequests",
// 	},
// 	83: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	84: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	85: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	86: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	87: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	88: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	89: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	90: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	91: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	92: {

// 		Category: "Records of Unusual Endeavors",
// 		Section:  "Sidequests",
// 	},
// 	93: {

// 		Category: "Lominsan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	94: {

// 		Category: "Gridanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	95: {

// 		Category: "Ul'dahn Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	96: {

// 		Category: "Coerthan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	97: {

// 		Category: "Mor Dhonan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	98: {

// 		Category: "Ishgardian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	99: {

// 		Category: "Abalathian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	100: {

// 		Category: "Dravanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	101: {

// 		Category: "Dravanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	102: {

// 		Category: "Dravanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	103: {

// 		Category: "Dravanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	104: {

// 		Category: "Azys Lla Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	105: {

// 		Category: "Gyr Abanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	106: {

// 		Category: "Gyr Abanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	107: {

// 		Category: "Gyr Abanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	108: {

// 		Category: "Gyr Abanian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	109: {

// 		Category: "Hingan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	110: {

// 		Category: "Othardian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	111: {

// 		Category: "Othardian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	112: {

// 		Category: "Othardian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	113: {

// 		Category: "Lakeland Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	114: {

// 		Category: "Kholusian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	115: {

// 		Category: "Lakeland Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	116: {

// 		Category: "Kholusian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	117: {

// 		Category: "Amh Araeng Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	118: {

// 		Category: "Il Mheg Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	119: {

// 		Category: "Rak'tika Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	120: {

// 		Category: "Tempest Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	121: {

// 		Category: "Sharlayan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	122: {

// 		Category: "Thavnairian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	123: {

// 		Category: "Sharlayan Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	124: {

// 		Category: "Thavnairian Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	125: {

// 		Category: "Garlean Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	126: {

// 		Category: "Mare Lamentorum Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	127: {

// 		Category: "Elpis Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	128: {

// 		Category: "Ultima Thule Sidequests",
// 		Section:  "Sidequests",
// 	},
// 	129: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	130: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	131: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	132: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	133: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	134: {

// 		Category: "Disciple of War Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	135: {

// 		Category: "Disciple of Magic Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	136: {

// 		Category: "Disciple of Magic Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	137: {

// 		Category: "Disciple of Magic Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	138: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	139: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	140: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	141: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	142: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	143: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	144: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	145: {

// 		Category: "Disciple of the Hand Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	146: {

// 		Category: "Disciple of the Land Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	147: {

// 		Category: "Disciple of the Land Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	148: {

// 		Category: "Disciple of the Land Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	149: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	150: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	151: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	152: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	153: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	154: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	155: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	156: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	157: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	158: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	159: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	160: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	161: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	162: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	163: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	164: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	165: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	166: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	167: {

// 		Category: "Disciple of War Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	168: {

// 		Category: "Disciple of Magic Job Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	169: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	170: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	171: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	172: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	173: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	174: {

// 		Category: "Crystalline Mean Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	175: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	176: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	177: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	178: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	179: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	180: {

// 		Category: "Studium Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	181: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	182: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	183: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	184: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	185: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	186: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	187: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	188: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	189: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	190: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	191: {

// 		Category: "Role Quests",
// 		Section:  "Class & Job Quests",
// 	},
// 	192: {

// 		Category: "Grand Company Quests",
// 		Section:  "Other Quests",
// 	},
// 	193: {

// 		Category: "Grand Company Quests",
// 		Section:  "Other Quests",
// 	},
// 	194: {

// 		Category: "Grand Company Quests",
// 		Section:  "Other Quests",
// 	},
// 	195: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	196: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	197: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	198: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	199: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	200: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	201: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	202: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	203: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	204: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	205: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	206: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	207: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	208: {

// 		Category: "Seasonal Events",
// 		Section:  "Other Quests",
// 	},
// 	209: {

// 		Category: "Special Quests",
// 		Section:  "Other Quests",
// 	},
// 	210: {

// 		Category: "Special Quests",
// 		Section:  "Other Quests",
// 	},
// 	211: {

// 		Category: "Battlecraft Leves",
// 		Section:  "Levequests",
// 	},
// 	212: {

// 		Category: "Carpentry Leves",
// 		Section:  "Levequests",
// 	},
// 	213: {

// 		Category: "Blacksmithing Leves",
// 		Section:  "Levequests",
// 	},
// 	214: {

// 		Category: "Armoring Leves",
// 		Section:  "Levequests",
// 	},
// 	215: {

// 		Category: "Goldsmithing Leves",
// 		Section:  "Levequests",
// 	},
// 	216: {

// 		Category: "Leatherworking Leves",
// 		Section:  "Levequests",
// 	},
// 	217: {

// 		Category: "Clothcrafting Leves",
// 		Section:  "Levequests",
// 	},
// 	218: {

// 		Category: "Alchemy Leves",
// 		Section:  "Levequests",
// 	},
// 	219: {

// 		Category: "Cooking Leves",
// 		Section:  "Levequests",
// 	},
// 	220: {

// 		Category: "Mining Leves",
// 		Section:  "Levequests",
// 	},
// 	221: {

// 		Category: "Botany Leves",
// 		Section:  "Levequests",
// 	},
// 	222: {

// 		Category: "Fishing Leves",
// 		Section:  "Levequests",
// 	},
// 	223: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	224: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	225: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	226: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	227: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	228: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	229: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	230: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	231: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	232: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	233: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	234: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	235: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	236: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	237: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	238: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	239: {

// 		Category: "Sephiroth Missions",
// 		Section:  "Other Quests",
// 	},
// 	240: {

// 		Category: "Company Leves",
// 		Section:  "Levequests",
// 	},
// 	241: {

// 		Category: "Company Leves",
// 		Section:  "Levequests",
// 	},
// 	242: {

// 		Category: "Company Leves",
// 		Section:  "Levequests",
// 	},
// 	243: {

// 		Category: "Instanced Dungeons",
// 		Section:  "Duty",
// 	},
// 	244: {

// 		Category: "Instanced Raids",
// 		Section:  "Duty",
// 	},
// 	245: {

// 		Category: "Instanced Battles",
// 		Section:  "Duty",
// 	},
// 	246: {

// 		Category: "Other Duties",
// 		Section:  "Duty",
// 	},
// 	247: {

// 		Category: "Other Duties",
// 		Section:  "Duty",
// 	},
// 	248: {

// 		Category: "Other Duties",
// 		Section:  "Duty",
// 	},
// 	249: {

// 		Category: "Other Duties",
// 		Section:  "Duty",
// 	},
// 	250: {

// 		Category: "Instanced Dungeons",
// 		Section:  "Duty",
// 	},
// 	251: {

// 		Category: "Legacy Quests",
// 		Section:  "Other Quests",
// 	},
// 	252: {

// 		Category: "Legacy Quests",
// 		Section:  "Other Quests",
// 	},
// 	253: {

// 		Category: "Legacy Quests",
// 		Section:  "Other Quests",
// 	},
// 	254: {

// 		Category: "Legacy Quests",
// 		Section:  "Other Quests",
// 	},
// 	255: {

// 		Category: "Legacy Quests",
// 		Section:  "Other Quests",
// 	},
// }
