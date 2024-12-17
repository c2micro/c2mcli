// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MlanParser struct {
	*antlr.BaseParser
}

var MlanParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mlanParserInit() {
	staticData := &MlanParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'while'", "'{'", "'}'", "'for'", "'return'", "'continue'",
		"'break'", "'='", "'-='", "'*='", "'/='", "'%='", "'**='", "'['", "','",
		"']'", "':'", "'('", "')'", "'if'", "'elif'", "'else'", "'fn'", "'include'",
		"", "'=='", "'!='", "'||'", "'&&'", "'**'", "'>='", "'<= '", "'+='",
		"'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'^'", "'!'", "'@'",
		"", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "WS", "Eq", "Neq", "Or", "And",
		"Pow", "GtEq", "LtEq", "AssignSum", "Gt", "Lt", "Multiply", "Division",
		"Modulus", "Add", "Subtract", "Xor", "Not", "Closure", "Bool", "Null",
		"Identifier", "Integer", "IntegerHex", "Float", "String", "Comment",
	}
	staticData.RuleNames = []string{
		"program", "block", "statement", "whileStatement", "forStatement", "returnStatement",
		"continueStatement", "breakStatement", "assignment", "list", "dictUnit",
		"dict", "index", "functionInvoke", "closureInvoke", "expression", "ifBlock",
		"elifBlock", "elseBlock", "ifStatement", "functionParameters", "functionDefinition",
		"closureDefinition", "includeSubmodule",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 352, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 5, 1, 55, 8, 1, 10, 1, 12, 1, 58, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 87,
		8, 3, 10, 3, 12, 3, 90, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 102, 8, 4, 10, 4, 12, 4, 105, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 3, 5, 111, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 146, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 152, 8, 9, 10,
		9, 12, 9, 155, 9, 9, 3, 9, 157, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 169, 8, 11, 10, 11, 12, 11, 172,
		9, 11, 3, 11, 174, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 187, 8, 13, 10, 13, 12, 13, 190,
		9, 13, 3, 13, 192, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 202, 8, 14, 10, 14, 12, 14, 205, 9, 14, 3, 14, 207, 8,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 3, 15, 232, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		5, 15, 260, 8, 15, 10, 15, 12, 15, 263, 9, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 5, 16, 269, 8, 16, 10, 16, 12, 16, 272, 9, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 5, 17, 280, 8, 17, 10, 17, 12, 17, 283, 9, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 290, 8, 18, 10, 18, 12, 18, 293,
		9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 299, 8, 19, 10, 19, 12, 19, 302,
		9, 19, 1, 19, 3, 19, 305, 8, 19, 1, 20, 1, 20, 1, 20, 5, 20, 310, 8, 20,
		10, 20, 12, 20, 313, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 319, 8,
		21, 1, 21, 1, 21, 1, 21, 5, 21, 324, 8, 21, 10, 21, 12, 21, 327, 9, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 3, 22, 334, 8, 22, 1, 22, 1, 22, 1,
		22, 5, 22, 339, 8, 22, 10, 22, 12, 22, 342, 9, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 0, 1, 30, 24, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 0, 4, 1, 0, 37, 39, 1, 0, 40, 41, 2, 0, 32, 33, 35, 36, 1, 0, 27, 28,
		390, 0, 48, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 82, 1,
		0, 0, 0, 8, 93, 1, 0, 0, 0, 10, 108, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14,
		114, 1, 0, 0, 0, 16, 145, 1, 0, 0, 0, 18, 147, 1, 0, 0, 0, 20, 160, 1,
		0, 0, 0, 22, 164, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0, 26, 181, 1, 0, 0, 0,
		28, 195, 1, 0, 0, 0, 30, 231, 1, 0, 0, 0, 32, 264, 1, 0, 0, 0, 34, 275,
		1, 0, 0, 0, 36, 286, 1, 0, 0, 0, 38, 296, 1, 0, 0, 0, 40, 306, 1, 0, 0,
		0, 42, 314, 1, 0, 0, 0, 44, 330, 1, 0, 0, 0, 46, 345, 1, 0, 0, 0, 48, 49,
		3, 2, 1, 0, 49, 50, 5, 0, 0, 1, 50, 1, 1, 0, 0, 0, 51, 55, 3, 4, 2, 0,
		52, 55, 3, 42, 21, 0, 53, 55, 3, 46, 23, 0, 54, 51, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		56, 57, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 3, 16,
		8, 0, 60, 61, 5, 1, 0, 0, 61, 81, 1, 0, 0, 0, 62, 63, 3, 28, 14, 0, 63,
		64, 5, 1, 0, 0, 64, 81, 1, 0, 0, 0, 65, 66, 3, 26, 13, 0, 66, 67, 5, 1,
		0, 0, 67, 81, 1, 0, 0, 0, 68, 81, 3, 6, 3, 0, 69, 81, 3, 8, 4, 0, 70, 81,
		3, 38, 19, 0, 71, 72, 3, 14, 7, 0, 72, 73, 5, 1, 0, 0, 73, 81, 1, 0, 0,
		0, 74, 75, 3, 12, 6, 0, 75, 76, 5, 1, 0, 0, 76, 81, 1, 0, 0, 0, 77, 78,
		3, 10, 5, 0, 78, 79, 5, 1, 0, 0, 79, 81, 1, 0, 0, 0, 80, 59, 1, 0, 0, 0,
		80, 62, 1, 0, 0, 0, 80, 65, 1, 0, 0, 0, 80, 68, 1, 0, 0, 0, 80, 69, 1,
		0, 0, 0, 80, 70, 1, 0, 0, 0, 80, 71, 1, 0, 0, 0, 80, 74, 1, 0, 0, 0, 80,
		77, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82, 83, 5, 2, 0, 0, 83, 84, 3, 30, 15,
		0, 84, 88, 5, 3, 0, 0, 85, 87, 3, 4, 2, 0, 86, 85, 1, 0, 0, 0, 87, 90,
		1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0,
		90, 88, 1, 0, 0, 0, 91, 92, 5, 4, 0, 0, 92, 7, 1, 0, 0, 0, 93, 94, 5, 5,
		0, 0, 94, 95, 3, 16, 8, 0, 95, 96, 5, 1, 0, 0, 96, 97, 3, 30, 15, 0, 97,
		98, 5, 1, 0, 0, 98, 99, 3, 16, 8, 0, 99, 103, 5, 3, 0, 0, 100, 102, 3,
		4, 2, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0,
		0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106,
		107, 5, 4, 0, 0, 107, 9, 1, 0, 0, 0, 108, 110, 5, 6, 0, 0, 109, 111, 3,
		30, 15, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 11, 1, 0, 0,
		0, 112, 113, 5, 7, 0, 0, 113, 13, 1, 0, 0, 0, 114, 115, 5, 8, 0, 0, 115,
		15, 1, 0, 0, 0, 116, 117, 5, 47, 0, 0, 117, 118, 5, 9, 0, 0, 118, 146,
		3, 30, 15, 0, 119, 120, 5, 47, 0, 0, 120, 121, 5, 9, 0, 0, 121, 146, 3,
		44, 22, 0, 122, 123, 5, 47, 0, 0, 123, 124, 5, 34, 0, 0, 124, 146, 3, 30,
		15, 0, 125, 126, 5, 47, 0, 0, 126, 127, 5, 10, 0, 0, 127, 146, 3, 30, 15,
		0, 128, 129, 5, 47, 0, 0, 129, 130, 5, 11, 0, 0, 130, 146, 3, 30, 15, 0,
		131, 132, 5, 47, 0, 0, 132, 133, 5, 12, 0, 0, 133, 146, 3, 30, 15, 0, 134,
		135, 5, 47, 0, 0, 135, 136, 5, 13, 0, 0, 136, 146, 3, 30, 15, 0, 137, 138,
		5, 47, 0, 0, 138, 139, 5, 14, 0, 0, 139, 146, 3, 30, 15, 0, 140, 141, 5,
		47, 0, 0, 141, 142, 3, 24, 12, 0, 142, 143, 5, 9, 0, 0, 143, 144, 3, 30,
		15, 0, 144, 146, 1, 0, 0, 0, 145, 116, 1, 0, 0, 0, 145, 119, 1, 0, 0, 0,
		145, 122, 1, 0, 0, 0, 145, 125, 1, 0, 0, 0, 145, 128, 1, 0, 0, 0, 145,
		131, 1, 0, 0, 0, 145, 134, 1, 0, 0, 0, 145, 137, 1, 0, 0, 0, 145, 140,
		1, 0, 0, 0, 146, 17, 1, 0, 0, 0, 147, 156, 5, 15, 0, 0, 148, 153, 3, 30,
		15, 0, 149, 150, 5, 16, 0, 0, 150, 152, 3, 30, 15, 0, 151, 149, 1, 0, 0,
		0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 148, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 5, 17, 0, 0, 159, 19, 1, 0,
		0, 0, 160, 161, 3, 30, 15, 0, 161, 162, 5, 18, 0, 0, 162, 163, 3, 30, 15,
		0, 163, 21, 1, 0, 0, 0, 164, 173, 5, 3, 0, 0, 165, 170, 3, 20, 10, 0, 166,
		167, 5, 16, 0, 0, 167, 169, 3, 20, 10, 0, 168, 166, 1, 0, 0, 0, 169, 172,
		1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 174, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 173, 165, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0,
		174, 175, 1, 0, 0, 0, 175, 176, 5, 4, 0, 0, 176, 23, 1, 0, 0, 0, 177, 178,
		5, 15, 0, 0, 178, 179, 3, 30, 15, 0, 179, 180, 5, 17, 0, 0, 180, 25, 1,
		0, 0, 0, 181, 182, 5, 47, 0, 0, 182, 191, 5, 19, 0, 0, 183, 188, 3, 30,
		15, 0, 184, 185, 5, 16, 0, 0, 185, 187, 3, 30, 15, 0, 186, 184, 1, 0, 0,
		0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189,
		192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191, 183, 1, 0, 0, 0, 191, 192,
		1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 5, 20, 0, 0, 194, 27, 1, 0,
		0, 0, 195, 196, 5, 44, 0, 0, 196, 197, 5, 47, 0, 0, 197, 206, 5, 19, 0,
		0, 198, 203, 3, 30, 15, 0, 199, 200, 5, 16, 0, 0, 200, 202, 3, 30, 15,
		0, 201, 199, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 198,
		1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 5, 20,
		0, 0, 209, 29, 1, 0, 0, 0, 210, 211, 6, 15, -1, 0, 211, 232, 5, 48, 0,
		0, 212, 232, 5, 49, 0, 0, 213, 232, 5, 46, 0, 0, 214, 232, 5, 45, 0, 0,
		215, 232, 5, 47, 0, 0, 216, 232, 5, 50, 0, 0, 217, 232, 5, 51, 0, 0, 218,
		232, 3, 44, 22, 0, 219, 232, 3, 26, 13, 0, 220, 232, 3, 28, 14, 0, 221,
		222, 5, 41, 0, 0, 222, 232, 3, 30, 15, 13, 223, 224, 5, 43, 0, 0, 224,
		232, 3, 30, 15, 12, 225, 226, 5, 19, 0, 0, 226, 227, 3, 30, 15, 0, 227,
		228, 5, 20, 0, 0, 228, 232, 1, 0, 0, 0, 229, 232, 3, 18, 9, 0, 230, 232,
		3, 22, 11, 0, 231, 210, 1, 0, 0, 0, 231, 212, 1, 0, 0, 0, 231, 213, 1,
		0, 0, 0, 231, 214, 1, 0, 0, 0, 231, 215, 1, 0, 0, 0, 231, 216, 1, 0, 0,
		0, 231, 217, 1, 0, 0, 0, 231, 218, 1, 0, 0, 0, 231, 219, 1, 0, 0, 0, 231,
		220, 1, 0, 0, 0, 231, 221, 1, 0, 0, 0, 231, 223, 1, 0, 0, 0, 231, 225,
		1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 261, 1, 0,
		0, 0, 233, 234, 10, 11, 0, 0, 234, 235, 5, 31, 0, 0, 235, 260, 3, 30, 15,
		11, 236, 237, 10, 10, 0, 0, 237, 238, 7, 0, 0, 0, 238, 260, 3, 30, 15,
		11, 239, 240, 10, 9, 0, 0, 240, 241, 7, 1, 0, 0, 241, 260, 3, 30, 15, 10,
		242, 243, 10, 8, 0, 0, 243, 244, 7, 2, 0, 0, 244, 260, 3, 30, 15, 9, 245,
		246, 10, 7, 0, 0, 246, 247, 7, 3, 0, 0, 247, 260, 3, 30, 15, 8, 248, 249,
		10, 6, 0, 0, 249, 250, 5, 42, 0, 0, 250, 260, 3, 30, 15, 7, 251, 252, 10,
		5, 0, 0, 252, 253, 5, 30, 0, 0, 253, 260, 3, 30, 15, 6, 254, 255, 10, 4,
		0, 0, 255, 256, 5, 29, 0, 0, 256, 260, 3, 30, 15, 5, 257, 258, 10, 14,
		0, 0, 258, 260, 3, 24, 12, 0, 259, 233, 1, 0, 0, 0, 259, 236, 1, 0, 0,
		0, 259, 239, 1, 0, 0, 0, 259, 242, 1, 0, 0, 0, 259, 245, 1, 0, 0, 0, 259,
		248, 1, 0, 0, 0, 259, 251, 1, 0, 0, 0, 259, 254, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0,
		0, 0, 262, 31, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265, 5, 21, 0, 0,
		265, 266, 3, 30, 15, 0, 266, 270, 5, 3, 0, 0, 267, 269, 3, 4, 2, 0, 268,
		267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271,
		1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274, 5, 4,
		0, 0, 274, 33, 1, 0, 0, 0, 275, 276, 5, 22, 0, 0, 276, 277, 3, 30, 15,
		0, 277, 281, 5, 3, 0, 0, 278, 280, 3, 4, 2, 0, 279, 278, 1, 0, 0, 0, 280,
		283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 284,
		1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 285, 5, 4, 0, 0, 285, 35, 1, 0,
		0, 0, 286, 287, 5, 23, 0, 0, 287, 291, 5, 3, 0, 0, 288, 290, 3, 4, 2, 0,
		289, 288, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 295,
		5, 4, 0, 0, 295, 37, 1, 0, 0, 0, 296, 300, 3, 32, 16, 0, 297, 299, 3, 34,
		17, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		300, 301, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303,
		305, 3, 36, 18, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 39,
		1, 0, 0, 0, 306, 311, 5, 47, 0, 0, 307, 308, 5, 16, 0, 0, 308, 310, 5,
		47, 0, 0, 309, 307, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0,
		0, 311, 312, 1, 0, 0, 0, 312, 41, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314,
		315, 5, 24, 0, 0, 315, 316, 5, 47, 0, 0, 316, 318, 5, 19, 0, 0, 317, 319,
		3, 40, 20, 0, 318, 317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 1,
		0, 0, 0, 320, 321, 5, 20, 0, 0, 321, 325, 5, 3, 0, 0, 322, 324, 3, 4, 2,
		0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325,
		326, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329,
		5, 4, 0, 0, 329, 43, 1, 0, 0, 0, 330, 331, 5, 24, 0, 0, 331, 333, 5, 19,
		0, 0, 332, 334, 3, 40, 20, 0, 333, 332, 1, 0, 0, 0, 333, 334, 1, 0, 0,
		0, 334, 335, 1, 0, 0, 0, 335, 336, 5, 20, 0, 0, 336, 340, 5, 3, 0, 0, 337,
		339, 3, 4, 2, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338,
		1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342, 340, 1, 0,
		0, 0, 343, 344, 5, 4, 0, 0, 344, 45, 1, 0, 0, 0, 345, 346, 5, 25, 0, 0,
		346, 347, 5, 19, 0, 0, 347, 348, 3, 30, 15, 0, 348, 349, 5, 20, 0, 0, 349,
		350, 5, 1, 0, 0, 350, 47, 1, 0, 0, 0, 28, 54, 56, 80, 88, 103, 110, 145,
		153, 156, 170, 173, 188, 191, 203, 206, 231, 259, 261, 270, 281, 291, 300,
		304, 311, 318, 325, 333, 340,
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

// MlanParserInit initializes any static state used to implement MlanParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMlanParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MlanParserInit() {
	staticData := &MlanParserStaticData
	staticData.once.Do(mlanParserInit)
}

// NewMlanParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMlanParser(input antlr.TokenStream) *MlanParser {
	MlanParserInit()
	this := new(MlanParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MlanParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Mlan.g4"

	return this
}

// MlanParser tokens.
const (
	MlanParserEOF        = antlr.TokenEOF
	MlanParserT__0       = 1
	MlanParserT__1       = 2
	MlanParserT__2       = 3
	MlanParserT__3       = 4
	MlanParserT__4       = 5
	MlanParserT__5       = 6
	MlanParserT__6       = 7
	MlanParserT__7       = 8
	MlanParserT__8       = 9
	MlanParserT__9       = 10
	MlanParserT__10      = 11
	MlanParserT__11      = 12
	MlanParserT__12      = 13
	MlanParserT__13      = 14
	MlanParserT__14      = 15
	MlanParserT__15      = 16
	MlanParserT__16      = 17
	MlanParserT__17      = 18
	MlanParserT__18      = 19
	MlanParserT__19      = 20
	MlanParserT__20      = 21
	MlanParserT__21      = 22
	MlanParserT__22      = 23
	MlanParserT__23      = 24
	MlanParserT__24      = 25
	MlanParserWS         = 26
	MlanParserEq         = 27
	MlanParserNeq        = 28
	MlanParserOr         = 29
	MlanParserAnd        = 30
	MlanParserPow        = 31
	MlanParserGtEq       = 32
	MlanParserLtEq       = 33
	MlanParserAssignSum  = 34
	MlanParserGt         = 35
	MlanParserLt         = 36
	MlanParserMultiply   = 37
	MlanParserDivision   = 38
	MlanParserModulus    = 39
	MlanParserAdd        = 40
	MlanParserSubtract   = 41
	MlanParserXor        = 42
	MlanParserNot        = 43
	MlanParserClosure    = 44
	MlanParserBool       = 45
	MlanParserNull       = 46
	MlanParserIdentifier = 47
	MlanParserInteger    = 48
	MlanParserIntegerHex = 49
	MlanParserFloat      = 50
	MlanParserString_    = 51
	MlanParserComment    = 52
)

// MlanParser rules.
const (
	MlanParserRULE_program            = 0
	MlanParserRULE_block              = 1
	MlanParserRULE_statement          = 2
	MlanParserRULE_whileStatement     = 3
	MlanParserRULE_forStatement       = 4
	MlanParserRULE_returnStatement    = 5
	MlanParserRULE_continueStatement  = 6
	MlanParserRULE_breakStatement     = 7
	MlanParserRULE_assignment         = 8
	MlanParserRULE_list               = 9
	MlanParserRULE_dictUnit           = 10
	MlanParserRULE_dict               = 11
	MlanParserRULE_index              = 12
	MlanParserRULE_functionInvoke     = 13
	MlanParserRULE_closureInvoke      = 14
	MlanParserRULE_expression         = 15
	MlanParserRULE_ifBlock            = 16
	MlanParserRULE_elifBlock          = 17
	MlanParserRULE_elseBlock          = 18
	MlanParserRULE_ifStatement        = 19
	MlanParserRULE_functionParameters = 20
	MlanParserRULE_functionDefinition = 21
	MlanParserRULE_closureDefinition  = 22
	MlanParserRULE_includeSubmodule   = 23
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MlanParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MlanParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Block()
	}
	{
		p.SetState(49)
		p.Match(MlanParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFunctionDefinition() []IFunctionDefinitionContext
	FunctionDefinition(i int) IFunctionDefinitionContext
	AllIncludeSubmodule() []IIncludeSubmoduleContext
	IncludeSubmodule(i int) IIncludeSubmoduleContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefinitionContext); ok {
			tst[i] = t.(IFunctionDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *BlockContext) AllIncludeSubmodule() []IIncludeSubmoduleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIncludeSubmoduleContext); ok {
			len++
		}
	}

	tst := make([]IIncludeSubmoduleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIncludeSubmoduleContext); ok {
			tst[i] = t.(IIncludeSubmoduleContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) IncludeSubmodule(i int) IIncludeSubmoduleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeSubmoduleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludeSubmoduleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MlanParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329726829028) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MlanParserT__1, MlanParserT__4, MlanParserT__5, MlanParserT__6, MlanParserT__7, MlanParserT__20, MlanParserClosure, MlanParserIdentifier:
			{
				p.SetState(51)
				p.Statement()
			}

		case MlanParserT__23:
			{
				p.SetState(52)
				p.FunctionDefinition()
			}

		case MlanParserT__24:
			{
				p.SetState(53)
				p.IncludeSubmodule()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ClosureInvoke() IClosureInvokeContext
	FunctionInvoke() IFunctionInvokeContext
	WhileStatement() IWhileStatementContext
	ForStatement() IForStatementContext
	IfStatement() IIfStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	ReturnStatement() IReturnStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ClosureInvoke() IClosureInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureInvokeContext)
}

func (s *StatementContext) FunctionInvoke() IFunctionInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionInvokeContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MlanParserRULE_statement)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Assignment()
		}
		{
			p.SetState(60)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.ClosureInvoke()
		}
		{
			p.SetState(63)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.FunctionInvoke()
		}
		{
			p.SetState(66)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.WhileStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.ForStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.IfStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.BreakStatement()
		}
		{
			p.SetState(72)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(74)
			p.ContinueStatement()
		}
		{
			p.SetState(75)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(77)
			p.ReturnStatement()
		}
		{
			p.SetState(78)
			p.Match(MlanParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *WhileStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MlanParserRULE_whileStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(MlanParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.expression(0)
	}
	{
		p.SetState(84)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(85)
			p.Statement()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MlanParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(MlanParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Assignment()
	}
	{
		p.SetState(95)
		p.Match(MlanParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.expression(0)
	}
	{
		p.SetState(97)
		p.Match(MlanParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Assignment()
	}
	{
		p.SetState(99)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MlanParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(MlanParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4497002574938120) != 0 {
		{
			p.SetState(109)
			p.expression(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MlanParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(MlanParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MlanParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(MlanParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyAll(ctx *AssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentRegularContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentRegularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentRegularContext {
	var p = new(AssignmentRegularContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentRegularContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentRegularContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentRegularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentRegularContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentRegularContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentRegularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentRegular(s)
	}
}

func (s *AssignmentRegularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentRegular(s)
	}
}

func (s *AssignmentRegularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentRegular(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentSumContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentSumContext {
	var p = new(AssignmentSumContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentSumContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentSumContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentSumContext) AssignSum() antlr.TerminalNode {
	return s.GetToken(MlanParserAssignSum, 0)
}

func (s *AssignmentSumContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentSumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentSum(s)
	}
}

func (s *AssignmentSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentSum(s)
	}
}

func (s *AssignmentSumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentSum(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentIndexRegularContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentIndexRegularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentIndexRegularContext {
	var p = new(AssignmentIndexRegularContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentIndexRegularContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentIndexRegularContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentIndexRegularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentIndexRegularContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *AssignmentIndexRegularContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentIndexRegularContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentIndexRegularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentIndexRegular(s)
	}
}

func (s *AssignmentIndexRegularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentIndexRegular(s)
	}
}

func (s *AssignmentIndexRegularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentIndexRegular(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentPowContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentPowContext {
	var p = new(AssignmentPowContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentPowContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentPowContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentPowContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentPowContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentPow(s)
	}
}

func (s *AssignmentPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentPow(s)
	}
}

func (s *AssignmentPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentDivContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentDivContext {
	var p = new(AssignmentDivContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentDivContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentDivContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentDivContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentDivContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentDiv(s)
	}
}

func (s *AssignmentDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentDiv(s)
	}
}

func (s *AssignmentDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentModContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentModContext {
	var p = new(AssignmentModContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentModContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentModContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentModContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentModContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentMod(s)
	}
}

func (s *AssignmentModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentMod(s)
	}
}

func (s *AssignmentModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentMulContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentMulContext {
	var p = new(AssignmentMulContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentMulContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentMulContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentMulContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentMulContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentMul(s)
	}
}

func (s *AssignmentMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentMul(s)
	}
}

func (s *AssignmentMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentSubContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentSubContext {
	var p = new(AssignmentSubContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentSubContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentSubContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentSubContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentSubContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentSub(s)
	}
}

func (s *AssignmentSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentSub(s)
	}
}

func (s *AssignmentSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentClosureContext struct {
	AssignmentContext
	varScalarName antlr.Token
}

func NewAssignmentClosureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentClosureContext {
	var p = new(AssignmentClosureContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentClosureContext) GetVarScalarName() antlr.Token { return s.varScalarName }

func (s *AssignmentClosureContext) SetVarScalarName(v antlr.Token) { s.varScalarName = v }

func (s *AssignmentClosureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentClosureContext) ClosureDefinition() IClosureDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureDefinitionContext)
}

func (s *AssignmentClosureContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *AssignmentClosureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterAssignmentClosure(s)
	}
}

func (s *AssignmentClosureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitAssignmentClosure(s)
	}
}

func (s *AssignmentClosureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitAssignmentClosure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MlanParserRULE_assignment)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentRegularContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(MlanParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expression(0)
		}

	case 2:
		localctx = NewAssignmentClosureContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentClosureContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(MlanParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.ClosureDefinition()
		}

	case 3:
		localctx = NewAssignmentSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentSumContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(MlanParserAssignSum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.expression(0)
		}

	case 4:
		localctx = NewAssignmentSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentSubContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(MlanParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.expression(0)
		}

	case 5:
		localctx = NewAssignmentMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentMulContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(MlanParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expression(0)
		}

	case 6:
		localctx = NewAssignmentDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentDivContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(MlanParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expression(0)
		}

	case 7:
		localctx = NewAssignmentModContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentModContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(MlanParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expression(0)
		}

	case 8:
		localctx = NewAssignmentPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(137)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentPowContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(MlanParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.expression(0)
		}

	case 9:
		localctx = NewAssignmentIndexRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(140)

			var _m = p.Match(MlanParserIdentifier)

			localctx.(*AssignmentIndexRegularContext).varScalarName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Index()
		}
		{
			p.SetState(142)
			p.Match(MlanParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MlanParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(MlanParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4497002574938120) != 0 {
		{
			p.SetState(148)
			p.expression(0)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MlanParserT__15 {
			{
				p.SetState(149)
				p.Match(MlanParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(150)
				p.expression(0)
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(158)
		p.Match(MlanParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictUnitContext is an interface to support dynamic dispatch.
type IDictUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsDictUnitContext differentiates from other interfaces.
	IsDictUnitContext()
}

type DictUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictUnitContext() *DictUnitContext {
	var p = new(DictUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_dictUnit
	return p
}

func InitEmptyDictUnitContext(p *DictUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_dictUnit
}

func (*DictUnitContext) IsDictUnitContext() {}

func NewDictUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictUnitContext {
	var p = new(DictUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_dictUnit

	return p
}

func (s *DictUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *DictUnitContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DictUnitContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DictUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterDictUnit(s)
	}
}

func (s *DictUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitDictUnit(s)
	}
}

func (s *DictUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitDictUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) DictUnit() (localctx IDictUnitContext) {
	localctx = NewDictUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MlanParserRULE_dictUnit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.expression(0)
	}
	{
		p.SetState(161)
		p.Match(MlanParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictContext is an interface to support dynamic dispatch.
type IDictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDictUnit() []IDictUnitContext
	DictUnit(i int) IDictUnitContext

	// IsDictContext differentiates from other interfaces.
	IsDictContext()
}

type DictContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictContext() *DictContext {
	var p = new(DictContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_dict
	return p
}

func InitEmptyDictContext(p *DictContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_dict
}

func (*DictContext) IsDictContext() {}

func NewDictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictContext {
	var p = new(DictContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_dict

	return p
}

func (s *DictContext) GetParser() antlr.Parser { return s.parser }

func (s *DictContext) AllDictUnit() []IDictUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictUnitContext); ok {
			len++
		}
	}

	tst := make([]IDictUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictUnitContext); ok {
			tst[i] = t.(IDictUnitContext)
			i++
		}
	}

	return tst
}

func (s *DictContext) DictUnit(i int) IDictUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictUnitContext)
}

func (s *DictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterDict(s)
	}
}

func (s *DictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitDict(s)
	}
}

func (s *DictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitDict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Dict() (localctx IDictContext) {
	localctx = NewDictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MlanParserRULE_dict)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4497002574938120) != 0 {
		{
			p.SetState(165)
			p.DictUnit()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MlanParserT__15 {
			{
				p.SetState(166)
				p.Match(MlanParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(167)
				p.DictUnit()
			}

			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(175)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MlanParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(MlanParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.expression(0)
	}
	{
		p.SetState(179)
		p.Match(MlanParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionInvokeContext is an interface to support dynamic dispatch.
type IFunctionInvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionInvokeContext differentiates from other interfaces.
	IsFunctionInvokeContext()
}

type FunctionInvokeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionInvokeContext() *FunctionInvokeContext {
	var p = new(FunctionInvokeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionInvoke
	return p
}

func InitEmptyFunctionInvokeContext(p *FunctionInvokeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionInvoke
}

func (*FunctionInvokeContext) IsFunctionInvokeContext() {}

func NewFunctionInvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionInvokeContext {
	var p = new(FunctionInvokeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_functionInvoke

	return p
}

func (s *FunctionInvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionInvokeContext) CopyAll(ctx *FunctionInvokeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierFunctionInvokeContext struct {
	FunctionInvokeContext
	varFunctionName antlr.Token
}

func NewIdentifierFunctionInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierFunctionInvokeContext {
	var p = new(IdentifierFunctionInvokeContext)

	InitEmptyFunctionInvokeContext(&p.FunctionInvokeContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionInvokeContext))

	return p
}

func (s *IdentifierFunctionInvokeContext) GetVarFunctionName() antlr.Token { return s.varFunctionName }

func (s *IdentifierFunctionInvokeContext) SetVarFunctionName(v antlr.Token) { s.varFunctionName = v }

func (s *IdentifierFunctionInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFunctionInvokeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *IdentifierFunctionInvokeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierFunctionInvokeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdentifierFunctionInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIdentifierFunctionInvoke(s)
	}
}

func (s *IdentifierFunctionInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIdentifierFunctionInvoke(s)
	}
}

func (s *IdentifierFunctionInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIdentifierFunctionInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) FunctionInvoke() (localctx IFunctionInvokeContext) {
	localctx = NewFunctionInvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MlanParserRULE_functionInvoke)
	var _la int

	localctx = NewIdentifierFunctionInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)

		var _m = p.Match(MlanParserIdentifier)

		localctx.(*IdentifierFunctionInvokeContext).varFunctionName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(MlanParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4497002574938120) != 0 {
		{
			p.SetState(183)
			p.expression(0)
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MlanParserT__15 {
			{
				p.SetState(184)
				p.Match(MlanParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(185)
				p.expression(0)
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(193)
		p.Match(MlanParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosureInvokeContext is an interface to support dynamic dispatch.
type IClosureInvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsClosureInvokeContext differentiates from other interfaces.
	IsClosureInvokeContext()
}

type ClosureInvokeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosureInvokeContext() *ClosureInvokeContext {
	var p = new(ClosureInvokeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_closureInvoke
	return p
}

func InitEmptyClosureInvokeContext(p *ClosureInvokeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_closureInvoke
}

func (*ClosureInvokeContext) IsClosureInvokeContext() {}

func NewClosureInvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosureInvokeContext {
	var p = new(ClosureInvokeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_closureInvoke

	return p
}

func (s *ClosureInvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosureInvokeContext) CopyAll(ctx *ClosureInvokeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ClosureInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureInvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierClosureInvokeContext struct {
	ClosureInvokeContext
	varClosureName antlr.Token
}

func NewIdentifierClosureInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierClosureInvokeContext {
	var p = new(IdentifierClosureInvokeContext)

	InitEmptyClosureInvokeContext(&p.ClosureInvokeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ClosureInvokeContext))

	return p
}

func (s *IdentifierClosureInvokeContext) GetVarClosureName() antlr.Token { return s.varClosureName }

func (s *IdentifierClosureInvokeContext) SetVarClosureName(v antlr.Token) { s.varClosureName = v }

func (s *IdentifierClosureInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierClosureInvokeContext) Closure() antlr.TerminalNode {
	return s.GetToken(MlanParserClosure, 0)
}

func (s *IdentifierClosureInvokeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *IdentifierClosureInvokeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierClosureInvokeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdentifierClosureInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIdentifierClosureInvoke(s)
	}
}

func (s *IdentifierClosureInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIdentifierClosureInvoke(s)
	}
}

func (s *IdentifierClosureInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIdentifierClosureInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ClosureInvoke() (localctx IClosureInvokeContext) {
	localctx = NewClosureInvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MlanParserRULE_closureInvoke)
	var _la int

	localctx = NewIdentifierClosureInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(MlanParserClosure)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)

		var _m = p.Match(MlanParserIdentifier)

		localctx.(*IdentifierClosureInvokeContext).varClosureName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(MlanParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4497002574938120) != 0 {
		{
			p.SetState(198)
			p.expression(0)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MlanParserT__15 {
			{
				p.SetState(199)
				p.Match(MlanParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.expression(0)
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(208)
		p.Match(MlanParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionIntegerHexContext struct {
	ExpressionContext
}

func NewExpressionIntegerHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIntegerHexContext {
	var p = new(ExpressionIntegerHexContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionIntegerHexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIntegerHexContext) IntegerHex() antlr.TerminalNode {
	return s.GetToken(MlanParserIntegerHex, 0)
}

func (s *ExpressionIntegerHexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionIntegerHex(s)
	}
}

func (s *ExpressionIntegerHexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionIntegerHex(s)
	}
}

func (s *ExpressionIntegerHexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionIntegerHex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionFunctionInvokeContext struct {
	ExpressionContext
}

func NewExpressionFunctionInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionFunctionInvokeContext {
	var p = new(ExpressionFunctionInvokeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionFunctionInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionFunctionInvokeContext) FunctionInvoke() IFunctionInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionInvokeContext)
}

func (s *ExpressionFunctionInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionFunctionInvoke(s)
	}
}

func (s *ExpressionFunctionInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionFunctionInvoke(s)
	}
}

func (s *ExpressionFunctionInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionFunctionInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionUnaryNegationContext struct {
	ExpressionContext
}

func NewExpressionUnaryNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionUnaryNegationContext {
	var p = new(ExpressionUnaryNegationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionUnaryNegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionUnaryNegationContext) Subtract() antlr.TerminalNode {
	return s.GetToken(MlanParserSubtract, 0)
}

func (s *ExpressionUnaryNegationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionUnaryNegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionUnaryNegation(s)
	}
}

func (s *ExpressionUnaryNegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionUnaryNegation(s)
	}
}

func (s *ExpressionUnaryNegationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionUnaryNegation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBoolContext struct {
	ExpressionContext
}

func NewExpressionBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBoolContext {
	var p = new(ExpressionBoolContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBoolContext) Bool() antlr.TerminalNode {
	return s.GetToken(MlanParserBool, 0)
}

func (s *ExpressionBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionBool(s)
	}
}

func (s *ExpressionBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionBool(s)
	}
}

func (s *ExpressionBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPowContext struct {
	ExpressionContext
}

func NewExpressionPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPowContext {
	var p = new(ExpressionPowContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPowContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionPowContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionPowContext) Pow() antlr.TerminalNode {
	return s.GetToken(MlanParserPow, 0)
}

func (s *ExpressionPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionPow(s)
	}
}

func (s *ExpressionPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionPow(s)
	}
}

func (s *ExpressionPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionXorContext struct {
	ExpressionContext
}

func NewExpressionXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionXorContext {
	var p = new(ExpressionXorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionXorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionXorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionXorContext) Xor() antlr.TerminalNode {
	return s.GetToken(MlanParserXor, 0)
}

func (s *ExpressionXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionXor(s)
	}
}

func (s *ExpressionXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionXor(s)
	}
}

func (s *ExpressionXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionEqualContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpressionEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionEqualContext {
	var p = new(ExpressionEqualContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionEqualContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionEqualContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionEqualContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionEqualContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionEqualContext) Eq() antlr.TerminalNode {
	return s.GetToken(MlanParserEq, 0)
}

func (s *ExpressionEqualContext) Neq() antlr.TerminalNode {
	return s.GetToken(MlanParserNeq, 0)
}

func (s *ExpressionEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionEqual(s)
	}
}

func (s *ExpressionEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionEqual(s)
	}
}

func (s *ExpressionEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionClosureContext struct {
	ExpressionContext
}

func NewExpressionClosureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionClosureContext {
	var p = new(ExpressionClosureContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionClosureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionClosureContext) ClosureDefinition() IClosureDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureDefinitionContext)
}

func (s *ExpressionClosureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionClosure(s)
	}
}

func (s *ExpressionClosureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionClosure(s)
	}
}

func (s *ExpressionClosureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionClosure(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionDictContext struct {
	ExpressionContext
}

func NewExpressionDictContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionDictContext {
	var p = new(ExpressionDictContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionDictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDictContext) Dict() IDictContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *ExpressionDictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionDict(s)
	}
}

func (s *ExpressionDictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionDict(s)
	}
}

func (s *ExpressionDictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionDict(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionIdentifierContext struct {
	ExpressionContext
}

func NewExpressionIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIdentifierContext {
	var p = new(ExpressionIdentifierContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *ExpressionIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionIdentifier(s)
	}
}

func (s *ExpressionIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionIdentifier(s)
	}
}

func (s *ExpressionIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionListContext struct {
	ExpressionContext
}

func NewExpressionListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListContext {
	var p = new(ExpressionListContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionSumSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpressionSumSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSumSubContext {
	var p = new(ExpressionSumSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionSumSubContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionSumSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionSumSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSumSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionSumSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionSumSubContext) Add() antlr.TerminalNode {
	return s.GetToken(MlanParserAdd, 0)
}

func (s *ExpressionSumSubContext) Subtract() antlr.TerminalNode {
	return s.GetToken(MlanParserSubtract, 0)
}

func (s *ExpressionSumSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionSumSub(s)
	}
}

func (s *ExpressionSumSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionSumSub(s)
	}
}

func (s *ExpressionSumSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionSumSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionComparisonContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpressionComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionComparisonContext {
	var p = new(ExpressionComparisonContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionComparisonContext) GtEq() antlr.TerminalNode {
	return s.GetToken(MlanParserGtEq, 0)
}

func (s *ExpressionComparisonContext) LtEq() antlr.TerminalNode {
	return s.GetToken(MlanParserLtEq, 0)
}

func (s *ExpressionComparisonContext) Gt() antlr.TerminalNode {
	return s.GetToken(MlanParserGt, 0)
}

func (s *ExpressionComparisonContext) Lt() antlr.TerminalNode {
	return s.GetToken(MlanParserLt, 0)
}

func (s *ExpressionComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionComparison(s)
	}
}

func (s *ExpressionComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionComparison(s)
	}
}

func (s *ExpressionComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionLogicalOrContext struct {
	ExpressionContext
}

func NewExpressionLogicalOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionLogicalOrContext {
	var p = new(ExpressionLogicalOrContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionLogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionLogicalOrContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionLogicalOrContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionLogicalOrContext) Or() antlr.TerminalNode {
	return s.GetToken(MlanParserOr, 0)
}

func (s *ExpressionLogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionLogicalOr(s)
	}
}

func (s *ExpressionLogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionLogicalOr(s)
	}
}

func (s *ExpressionLogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionIndexContext struct {
	ExpressionContext
}

func NewExpressionIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIndexContext {
	var p = new(ExpressionIndexContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionIndexContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *ExpressionIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionIndex(s)
	}
}

func (s *ExpressionIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionIndex(s)
	}
}

func (s *ExpressionIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionLogicalNotContext struct {
	ExpressionContext
}

func NewExpressionLogicalNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionLogicalNotContext {
	var p = new(ExpressionLogicalNotContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionLogicalNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionLogicalNotContext) Not() antlr.TerminalNode {
	return s.GetToken(MlanParserNot, 0)
}

func (s *ExpressionLogicalNotContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionLogicalNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionLogicalNot(s)
	}
}

func (s *ExpressionLogicalNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionLogicalNot(s)
	}
}

func (s *ExpressionLogicalNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionLogicalNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionClosureInvokeContext struct {
	ExpressionContext
}

func NewExpressionClosureInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionClosureInvokeContext {
	var p = new(ExpressionClosureInvokeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionClosureInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionClosureInvokeContext) ClosureInvoke() IClosureInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureInvokeContext)
}

func (s *ExpressionClosureInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionClosureInvoke(s)
	}
}

func (s *ExpressionClosureInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionClosureInvoke(s)
	}
}

func (s *ExpressionClosureInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionClosureInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionParenthesesContext struct {
	ExpressionContext
}

func NewExpressionParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionParenthesesContext {
	var p = new(ExpressionParenthesesContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionParenthesesContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionParentheses(s)
	}
}

func (s *ExpressionParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionParentheses(s)
	}
}

func (s *ExpressionParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionMulDivModContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpressionMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionMulDivModContext {
	var p = new(ExpressionMulDivModContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionMulDivModContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionMulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMulDivModContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionMulDivModContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionMulDivModContext) Multiply() antlr.TerminalNode {
	return s.GetToken(MlanParserMultiply, 0)
}

func (s *ExpressionMulDivModContext) Division() antlr.TerminalNode {
	return s.GetToken(MlanParserDivision, 0)
}

func (s *ExpressionMulDivModContext) Modulus() antlr.TerminalNode {
	return s.GetToken(MlanParserModulus, 0)
}

func (s *ExpressionMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionMulDivMod(s)
	}
}

func (s *ExpressionMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionMulDivMod(s)
	}
}

func (s *ExpressionMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionLogicalAndContext struct {
	ExpressionContext
}

func NewExpressionLogicalAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionLogicalAndContext {
	var p = new(ExpressionLogicalAndContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionLogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionLogicalAndContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionLogicalAndContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionLogicalAndContext) And() antlr.TerminalNode {
	return s.GetToken(MlanParserAnd, 0)
}

func (s *ExpressionLogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionLogicalAnd(s)
	}
}

func (s *ExpressionLogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionLogicalAnd(s)
	}
}

func (s *ExpressionLogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionFloatContext struct {
	ExpressionContext
}

func NewExpressionFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionFloatContext {
	var p = new(ExpressionFloatContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionFloatContext) Float() antlr.TerminalNode {
	return s.GetToken(MlanParserFloat, 0)
}

func (s *ExpressionFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionFloat(s)
	}
}

func (s *ExpressionFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionFloat(s)
	}
}

func (s *ExpressionFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionIntegerContext struct {
	ExpressionContext
}

func NewExpressionIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIntegerContext {
	var p = new(ExpressionIntegerContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIntegerContext) Integer() antlr.TerminalNode {
	return s.GetToken(MlanParserInteger, 0)
}

func (s *ExpressionIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionInteger(s)
	}
}

func (s *ExpressionIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionInteger(s)
	}
}

func (s *ExpressionIntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionNullContext struct {
	ExpressionContext
}

func NewExpressionNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionNullContext {
	var p = new(ExpressionNullContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionNullContext) Null() antlr.TerminalNode {
	return s.GetToken(MlanParserNull, 0)
}

func (s *ExpressionNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionNull(s)
	}
}

func (s *ExpressionNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionNull(s)
	}
}

func (s *ExpressionNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionStringContext struct {
	ExpressionContext
}

func NewExpressionStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStringContext {
	var p = new(ExpressionStringContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStringContext) String_() antlr.TerminalNode {
	return s.GetToken(MlanParserString_, 0)
}

func (s *ExpressionStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterExpressionString(s)
	}
}

func (s *ExpressionStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitExpressionString(s)
	}
}

func (s *ExpressionStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitExpressionString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MlanParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, MlanParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(211)
			p.Match(MlanParserInteger)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpressionIntegerHexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(212)
			p.Match(MlanParserIntegerHex)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExpressionNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(213)
			p.Match(MlanParserNull)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpressionBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(214)
			p.Match(MlanParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpressionIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(215)
			p.Match(MlanParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpressionFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(216)
			p.Match(MlanParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpressionStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(217)
			p.Match(MlanParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpressionClosureContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(218)
			p.ClosureDefinition()
		}

	case 9:
		localctx = NewExpressionFunctionInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(219)
			p.FunctionInvoke()
		}

	case 10:
		localctx = NewExpressionClosureInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)
			p.ClosureInvoke()
		}

	case 11:
		localctx = NewExpressionUnaryNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(221)
			p.Match(MlanParserSubtract)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.expression(13)
		}

	case 12:
		localctx = NewExpressionLogicalNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Match(MlanParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.expression(12)
		}

	case 13:
		localctx = NewExpressionParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(225)
			p.Match(MlanParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expression(0)
		}
		{
			p.SetState(227)
			p.Match(MlanParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewExpressionListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(229)
			p.List()
		}

	case 15:
		localctx = NewExpressionDictContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Dict()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(234)
					p.Match(MlanParserPow)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(235)
					p.expression(11)
				}

			case 2:
				localctx = NewExpressionMulDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(237)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionMulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&962072674304) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionMulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(238)
					p.expression(11)
				}

			case 3:
				localctx = NewExpressionSumSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(240)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionSumSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MlanParserAdd || _la == MlanParserSubtract) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionSumSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(241)
					p.expression(10)
				}

			case 4:
				localctx = NewExpressionComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(243)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&115964116992) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(244)
					p.expression(9)
				}

			case 5:
				localctx = NewExpressionEqualContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(246)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionEqualContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MlanParserEq || _la == MlanParserNeq) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionEqualContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(247)
					p.expression(8)
				}

			case 6:
				localctx = NewExpressionXorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(249)
					p.Match(MlanParserXor)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(250)
					p.expression(7)
				}

			case 7:
				localctx = NewExpressionLogicalAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(252)
					p.Match(MlanParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(253)
					p.expression(6)
				}

			case 8:
				localctx = NewExpressionLogicalOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(255)
					p.Match(MlanParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.expression(5)
				}

			case 9:
				localctx = NewExpressionIndexContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MlanParserRULE_expression)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(258)
					p.Index()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) CopyAll(ctx *IfBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfBlockStatementContext struct {
	IfBlockContext
}

func NewIfBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfBlockStatementContext {
	var p = new(IfBlockStatementContext)

	InitEmptyIfBlockContext(&p.IfBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfBlockContext))

	return p
}

func (s *IfBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfBlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIfBlockStatement(s)
	}
}

func (s *IfBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIfBlockStatement(s)
	}
}

func (s *IfBlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIfBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MlanParserRULE_ifBlock)
	var _la int

	localctx = NewIfBlockStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(MlanParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.expression(0)
	}
	{
		p.SetState(266)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(267)
			p.Statement()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElifBlockContext is an interface to support dynamic dispatch.
type IElifBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElifBlockContext differentiates from other interfaces.
	IsElifBlockContext()
}

type ElifBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifBlockContext() *ElifBlockContext {
	var p = new(ElifBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_elifBlock
	return p
}

func InitEmptyElifBlockContext(p *ElifBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_elifBlock
}

func (*ElifBlockContext) IsElifBlockContext() {}

func NewElifBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifBlockContext {
	var p = new(ElifBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_elifBlock

	return p
}

func (s *ElifBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifBlockContext) CopyAll(ctx *ElifBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElifBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElifBlockStatementContext struct {
	ElifBlockContext
}

func NewElifBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElifBlockStatementContext {
	var p = new(ElifBlockStatementContext)

	InitEmptyElifBlockContext(&p.ElifBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElifBlockContext))

	return p
}

func (s *ElifBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifBlockStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElifBlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ElifBlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElifBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterElifBlockStatement(s)
	}
}

func (s *ElifBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitElifBlockStatement(s)
	}
}

func (s *ElifBlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitElifBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ElifBlock() (localctx IElifBlockContext) {
	localctx = NewElifBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MlanParserRULE_elifBlock)
	var _la int

	localctx = NewElifBlockStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(MlanParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.expression(0)
	}
	{
		p.SetState(277)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(278)
			p.Statement()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(284)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) CopyAll(ctx *ElseBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseBlockStatementContext struct {
	ElseBlockContext
}

func NewElseBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseBlockStatementContext {
	var p = new(ElseBlockStatementContext)

	InitEmptyElseBlockContext(&p.ElseBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElseBlockContext))

	return p
}

func (s *ElseBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ElseBlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterElseBlockStatement(s)
	}
}

func (s *ElseBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitElseBlockStatement(s)
	}
}

func (s *ElseBlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitElseBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MlanParserRULE_elseBlock)
	var _la int

	localctx = NewElseBlockStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(MlanParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(288)
			p.Statement()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(294)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfBlock() IIfBlockContext
	AllElifBlock() []IElifBlockContext
	ElifBlock(i int) IElifBlockContext
	ElseBlock() IElseBlockContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfStatementContext) AllElifBlock() []IElifBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElifBlockContext); ok {
			len++
		}
	}

	tst := make([]IElifBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElifBlockContext); ok {
			tst[i] = t.(IElifBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElifBlock(i int) IElifBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifBlockContext)
}

func (s *IfStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MlanParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.IfBlock()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MlanParserT__21 {
		{
			p.SetState(297)
			p.ElifBlock()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MlanParserT__22 {
		{
			p.SetState(303)
			p.ElseBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionParametersContext is an interface to support dynamic dispatch.
type IFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsFunctionParametersContext differentiates from other interfaces.
	IsFunctionParametersContext()
}

type FunctionParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParametersContext() *FunctionParametersContext {
	var p = new(FunctionParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionParameters
	return p
}

func InitEmptyFunctionParametersContext(p *FunctionParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionParameters
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_functionParameters

	return p
}

func (s *FunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParametersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MlanParserIdentifier)
}

func (s *FunctionParametersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, i)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitFunctionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MlanParserRULE_functionParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(MlanParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MlanParserT__15 {
		{
			p.SetState(307)
			p.Match(MlanParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(MlanParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarFunctionName returns the varFunctionName token.
	GetVarFunctionName() antlr.Token

	// SetVarFunctionName sets the varFunctionName token.
	SetVarFunctionName(antlr.Token)

	// Getter signatures
	Identifier() antlr.TerminalNode
	FunctionParameters() IFunctionParametersContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	varFunctionName antlr.Token
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) GetVarFunctionName() antlr.Token { return s.varFunctionName }

func (s *FunctionDefinitionContext) SetVarFunctionName(v antlr.Token) { s.varFunctionName = v }

func (s *FunctionDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MlanParserIdentifier, 0)
}

func (s *FunctionDefinitionContext) FunctionParameters() IFunctionParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionDefinitionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDefinitionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MlanParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(MlanParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)

		var _m = p.Match(MlanParserIdentifier)

		localctx.(*FunctionDefinitionContext).varFunctionName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(MlanParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MlanParserIdentifier {
		{
			p.SetState(317)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(320)
		p.Match(MlanParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(322)
			p.Statement()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(328)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosureDefinitionContext is an interface to support dynamic dispatch.
type IClosureDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionParameters() IFunctionParametersContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsClosureDefinitionContext differentiates from other interfaces.
	IsClosureDefinitionContext()
}

type ClosureDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosureDefinitionContext() *ClosureDefinitionContext {
	var p = new(ClosureDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_closureDefinition
	return p
}

func InitEmptyClosureDefinitionContext(p *ClosureDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_closureDefinition
}

func (*ClosureDefinitionContext) IsClosureDefinitionContext() {}

func NewClosureDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosureDefinitionContext {
	var p = new(ClosureDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_closureDefinition

	return p
}

func (s *ClosureDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosureDefinitionContext) FunctionParameters() IFunctionParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *ClosureDefinitionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ClosureDefinitionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ClosureDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosureDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterClosureDefinition(s)
	}
}

func (s *ClosureDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitClosureDefinition(s)
	}
}

func (s *ClosureDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitClosureDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) ClosureDefinition() (localctx IClosureDefinitionContext) {
	localctx = NewClosureDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MlanParserRULE_closureDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(MlanParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(MlanParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MlanParserIdentifier {
		{
			p.SetState(332)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(335)
		p.Match(MlanParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(MlanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329676497380) != 0 {
		{
			p.SetState(337)
			p.Statement()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(343)
		p.Match(MlanParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncludeSubmoduleContext is an interface to support dynamic dispatch.
type IIncludeSubmoduleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIncludeSubmoduleContext differentiates from other interfaces.
	IsIncludeSubmoduleContext()
}

type IncludeSubmoduleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeSubmoduleContext() *IncludeSubmoduleContext {
	var p = new(IncludeSubmoduleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_includeSubmodule
	return p
}

func InitEmptyIncludeSubmoduleContext(p *IncludeSubmoduleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MlanParserRULE_includeSubmodule
}

func (*IncludeSubmoduleContext) IsIncludeSubmoduleContext() {}

func NewIncludeSubmoduleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeSubmoduleContext {
	var p = new(IncludeSubmoduleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MlanParserRULE_includeSubmodule

	return p
}

func (s *IncludeSubmoduleContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeSubmoduleContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IncludeSubmoduleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeSubmoduleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeSubmoduleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.EnterIncludeSubmodule(s)
	}
}

func (s *IncludeSubmoduleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MlanListener); ok {
		listenerT.ExitIncludeSubmodule(s)
	}
}

func (s *IncludeSubmoduleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MlanVisitor:
		return t.VisitIncludeSubmodule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MlanParser) IncludeSubmodule() (localctx IIncludeSubmoduleContext) {
	localctx = NewIncludeSubmoduleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MlanParserRULE_includeSubmodule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(MlanParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(MlanParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.expression(0)
	}
	{
		p.SetState(348)
		p.Match(MlanParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(MlanParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MlanParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MlanParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
