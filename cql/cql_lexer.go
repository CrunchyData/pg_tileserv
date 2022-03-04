// Generated from CqlLexer.g4 by ANTLR 4.7.

package cql

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 71, 657,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9,
	17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22,
	4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4,
	28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33,
	9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9,
	38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43,
	4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4,
	49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54,
	9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9,
	59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64,
	4, 65, 9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4,
	70, 9, 70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75,
	9, 75, 4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9,
	80, 4, 81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85,
	4, 86, 9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4,
	91, 9, 91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96,
	9, 96, 4, 97, 9, 97, 4, 98, 9, 98, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 257,
	10, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 285, 10, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 5, 45, 333, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 5, 46, 403, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	56, 3, 56, 5, 56, 508, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58,
	3, 58, 7, 58, 517, 10, 58, 12, 58, 14, 58, 520, 11, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 5, 58, 526, 10, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60,
	5, 60, 534, 10, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3,
	64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69,
	3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3,
	75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80,
	3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 3, 83, 3, 84, 3, 84, 3, 85, 3,
	85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 5, 85, 591, 10, 85, 3, 86, 3, 86,
	5, 86, 595, 10, 86, 3, 87, 5, 87, 598, 10, 87, 3, 87, 3, 87, 5, 87, 602,
	10, 87, 3, 88, 3, 88, 3, 88, 5, 88, 607, 10, 88, 5, 88, 609, 10, 88, 3,
	88, 3, 88, 3, 88, 5, 88, 614, 10, 88, 3, 89, 3, 89, 3, 89, 3, 89, 3, 90,
	3, 90, 3, 91, 3, 91, 3, 92, 5, 92, 625, 10, 92, 3, 92, 3, 92, 3, 93, 6,
	93, 630, 10, 93, 13, 93, 14, 93, 631, 3, 94, 3, 94, 5, 94, 636, 10, 94,
	3, 95, 6, 95, 639, 10, 95, 13, 95, 14, 95, 640, 3, 95, 3, 95, 3, 96, 3,
	96, 3, 96, 3, 96, 3, 97, 3, 97, 3, 97, 3, 97, 3, 97, 3, 98, 3, 98, 3, 98,
	3, 98, 2, 2, 99, 4, 2, 6, 2, 8, 2, 10, 2, 12, 2, 14, 2, 16, 2, 18, 2, 20,
	2, 22, 2, 24, 2, 26, 2, 28, 2, 30, 2, 32, 2, 34, 2, 36, 2, 38, 2, 40, 2,
	42, 2, 44, 2, 46, 2, 48, 2, 50, 2, 52, 2, 54, 2, 56, 3, 58, 4, 60, 5, 62,
	6, 64, 7, 66, 8, 68, 9, 70, 10, 72, 11, 74, 12, 76, 13, 78, 14, 80, 15,
	82, 16, 84, 17, 86, 18, 88, 19, 90, 20, 92, 21, 94, 22, 96, 23, 98, 24,
	100, 25, 102, 26, 104, 27, 106, 28, 108, 29, 110, 30, 112, 31, 114, 2,
	116, 32, 118, 33, 120, 34, 122, 35, 124, 36, 126, 37, 128, 38, 130, 39,
	132, 40, 134, 41, 136, 42, 138, 43, 140, 44, 142, 45, 144, 46, 146, 47,
	148, 48, 150, 49, 152, 50, 154, 51, 156, 52, 158, 53, 160, 54, 162, 55,
	164, 56, 166, 57, 168, 58, 170, 59, 172, 60, 174, 61, 176, 62, 178, 63,
	180, 64, 182, 65, 184, 66, 186, 67, 188, 68, 190, 69, 192, 70, 194, 71,
	196, 2, 4, 2, 3, 32, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2,
	69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2,
	72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2,
	75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2,
	78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2,
	81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2,
	84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2,
	87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2,
	90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 4, 2,
	67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 41,
	41, 2, 668, 2, 56, 3, 2, 2, 2, 2, 58, 3, 2, 2, 2, 2, 60, 3, 2, 2, 2, 2,
	62, 3, 2, 2, 2, 2, 64, 3, 2, 2, 2, 2, 66, 3, 2, 2, 2, 2, 68, 3, 2, 2, 2,
	2, 70, 3, 2, 2, 2, 2, 72, 3, 2, 2, 2, 2, 74, 3, 2, 2, 2, 2, 76, 3, 2, 2,
	2, 2, 78, 3, 2, 2, 2, 2, 80, 3, 2, 2, 2, 2, 82, 3, 2, 2, 2, 2, 84, 3, 2,
	2, 2, 2, 86, 3, 2, 2, 2, 2, 88, 3, 2, 2, 2, 2, 90, 3, 2, 2, 2, 2, 92, 3,
	2, 2, 2, 2, 94, 3, 2, 2, 2, 2, 96, 3, 2, 2, 2, 2, 98, 3, 2, 2, 2, 2, 100,
	3, 2, 2, 2, 2, 102, 3, 2, 2, 2, 2, 104, 3, 2, 2, 2, 2, 106, 3, 2, 2, 2,
	2, 108, 3, 2, 2, 2, 2, 110, 3, 2, 2, 2, 2, 112, 3, 2, 2, 2, 2, 114, 3,
	2, 2, 2, 2, 116, 3, 2, 2, 2, 2, 118, 3, 2, 2, 2, 2, 120, 3, 2, 2, 2, 2,
	122, 3, 2, 2, 2, 2, 124, 3, 2, 2, 2, 2, 126, 3, 2, 2, 2, 2, 128, 3, 2,
	2, 2, 2, 130, 3, 2, 2, 2, 2, 132, 3, 2, 2, 2, 2, 134, 3, 2, 2, 2, 2, 136,
	3, 2, 2, 2, 2, 138, 3, 2, 2, 2, 2, 140, 3, 2, 2, 2, 2, 142, 3, 2, 2, 2,
	2, 144, 3, 2, 2, 2, 2, 146, 3, 2, 2, 2, 2, 148, 3, 2, 2, 2, 2, 150, 3,
	2, 2, 2, 2, 152, 3, 2, 2, 2, 2, 154, 3, 2, 2, 2, 2, 156, 3, 2, 2, 2, 2,
	158, 3, 2, 2, 2, 2, 160, 3, 2, 2, 2, 2, 162, 3, 2, 2, 2, 2, 164, 3, 2,
	2, 2, 2, 166, 3, 2, 2, 2, 2, 168, 3, 2, 2, 2, 2, 170, 3, 2, 2, 2, 2, 172,
	3, 2, 2, 2, 2, 174, 3, 2, 2, 2, 2, 176, 3, 2, 2, 2, 2, 178, 3, 2, 2, 2,
	2, 180, 3, 2, 2, 2, 2, 182, 3, 2, 2, 2, 2, 184, 3, 2, 2, 2, 2, 186, 3,
	2, 2, 2, 2, 188, 3, 2, 2, 2, 2, 190, 3, 2, 2, 2, 3, 192, 3, 2, 2, 2, 3,
	194, 3, 2, 2, 2, 3, 196, 3, 2, 2, 2, 4, 198, 3, 2, 2, 2, 6, 200, 3, 2,
	2, 2, 8, 202, 3, 2, 2, 2, 10, 204, 3, 2, 2, 2, 12, 206, 3, 2, 2, 2, 14,
	208, 3, 2, 2, 2, 16, 210, 3, 2, 2, 2, 18, 212, 3, 2, 2, 2, 20, 214, 3,
	2, 2, 2, 22, 216, 3, 2, 2, 2, 24, 218, 3, 2, 2, 2, 26, 220, 3, 2, 2, 2,
	28, 222, 3, 2, 2, 2, 30, 224, 3, 2, 2, 2, 32, 226, 3, 2, 2, 2, 34, 228,
	3, 2, 2, 2, 36, 230, 3, 2, 2, 2, 38, 232, 3, 2, 2, 2, 40, 234, 3, 2, 2,
	2, 42, 236, 3, 2, 2, 2, 44, 238, 3, 2, 2, 2, 46, 240, 3, 2, 2, 2, 48, 242,
	3, 2, 2, 2, 50, 244, 3, 2, 2, 2, 52, 246, 3, 2, 2, 2, 54, 248, 3, 2, 2,
	2, 56, 256, 3, 2, 2, 2, 58, 258, 3, 2, 2, 2, 60, 260, 3, 2, 2, 2, 62, 262,
	3, 2, 2, 2, 64, 264, 3, 2, 2, 2, 66, 267, 3, 2, 2, 2, 68, 270, 3, 2, 2,
	2, 70, 284, 3, 2, 2, 2, 72, 286, 3, 2, 2, 2, 74, 290, 3, 2, 2, 2, 76, 293,
	3, 2, 2, 2, 78, 297, 3, 2, 2, 2, 80, 302, 3, 2, 2, 2, 82, 308, 3, 2, 2,
	2, 84, 316, 3, 2, 2, 2, 86, 319, 3, 2, 2, 2, 88, 324, 3, 2, 2, 2, 90, 332,
	3, 2, 2, 2, 92, 402, 3, 2, 2, 2, 94, 404, 3, 2, 2, 2, 96, 412, 3, 2, 2,
	2, 98, 418, 3, 2, 2, 2, 100, 429, 3, 2, 2, 2, 102, 437, 3, 2, 2, 2, 104,
	448, 3, 2, 2, 2, 106, 464, 3, 2, 2, 2, 108, 477, 3, 2, 2, 2, 110, 496,
	3, 2, 2, 2, 112, 507, 3, 2, 2, 2, 114, 509, 3, 2, 2, 2, 116, 525, 3, 2,
	2, 2, 118, 527, 3, 2, 2, 2, 120, 533, 3, 2, 2, 2, 122, 535, 3, 2, 2, 2,
	124, 537, 3, 2, 2, 2, 126, 539, 3, 2, 2, 2, 128, 541, 3, 2, 2, 2, 130,
	543, 3, 2, 2, 2, 132, 545, 3, 2, 2, 2, 134, 547, 3, 2, 2, 2, 136, 549,
	3, 2, 2, 2, 138, 551, 3, 2, 2, 2, 140, 553, 3, 2, 2, 2, 142, 555, 3, 2,
	2, 2, 144, 557, 3, 2, 2, 2, 146, 559, 3, 2, 2, 2, 148, 561, 3, 2, 2, 2,
	150, 563, 3, 2, 2, 2, 152, 565, 3, 2, 2, 2, 154, 567, 3, 2, 2, 2, 156,
	569, 3, 2, 2, 2, 158, 571, 3, 2, 2, 2, 160, 573, 3, 2, 2, 2, 162, 575,
	3, 2, 2, 2, 164, 577, 3, 2, 2, 2, 166, 579, 3, 2, 2, 2, 168, 581, 3, 2,
	2, 2, 170, 590, 3, 2, 2, 2, 172, 594, 3, 2, 2, 2, 174, 601, 3, 2, 2, 2,
	176, 613, 3, 2, 2, 2, 178, 615, 3, 2, 2, 2, 180, 619, 3, 2, 2, 2, 182,
	621, 3, 2, 2, 2, 184, 624, 3, 2, 2, 2, 186, 629, 3, 2, 2, 2, 188, 635,
	3, 2, 2, 2, 190, 638, 3, 2, 2, 2, 192, 644, 3, 2, 2, 2, 194, 648, 3, 2,
	2, 2, 196, 653, 3, 2, 2, 2, 198, 199, 9, 2, 2, 2, 199, 5, 3, 2, 2, 2, 200,
	201, 9, 3, 2, 2, 201, 7, 3, 2, 2, 2, 202, 203, 9, 4, 2, 2, 203, 9, 3, 2,
	2, 2, 204, 205, 9, 5, 2, 2, 205, 11, 3, 2, 2, 2, 206, 207, 9, 6, 2, 2,
	207, 13, 3, 2, 2, 2, 208, 209, 9, 7, 2, 2, 209, 15, 3, 2, 2, 2, 210, 211,
	9, 8, 2, 2, 211, 17, 3, 2, 2, 2, 212, 213, 9, 9, 2, 2, 213, 19, 3, 2, 2,
	2, 214, 215, 9, 10, 2, 2, 215, 21, 3, 2, 2, 2, 216, 217, 9, 11, 2, 2, 217,
	23, 3, 2, 2, 2, 218, 219, 9, 12, 2, 2, 219, 25, 3, 2, 2, 2, 220, 221, 9,
	13, 2, 2, 221, 27, 3, 2, 2, 2, 222, 223, 9, 14, 2, 2, 223, 29, 3, 2, 2,
	2, 224, 225, 9, 15, 2, 2, 225, 31, 3, 2, 2, 2, 226, 227, 9, 16, 2, 2, 227,
	33, 3, 2, 2, 2, 228, 229, 9, 17, 2, 2, 229, 35, 3, 2, 2, 2, 230, 231, 9,
	18, 2, 2, 231, 37, 3, 2, 2, 2, 232, 233, 9, 19, 2, 2, 233, 39, 3, 2, 2,
	2, 234, 235, 9, 20, 2, 2, 235, 41, 3, 2, 2, 2, 236, 237, 9, 21, 2, 2, 237,
	43, 3, 2, 2, 2, 238, 239, 9, 22, 2, 2, 239, 45, 3, 2, 2, 2, 240, 241, 9,
	23, 2, 2, 241, 47, 3, 2, 2, 2, 242, 243, 9, 24, 2, 2, 243, 49, 3, 2, 2,
	2, 244, 245, 9, 25, 2, 2, 245, 51, 3, 2, 2, 2, 246, 247, 9, 26, 2, 2, 247,
	53, 3, 2, 2, 2, 248, 249, 9, 27, 2, 2, 249, 55, 3, 2, 2, 2, 250, 257, 5,
	60, 30, 2, 251, 257, 5, 64, 32, 2, 252, 257, 5, 58, 29, 2, 253, 257, 5,
	62, 31, 2, 254, 257, 5, 68, 34, 2, 255, 257, 5, 66, 33, 2, 256, 250, 3,
	2, 2, 2, 256, 251, 3, 2, 2, 2, 256, 252, 3, 2, 2, 2, 256, 253, 3, 2, 2,
	2, 256, 254, 3, 2, 2, 2, 256, 255, 3, 2, 2, 2, 257, 57, 3, 2, 2, 2, 258,
	259, 7, 62, 2, 2, 259, 59, 3, 2, 2, 2, 260, 261, 7, 63, 2, 2, 261, 61,
	3, 2, 2, 2, 262, 263, 7, 64, 2, 2, 263, 63, 3, 2, 2, 2, 264, 265, 5, 58,
	29, 2, 265, 266, 5, 62, 31, 2, 266, 65, 3, 2, 2, 2, 267, 268, 5, 62, 31,
	2, 268, 269, 5, 60, 30, 2, 269, 67, 3, 2, 2, 2, 270, 271, 5, 58, 29, 2,
	271, 272, 5, 60, 30, 2, 272, 69, 3, 2, 2, 2, 273, 274, 5, 42, 21, 2, 274,
	275, 5, 38, 19, 2, 275, 276, 5, 44, 22, 2, 276, 277, 5, 12, 6, 2, 277,
	285, 3, 2, 2, 2, 278, 279, 5, 14, 7, 2, 279, 280, 5, 4, 2, 2, 280, 281,
	5, 26, 13, 2, 281, 282, 5, 40, 20, 2, 282, 283, 5, 12, 6, 2, 283, 285,
	3, 2, 2, 2, 284, 273, 3, 2, 2, 2, 284, 278, 3, 2, 2, 2, 285, 71, 3, 2,
	2, 2, 286, 287, 5, 4, 2, 2, 287, 288, 5, 30, 15, 2, 288, 289, 5, 10, 5,
	2, 289, 73, 3, 2, 2, 2, 290, 291, 5, 32, 16, 2, 291, 292, 5, 38, 19, 2,
	292, 75, 3, 2, 2, 2, 293, 294, 5, 30, 15, 2, 294, 295, 5, 32, 16, 2, 295,
	296, 5, 42, 21, 2, 296, 77, 3, 2, 2, 2, 297, 298, 5, 26, 13, 2, 298, 299,
	5, 20, 10, 2, 299, 300, 5, 24, 12, 2, 300, 301, 5, 12, 6, 2, 301, 79, 3,
	2, 2, 2, 302, 303, 5, 20, 10, 2, 303, 304, 5, 26, 13, 2, 304, 305, 5, 20,
	10, 2, 305, 306, 5, 24, 12, 2, 306, 307, 5, 12, 6, 2, 307, 81, 3, 2, 2,
	2, 308, 309, 5, 6, 3, 2, 309, 310, 5, 12, 6, 2, 310, 311, 5, 42, 21, 2,
	311, 312, 5, 48, 24, 2, 312, 313, 5, 12, 6, 2, 313, 314, 5, 12, 6, 2, 314,
	315, 5, 30, 15, 2, 315, 83, 3, 2, 2, 2, 316, 317, 5, 20, 10, 2, 317, 318,
	5, 40, 20, 2, 318, 85, 3, 2, 2, 2, 319, 320, 5, 30, 15, 2, 320, 321, 5,
	44, 22, 2, 321, 322, 5, 26, 13, 2, 322, 323, 5, 26, 13, 2, 323, 87, 3,
	2, 2, 2, 324, 325, 5, 20, 10, 2, 325, 326, 5, 30, 15, 2, 326, 89, 3, 2,
	2, 2, 327, 333, 5, 150, 75, 2, 328, 333, 5, 154, 77, 2, 329, 333, 5, 148,
	74, 2, 330, 333, 5, 158, 79, 2, 331, 333, 5, 134, 67, 2, 332, 327, 3, 2,
	2, 2, 332, 328, 3, 2, 2, 2, 332, 329, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2,
	332, 331, 3, 2, 2, 2, 333, 91, 3, 2, 2, 2, 334, 335, 5, 12, 6, 2, 335,
	336, 5, 36, 18, 2, 336, 337, 5, 44, 22, 2, 337, 338, 5, 4, 2, 2, 338, 339,
	5, 26, 13, 2, 339, 340, 5, 40, 20, 2, 340, 403, 3, 2, 2, 2, 341, 342, 5,
	10, 5, 2, 342, 343, 5, 20, 10, 2, 343, 344, 5, 40, 20, 2, 344, 345, 5,
	22, 11, 2, 345, 346, 5, 32, 16, 2, 346, 347, 5, 20, 10, 2, 347, 348, 5,
	30, 15, 2, 348, 349, 5, 42, 21, 2, 349, 403, 3, 2, 2, 2, 350, 351, 5, 42,
	21, 2, 351, 352, 5, 32, 16, 2, 352, 353, 5, 44, 22, 2, 353, 354, 5, 8,
	4, 2, 354, 355, 5, 18, 9, 2, 355, 356, 5, 12, 6, 2, 356, 357, 5, 40, 20,
	2, 357, 403, 3, 2, 2, 2, 358, 359, 5, 48, 24, 2, 359, 360, 5, 20, 10, 2,
	360, 361, 5, 42, 21, 2, 361, 362, 5, 18, 9, 2, 362, 363, 5, 20, 10, 2,
	363, 364, 5, 30, 15, 2, 364, 403, 3, 2, 2, 2, 365, 366, 5, 32, 16, 2, 366,
	367, 5, 46, 23, 2, 367, 368, 5, 12, 6, 2, 368, 369, 5, 38, 19, 2, 369,
	370, 5, 26, 13, 2, 370, 371, 5, 4, 2, 2, 371, 372, 5, 34, 17, 2, 372, 373,
	5, 40, 20, 2, 373, 403, 3, 2, 2, 2, 374, 375, 5, 8, 4, 2, 375, 376, 5,
	38, 19, 2, 376, 377, 5, 32, 16, 2, 377, 378, 5, 40, 20, 2, 378, 379, 5,
	40, 20, 2, 379, 380, 5, 12, 6, 2, 380, 381, 5, 40, 20, 2, 381, 403, 3,
	2, 2, 2, 382, 383, 5, 20, 10, 2, 383, 384, 5, 30, 15, 2, 384, 385, 5, 42,
	21, 2, 385, 386, 5, 12, 6, 2, 386, 387, 5, 38, 19, 2, 387, 388, 5, 40,
	20, 2, 388, 389, 5, 12, 6, 2, 389, 390, 5, 8, 4, 2, 390, 391, 5, 42, 21,
	2, 391, 392, 5, 40, 20, 2, 392, 403, 3, 2, 2, 2, 393, 394, 5, 8, 4, 2,
	394, 395, 5, 32, 16, 2, 395, 396, 5, 30, 15, 2, 396, 397, 5, 42, 21, 2,
	397, 398, 5, 4, 2, 2, 398, 399, 5, 20, 10, 2, 399, 400, 5, 30, 15, 2, 400,
	401, 5, 40, 20, 2, 401, 403, 3, 2, 2, 2, 402, 334, 3, 2, 2, 2, 402, 341,
	3, 2, 2, 2, 402, 350, 3, 2, 2, 2, 402, 358, 3, 2, 2, 2, 402, 365, 3, 2,
	2, 2, 402, 374, 3, 2, 2, 2, 402, 382, 3, 2, 2, 2, 402, 393, 3, 2, 2, 2,
	403, 93, 3, 2, 2, 2, 404, 405, 5, 10, 5, 2, 405, 406, 5, 48, 24, 2, 406,
	407, 5, 20, 10, 2, 407, 408, 5, 42, 21, 2, 408, 409, 5, 18, 9, 2, 409,
	410, 5, 20, 10, 2, 410, 411, 5, 30, 15, 2, 411, 95, 3, 2, 2, 2, 412, 413,
	5, 34, 17, 2, 413, 414, 5, 32, 16, 2, 414, 415, 5, 20, 10, 2, 415, 416,
	5, 30, 15, 2, 416, 417, 5, 42, 21, 2, 417, 97, 3, 2, 2, 2, 418, 419, 5,
	26, 13, 2, 419, 420, 5, 20, 10, 2, 420, 421, 5, 30, 15, 2, 421, 422, 5,
	12, 6, 2, 422, 423, 5, 40, 20, 2, 423, 424, 5, 42, 21, 2, 424, 425, 5,
	38, 19, 2, 425, 426, 5, 20, 10, 2, 426, 427, 5, 30, 15, 2, 427, 428, 5,
	16, 8, 2, 428, 99, 3, 2, 2, 2, 429, 430, 5, 34, 17, 2, 430, 431, 5, 32,
	16, 2, 431, 432, 5, 26, 13, 2, 432, 433, 5, 52, 26, 2, 433, 434, 5, 16,
	8, 2, 434, 435, 5, 32, 16, 2, 435, 436, 5, 30, 15, 2, 436, 101, 3, 2, 2,
	2, 437, 438, 5, 28, 14, 2, 438, 439, 5, 44, 22, 2, 439, 440, 5, 26, 13,
	2, 440, 441, 5, 42, 21, 2, 441, 442, 5, 20, 10, 2, 442, 443, 5, 34, 17,
	2, 443, 444, 5, 32, 16, 2, 444, 445, 5, 20, 10, 2, 445, 446, 5, 30, 15,
	2, 446, 447, 5, 42, 21, 2, 447, 103, 3, 2, 2, 2, 448, 449, 5, 28, 14, 2,
	449, 450, 5, 44, 22, 2, 450, 451, 5, 26, 13, 2, 451, 452, 5, 42, 21, 2,
	452, 453, 5, 20, 10, 2, 453, 454, 5, 26, 13, 2, 454, 455, 5, 20, 10, 2,
	455, 456, 5, 30, 15, 2, 456, 457, 5, 12, 6, 2, 457, 458, 5, 40, 20, 2,
	458, 459, 5, 42, 21, 2, 459, 460, 5, 38, 19, 2, 460, 461, 5, 20, 10, 2,
	461, 462, 5, 30, 15, 2, 462, 463, 5, 16, 8, 2, 463, 105, 3, 2, 2, 2, 464,
	465, 5, 28, 14, 2, 465, 466, 5, 44, 22, 2, 466, 467, 5, 26, 13, 2, 467,
	468, 5, 42, 21, 2, 468, 469, 5, 20, 10, 2, 469, 470, 5, 34, 17, 2, 470,
	471, 5, 32, 16, 2, 471, 472, 5, 26, 13, 2, 472, 473, 5, 52, 26, 2, 473,
	474, 5, 16, 8, 2, 474, 475, 5, 32, 16, 2, 475, 476, 5, 30, 15, 2, 476,
	107, 3, 2, 2, 2, 477, 478, 5, 16, 8, 2, 478, 479, 5, 12, 6, 2, 479, 480,
	5, 32, 16, 2, 480, 481, 5, 28, 14, 2, 481, 482, 5, 12, 6, 2, 482, 483,
	5, 42, 21, 2, 483, 484, 5, 38, 19, 2, 484, 485, 5, 52, 26, 2, 485, 486,
	5, 8, 4, 2, 486, 487, 5, 32, 16, 2, 487, 488, 5, 26, 13, 2, 488, 489, 5,
	26, 13, 2, 489, 490, 5, 12, 6, 2, 490, 491, 5, 8, 4, 2, 491, 492, 5, 42,
	21, 2, 492, 493, 5, 20, 10, 2, 493, 494, 5, 32, 16, 2, 494, 495, 5, 30,
	15, 2, 495, 109, 3, 2, 2, 2, 496, 497, 5, 12, 6, 2, 497, 498, 5, 30, 15,
	2, 498, 499, 5, 46, 23, 2, 499, 500, 5, 12, 6, 2, 500, 501, 5, 26, 13,
	2, 501, 502, 5, 32, 16, 2, 502, 503, 5, 34, 17, 2, 503, 504, 5, 12, 6,
	2, 504, 111, 3, 2, 2, 2, 505, 508, 5, 172, 86, 2, 506, 508, 5, 174, 87,
	2, 507, 505, 3, 2, 2, 2, 507, 506, 3, 2, 2, 2, 508, 113, 3, 2, 2, 2, 509,
	510, 5, 138, 69, 2, 510, 511, 3, 2, 2, 2, 511, 512, 8, 57, 2, 2, 512, 513,
	8, 57, 3, 2, 513, 115, 3, 2, 2, 2, 514, 518, 5, 118, 59, 2, 515, 517, 5,
	120, 60, 2, 516, 515, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516, 3, 2,
	2, 2, 518, 519, 3, 2, 2, 2, 519, 526, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2,
	521, 522, 5, 132, 66, 2, 522, 523, 5, 116, 58, 2, 523, 524, 5, 132, 66,
	2, 524, 526, 3, 2, 2, 2, 525, 514, 3, 2, 2, 2, 525, 521, 3, 2, 2, 2, 526,
	117, 3, 2, 2, 2, 527, 528, 5, 122, 61, 2, 528, 119, 3, 2, 2, 2, 529, 534,
	5, 122, 61, 2, 530, 534, 5, 124, 62, 2, 531, 534, 5, 130, 65, 2, 532, 534,
	5, 128, 64, 2, 533, 529, 3, 2, 2, 2, 533, 530, 3, 2, 2, 2, 533, 531, 3,
	2, 2, 2, 533, 532, 3, 2, 2, 2, 534, 121, 3, 2, 2, 2, 535, 536, 9, 28, 2,
	2, 536, 123, 3, 2, 2, 2, 537, 538, 9, 29, 2, 2, 538, 125, 3, 2, 2, 2, 539,
	540, 7, 37, 2, 2, 540, 127, 3, 2, 2, 2, 541, 542, 7, 38, 2, 2, 542, 129,
	3, 2, 2, 2, 543, 544, 7, 97, 2, 2, 544, 131, 3, 2, 2, 2, 545, 546, 7, 36,
	2, 2, 546, 133, 3, 2, 2, 2, 547, 548, 7, 39, 2, 2, 548, 135, 3, 2, 2, 2,
	549, 550, 7, 40, 2, 2, 550, 137, 3, 2, 2, 2, 551, 552, 7, 41, 2, 2, 552,
	139, 3, 2, 2, 2, 553, 554, 7, 42, 2, 2, 554, 141, 3, 2, 2, 2, 555, 556,
	7, 43, 2, 2, 556, 143, 3, 2, 2, 2, 557, 558, 7, 93, 2, 2, 558, 145, 3,
	2, 2, 2, 559, 560, 7, 95, 2, 2, 560, 147, 3, 2, 2, 2, 561, 562, 7, 44,
	2, 2, 562, 149, 3, 2, 2, 2, 563, 564, 7, 45, 2, 2, 564, 151, 3, 2, 2, 2,
	565, 566, 7, 46, 2, 2, 566, 153, 3, 2, 2, 2, 567, 568, 7, 47, 2, 2, 568,
	155, 3, 2, 2, 2, 569, 570, 7, 48, 2, 2, 570, 157, 3, 2, 2, 2, 571, 572,
	7, 49, 2, 2, 572, 159, 3, 2, 2, 2, 573, 574, 7, 60, 2, 2, 574, 161, 3,
	2, 2, 2, 575, 576, 7, 61, 2, 2, 576, 163, 3, 2, 2, 2, 577, 578, 7, 65,
	2, 2, 578, 165, 3, 2, 2, 2, 579, 580, 7, 126, 2, 2, 580, 167, 3, 2, 2,
	2, 581, 582, 4, 50, 51, 2, 582, 169, 3, 2, 2, 2, 583, 591, 5, 124, 62,
	2, 584, 591, 5, 4, 2, 2, 585, 591, 5, 6, 3, 2, 586, 591, 5, 8, 4, 2, 587,
	591, 5, 10, 5, 2, 588, 591, 5, 12, 6, 2, 589, 591, 5, 14, 7, 2, 590, 583,
	3, 2, 2, 2, 590, 584, 3, 2, 2, 2, 590, 585, 3, 2, 2, 2, 590, 586, 3, 2,
	2, 2, 590, 587, 3, 2, 2, 2, 590, 588, 3, 2, 2, 2, 590, 589, 3, 2, 2, 2,
	591, 171, 3, 2, 2, 2, 592, 595, 5, 176, 88, 2, 593, 595, 5, 178, 89, 2,
	594, 592, 3, 2, 2, 2, 594, 593, 3, 2, 2, 2, 595, 173, 3, 2, 2, 2, 596,
	598, 5, 188, 94, 2, 597, 596, 3, 2, 2, 2, 597, 598, 3, 2, 2, 2, 598, 599,
	3, 2, 2, 2, 599, 602, 5, 176, 88, 2, 600, 602, 5, 178, 89, 2, 601, 597,
	3, 2, 2, 2, 601, 600, 3, 2, 2, 2, 602, 175, 3, 2, 2, 2, 603, 608, 5, 186,
	93, 2, 604, 606, 5, 156, 78, 2, 605, 607, 5, 186, 93, 2, 606, 605, 3, 2,
	2, 2, 606, 607, 3, 2, 2, 2, 607, 609, 3, 2, 2, 2, 608, 604, 3, 2, 2, 2,
	608, 609, 3, 2, 2, 2, 609, 614, 3, 2, 2, 2, 610, 611, 5, 156, 78, 2, 611,
	612, 5, 186, 93, 2, 612, 614, 3, 2, 2, 2, 613, 603, 3, 2, 2, 2, 613, 610,
	3, 2, 2, 2, 614, 177, 3, 2, 2, 2, 615, 616, 5, 180, 90, 2, 616, 617, 7,
	71, 2, 2, 617, 618, 5, 182, 91, 2, 618, 179, 3, 2, 2, 2, 619, 620, 5, 176,
	88, 2, 620, 181, 3, 2, 2, 2, 621, 622, 5, 184, 92, 2, 622, 183, 3, 2, 2,
	2, 623, 625, 5, 188, 94, 2, 624, 623, 3, 2, 2, 2, 624, 625, 3, 2, 2, 2,
	625, 626, 3, 2, 2, 2, 626, 627, 5, 186, 93, 2, 627, 185, 3, 2, 2, 2, 628,
	630, 5, 124, 62, 2, 629, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631, 629,
	3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 187, 3, 2, 2, 2, 633, 636, 5, 150,
	75, 2, 634, 636, 5, 154, 77, 2, 635, 633, 3, 2, 2, 2, 635, 634, 3, 2, 2,
	2, 636, 189, 3, 2, 2, 2, 637, 639, 9, 30, 2, 2, 638, 637, 3, 2, 2, 2, 639,
	640, 3, 2, 2, 2, 640, 638, 3, 2, 2, 2, 640, 641, 3, 2, 2, 2, 641, 642,
	3, 2, 2, 2, 642, 643, 8, 95, 4, 2, 643, 191, 3, 2, 2, 2, 644, 645, 7, 41,
	2, 2, 645, 646, 3, 2, 2, 2, 646, 647, 8, 96, 5, 2, 647, 193, 3, 2, 2, 2,
	648, 649, 7, 41, 2, 2, 649, 650, 7, 41, 2, 2, 650, 651, 3, 2, 2, 2, 651,
	652, 8, 97, 2, 2, 652, 195, 3, 2, 2, 2, 653, 654, 10, 31, 2, 2, 654, 655,
	3, 2, 2, 2, 655, 656, 8, 98, 2, 2, 656, 197, 3, 2, 2, 2, 23, 2, 3, 256,
	284, 332, 402, 507, 518, 525, 533, 590, 594, 597, 601, 606, 608, 613, 624,
	631, 635, 640, 6, 5, 2, 2, 4, 3, 2, 8, 2, 2, 4, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "STR",
}

var lexerLiteralNames = []string{
	"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "'#'", "'$'", "'_'", "'\"'", "'%'", "'&'", "", "'('", "')'", "'['",
	"']'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "':'", "';'", "'?'", "'|'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "''''",
}

var lexerSymbolicNames = []string{
	"", "ComparisonOperator", "LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral",
	"AND", "OR", "NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "IN", "ArithmeticOperator",
	"SpatialOperator", "DistanceOperator", "POINT", "LINESTRING", "POLYGON",
	"MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON", "GEOMETRYCOLLECTION",
	"ENVELOPE", "NumericLiteral", "Identifier", "IdentifierStart", "IdentifierPart",
	"ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR", "UNDERSCORE", "DOUBLEQUOTE", "PERCENT",
	"AMPERSAND", "QUOTE", "LEFTPAREN", "RIGHTPAREN", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET",
	"ASTERISK", "PLUS", "COMMA", "MINUS", "PERIOD", "SOLIDUS", "COLON", "SEMICOLON",
	"QUESTIONMARK", "VERTICALBAR", "BIT", "HEXIT", "UnsignedNumericLiteral",
	"SignedNumericLiteral", "ExactNumericLiteral", "ApproximateNumericLiteral",
	"Mantissa", "Exponent", "SignedInteger", "UnsignedInteger", "Sign", "WS",
	"CharacterStringLiteral", "QuotedQuote",
}

var lexerRuleNames = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ComparisonOperator",
	"LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral", "AND", "OR",
	"NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "IN", "ArithmeticOperator",
	"SpatialOperator", "DistanceOperator", "POINT", "LINESTRING", "POLYGON",
	"MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON", "GEOMETRYCOLLECTION",
	"ENVELOPE", "NumericLiteral", "CharacterStringLiteralStart", "Identifier",
	"IdentifierStart", "IdentifierPart", "ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR",
	"UNDERSCORE", "DOUBLEQUOTE", "PERCENT", "AMPERSAND", "QUOTE", "LEFTPAREN",
	"RIGHTPAREN", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "ASTERISK", "PLUS",
	"COMMA", "MINUS", "PERIOD", "SOLIDUS", "COLON", "SEMICOLON", "QUESTIONMARK",
	"VERTICALBAR", "BIT", "HEXIT", "UnsignedNumericLiteral", "SignedNumericLiteral",
	"ExactNumericLiteral", "ApproximateNumericLiteral", "Mantissa", "Exponent",
	"SignedInteger", "UnsignedInteger", "Sign", "WS", "CharacterStringLiteral",
	"QuotedQuote", "Character",
}

type CqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCqlLexer(input antlr.CharStream) *CqlLexer {

	l := new(CqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CqlLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CqlLexer tokens.
const (
	CqlLexerComparisonOperator        = 1
	CqlLexerLT                        = 2
	CqlLexerEQ                        = 3
	CqlLexerGT                        = 4
	CqlLexerNEQ                       = 5
	CqlLexerGTEQ                      = 6
	CqlLexerLTEQ                      = 7
	CqlLexerBooleanLiteral            = 8
	CqlLexerAND                       = 9
	CqlLexerOR                        = 10
	CqlLexerNOT                       = 11
	CqlLexerLIKE                      = 12
	CqlLexerILIKE                     = 13
	CqlLexerBETWEEN                   = 14
	CqlLexerIS                        = 15
	CqlLexerNULL                      = 16
	CqlLexerIN                        = 17
	CqlLexerArithmeticOperator        = 18
	CqlLexerSpatialOperator           = 19
	CqlLexerDistanceOperator          = 20
	CqlLexerPOINT                     = 21
	CqlLexerLINESTRING                = 22
	CqlLexerPOLYGON                   = 23
	CqlLexerMULTIPOINT                = 24
	CqlLexerMULTILINESTRING           = 25
	CqlLexerMULTIPOLYGON              = 26
	CqlLexerGEOMETRYCOLLECTION        = 27
	CqlLexerENVELOPE                  = 28
	CqlLexerNumericLiteral            = 29
	CqlLexerIdentifier                = 30
	CqlLexerIdentifierStart           = 31
	CqlLexerIdentifierPart            = 32
	CqlLexerALPHA                     = 33
	CqlLexerDIGIT                     = 34
	CqlLexerOCTOTHORP                 = 35
	CqlLexerDOLLAR                    = 36
	CqlLexerUNDERSCORE                = 37
	CqlLexerDOUBLEQUOTE               = 38
	CqlLexerPERCENT                   = 39
	CqlLexerAMPERSAND                 = 40
	CqlLexerQUOTE                     = 41
	CqlLexerLEFTPAREN                 = 42
	CqlLexerRIGHTPAREN                = 43
	CqlLexerLEFTSQUAREBRACKET         = 44
	CqlLexerRIGHTSQUAREBRACKET        = 45
	CqlLexerASTERISK                  = 46
	CqlLexerPLUS                      = 47
	CqlLexerCOMMA                     = 48
	CqlLexerMINUS                     = 49
	CqlLexerPERIOD                    = 50
	CqlLexerSOLIDUS                   = 51
	CqlLexerCOLON                     = 52
	CqlLexerSEMICOLON                 = 53
	CqlLexerQUESTIONMARK              = 54
	CqlLexerVERTICALBAR               = 55
	CqlLexerBIT                       = 56
	CqlLexerHEXIT                     = 57
	CqlLexerUnsignedNumericLiteral    = 58
	CqlLexerSignedNumericLiteral      = 59
	CqlLexerExactNumericLiteral       = 60
	CqlLexerApproximateNumericLiteral = 61
	CqlLexerMantissa                  = 62
	CqlLexerExponent                  = 63
	CqlLexerSignedInteger             = 64
	CqlLexerUnsignedInteger           = 65
	CqlLexerSign                      = 66
	CqlLexerWS                        = 67
	CqlLexerCharacterStringLiteral    = 68
	CqlLexerQuotedQuote               = 69
)

// CqlLexerSTR is the CqlLexer mode.
const CqlLexerSTR = 1