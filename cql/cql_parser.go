// Generated from CQLParser.g4 by ANTLR 4.7.

package cql // CQLParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 71, 325,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 78,
	10, 3, 12, 3, 14, 3, 81, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 89, 10, 4, 12, 4, 14, 4, 92, 11, 4, 3, 5, 5, 5, 95, 10, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 104, 10, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 113, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 5, 9, 121, 10, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 128, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 138, 10,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 148,
	10, 12, 3, 12, 3, 12, 3, 12, 7, 12, 153, 10, 12, 12, 12, 14, 12, 156, 11,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 162, 10, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 5, 20, 190, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 5, 21, 200, 10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 7, 26, 219, 10, 26, 12, 26, 14, 26, 222, 11, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 231, 10, 27, 12, 27, 14, 27,
	234, 11, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 243,
	10, 28, 12, 28, 14, 28, 246, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 7, 29, 255, 10, 29, 12, 29, 14, 29, 258, 11, 29, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 267, 10, 30, 12, 30, 14,
	30, 270, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 289,
	10, 32, 12, 32, 14, 32, 292, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 5, 34, 301, 10, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7,
	34, 308, 10, 34, 12, 34, 14, 34, 311, 11, 34, 3, 34, 3, 34, 3, 34, 7, 34,
	316, 10, 34, 12, 34, 14, 34, 319, 11, 34, 5, 34, 321, 10, 34, 3, 34, 3,
	34, 3, 34, 2, 5, 4, 6, 22, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 2, 3, 3, 2, 14, 15, 2, 327, 2, 68, 3, 2, 2, 2, 4, 71, 3,
	2, 2, 2, 6, 82, 3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10, 103, 3, 2, 2, 2, 12,
	112, 3, 2, 2, 2, 14, 114, 3, 2, 2, 2, 16, 118, 3, 2, 2, 2, 18, 125, 3,
	2, 2, 2, 20, 134, 3, 2, 2, 2, 22, 147, 3, 2, 2, 2, 24, 161, 3, 2, 2, 2,
	26, 163, 3, 2, 2, 2, 28, 165, 3, 2, 2, 2, 30, 167, 3, 2, 2, 2, 32, 169,
	3, 2, 2, 2, 34, 171, 3, 2, 2, 2, 36, 178, 3, 2, 2, 2, 38, 189, 3, 2, 2,
	2, 40, 199, 3, 2, 2, 2, 42, 201, 3, 2, 2, 2, 44, 204, 3, 2, 2, 2, 46, 208,
	3, 2, 2, 2, 48, 211, 3, 2, 2, 2, 50, 214, 3, 2, 2, 2, 52, 225, 3, 2, 2,
	2, 54, 237, 3, 2, 2, 2, 56, 249, 3, 2, 2, 2, 58, 261, 3, 2, 2, 2, 60, 273,
	3, 2, 2, 2, 62, 284, 3, 2, 2, 2, 64, 295, 3, 2, 2, 2, 66, 298, 3, 2, 2,
	2, 68, 69, 5, 4, 3, 2, 69, 70, 7, 2, 2, 3, 70, 3, 3, 2, 2, 2, 71, 72, 8,
	3, 1, 2, 72, 73, 5, 6, 4, 2, 73, 79, 3, 2, 2, 2, 74, 75, 12, 3, 2, 2, 75,
	76, 7, 12, 2, 2, 76, 78, 5, 6, 4, 2, 77, 74, 3, 2, 2, 2, 78, 81, 3, 2,
	2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 5, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 82, 83, 8, 4, 1, 2, 83, 84, 5, 8, 5, 2, 84, 90, 3, 2, 2, 2,
	85, 86, 12, 3, 2, 2, 86, 87, 7, 11, 2, 2, 87, 89, 5, 8, 5, 2, 88, 85, 3,
	2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91,
	7, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 95, 7, 13, 2, 2, 94, 93, 3, 2, 2,
	2, 94, 95, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 5, 10, 6, 2, 97, 9,
	3, 2, 2, 2, 98, 104, 5, 12, 7, 2, 99, 100, 7, 44, 2, 2, 100, 101, 5, 4,
	3, 2, 101, 102, 7, 45, 2, 2, 102, 104, 3, 2, 2, 2, 103, 98, 3, 2, 2, 2,
	103, 99, 3, 2, 2, 2, 104, 11, 3, 2, 2, 2, 105, 113, 5, 14, 8, 2, 106, 113,
	5, 16, 9, 2, 107, 113, 5, 18, 10, 2, 108, 113, 5, 20, 11, 2, 109, 113,
	5, 66, 34, 2, 110, 113, 5, 34, 18, 2, 111, 113, 5, 36, 19, 2, 112, 105,
	3, 2, 2, 2, 112, 106, 3, 2, 2, 2, 112, 107, 3, 2, 2, 2, 112, 108, 3, 2,
	2, 2, 112, 109, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2,
	113, 13, 3, 2, 2, 2, 114, 115, 5, 22, 12, 2, 115, 116, 7, 3, 2, 2, 116,
	117, 5, 22, 12, 2, 117, 15, 3, 2, 2, 2, 118, 120, 5, 26, 14, 2, 119, 121,
	7, 13, 2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2,
	2, 2, 122, 123, 9, 2, 2, 2, 123, 124, 5, 28, 15, 2, 124, 17, 3, 2, 2, 2,
	125, 127, 5, 26, 14, 2, 126, 128, 7, 13, 2, 2, 127, 126, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 7, 16, 2, 2, 130, 131,
	5, 22, 12, 2, 131, 132, 7, 11, 2, 2, 132, 133, 5, 22, 12, 2, 133, 19, 3,
	2, 2, 2, 134, 135, 5, 26, 14, 2, 135, 137, 7, 17, 2, 2, 136, 138, 7, 13,
	2, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2,
	139, 140, 7, 18, 2, 2, 140, 21, 3, 2, 2, 2, 141, 142, 8, 12, 1, 2, 142,
	148, 5, 24, 13, 2, 143, 144, 7, 44, 2, 2, 144, 145, 5, 22, 12, 2, 145,
	146, 7, 45, 2, 2, 146, 148, 3, 2, 2, 2, 147, 141, 3, 2, 2, 2, 147, 143,
	3, 2, 2, 2, 148, 154, 3, 2, 2, 2, 149, 150, 12, 3, 2, 2, 150, 151, 7, 20,
	2, 2, 151, 153, 5, 22, 12, 4, 152, 149, 3, 2, 2, 2, 153, 156, 3, 2, 2,
	2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 23, 3, 2, 2, 2, 156,
	154, 3, 2, 2, 2, 157, 162, 5, 26, 14, 2, 158, 162, 5, 28, 15, 2, 159, 162,
	5, 30, 16, 2, 160, 162, 5, 32, 17, 2, 161, 157, 3, 2, 2, 2, 161, 158, 3,
	2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 160, 3, 2, 2, 2, 162, 25, 3, 2, 2,
	2, 163, 164, 7, 32, 2, 2, 164, 27, 3, 2, 2, 2, 165, 166, 7, 70, 2, 2, 166,
	29, 3, 2, 2, 2, 167, 168, 7, 31, 2, 2, 168, 31, 3, 2, 2, 2, 169, 170, 7,
	10, 2, 2, 170, 33, 3, 2, 2, 2, 171, 172, 7, 21, 2, 2, 172, 173, 7, 44,
	2, 2, 173, 174, 5, 38, 20, 2, 174, 175, 7, 50, 2, 2, 175, 176, 5, 38, 20,
	2, 176, 177, 7, 45, 2, 2, 177, 35, 3, 2, 2, 2, 178, 179, 7, 22, 2, 2, 179,
	180, 7, 44, 2, 2, 180, 181, 5, 38, 20, 2, 181, 182, 7, 50, 2, 2, 182, 183,
	5, 38, 20, 2, 183, 184, 7, 50, 2, 2, 184, 185, 7, 31, 2, 2, 185, 186, 7,
	45, 2, 2, 186, 37, 3, 2, 2, 2, 187, 190, 5, 26, 14, 2, 188, 190, 5, 40,
	21, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 39, 3, 2, 2, 2,
	191, 200, 5, 42, 22, 2, 192, 200, 5, 46, 24, 2, 193, 200, 5, 48, 25, 2,
	194, 200, 5, 52, 27, 2, 195, 200, 5, 54, 28, 2, 196, 200, 5, 56, 29, 2,
	197, 200, 5, 58, 30, 2, 198, 200, 5, 60, 31, 2, 199, 191, 3, 2, 2, 2, 199,
	192, 3, 2, 2, 2, 199, 193, 3, 2, 2, 2, 199, 194, 3, 2, 2, 2, 199, 195,
	3, 2, 2, 2, 199, 196, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2,
	2, 2, 200, 41, 3, 2, 2, 2, 201, 202, 7, 23, 2, 2, 202, 203, 5, 44, 23,
	2, 203, 43, 3, 2, 2, 2, 204, 205, 7, 44, 2, 2, 205, 206, 5, 64, 33, 2,
	206, 207, 7, 45, 2, 2, 207, 45, 3, 2, 2, 2, 208, 209, 7, 24, 2, 2, 209,
	210, 5, 62, 32, 2, 210, 47, 3, 2, 2, 2, 211, 212, 7, 25, 2, 2, 212, 213,
	5, 50, 26, 2, 213, 49, 3, 2, 2, 2, 214, 215, 7, 44, 2, 2, 215, 220, 5,
	62, 32, 2, 216, 217, 7, 50, 2, 2, 217, 219, 5, 62, 32, 2, 218, 216, 3,
	2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2,
	2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 224, 7, 45, 2, 2, 224,
	51, 3, 2, 2, 2, 225, 226, 7, 26, 2, 2, 226, 227, 7, 44, 2, 2, 227, 232,
	5, 44, 23, 2, 228, 229, 7, 50, 2, 2, 229, 231, 5, 44, 23, 2, 230, 228,
	3, 2, 2, 2, 231, 234, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2,
	2, 2, 233, 235, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 235, 236, 7, 45, 2, 2,
	236, 53, 3, 2, 2, 2, 237, 238, 7, 27, 2, 2, 238, 239, 7, 44, 2, 2, 239,
	244, 5, 62, 32, 2, 240, 241, 7, 50, 2, 2, 241, 243, 5, 62, 32, 2, 242,
	240, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 244, 245,
	3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 247, 248, 7, 45,
	2, 2, 248, 55, 3, 2, 2, 2, 249, 250, 7, 28, 2, 2, 250, 251, 7, 44, 2, 2,
	251, 256, 5, 50, 26, 2, 252, 253, 7, 50, 2, 2, 253, 255, 5, 50, 26, 2,
	254, 252, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 260,
	7, 45, 2, 2, 260, 57, 3, 2, 2, 2, 261, 262, 7, 29, 2, 2, 262, 263, 7, 44,
	2, 2, 263, 268, 5, 40, 21, 2, 264, 265, 7, 50, 2, 2, 265, 267, 5, 40, 21,
	2, 266, 264, 3, 2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268,
	269, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 272,
	7, 45, 2, 2, 272, 59, 3, 2, 2, 2, 273, 274, 7, 30, 2, 2, 274, 275, 7, 44,
	2, 2, 275, 276, 7, 31, 2, 2, 276, 277, 7, 50, 2, 2, 277, 278, 7, 31, 2,
	2, 278, 279, 7, 50, 2, 2, 279, 280, 7, 31, 2, 2, 280, 281, 7, 50, 2, 2,
	281, 282, 7, 31, 2, 2, 282, 283, 7, 45, 2, 2, 283, 61, 3, 2, 2, 2, 284,
	285, 7, 44, 2, 2, 285, 290, 5, 64, 33, 2, 286, 287, 7, 50, 2, 2, 287, 289,
	5, 64, 33, 2, 288, 286, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3,
	2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 293, 3, 2, 2, 2, 292, 290, 3, 2, 2,
	2, 293, 294, 7, 45, 2, 2, 294, 63, 3, 2, 2, 2, 295, 296, 7, 31, 2, 2, 296,
	297, 7, 31, 2, 2, 297, 65, 3, 2, 2, 2, 298, 300, 5, 26, 14, 2, 299, 301,
	7, 13, 2, 2, 300, 299, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 303, 7, 19, 2, 2, 303, 320, 7, 44, 2, 2, 304, 309, 5, 28, 15,
	2, 305, 306, 7, 50, 2, 2, 306, 308, 5, 28, 15, 2, 307, 305, 3, 2, 2, 2,
	308, 311, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310,
	321, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 312, 317, 5, 30, 16, 2, 313, 314,
	7, 50, 2, 2, 314, 316, 5, 30, 16, 2, 315, 313, 3, 2, 2, 2, 316, 319, 3,
	2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 321, 3, 2, 2,
	2, 319, 317, 3, 2, 2, 2, 320, 304, 3, 2, 2, 2, 320, 312, 3, 2, 2, 2, 321,
	322, 3, 2, 2, 2, 322, 323, 7, 45, 2, 2, 323, 67, 3, 2, 2, 2, 25, 79, 90,
	94, 103, 112, 120, 127, 137, 147, 154, 161, 189, 199, 220, 232, 244, 256,
	268, 290, 300, 309, 317, 320,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "'#'", "'$'", "'_'", "'\"'", "'%'", "'&'", "", "'('", "')'", "'['",
	"']'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "':'", "';'", "'?'", "'|'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "''''",
}
var symbolicNames = []string{
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

var ruleNames = []string{
	"cqlFilter", "booleanValueExpression", "booleanTerm", "booleanFactor",
	"booleanPrimary", "predicate", "binaryComparisonPredicate", "likePredicate",
	"betweenPredicate", "isNullPredicate", "scalarExpression", "scalarValue",
	"propertyName", "characterLiteral", "numericLiteral", "booleanLiteral",
	"spatialPredicate", "distancePredicate", "geomExpression", "geomLiteral",
	"point", "pointList", "linestring", "polygon", "polygonDef", "multiPoint",
	"multiLinestring", "multiPolygon", "geometryCollection", "envelope", "coordList",
	"coordinate", "inPredicate",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CQLParser struct {
	*antlr.BaseParser
}

func NewCQLParser(input antlr.TokenStream) *CQLParser {
	this := new(CQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CQLParser.g4"

	return this
}

// CQLParser tokens.
const (
	CQLParserEOF                       = antlr.TokenEOF
	CQLParserComparisonOperator        = 1
	CQLParserLT                        = 2
	CQLParserEQ                        = 3
	CQLParserGT                        = 4
	CQLParserNEQ                       = 5
	CQLParserGTEQ                      = 6
	CQLParserLTEQ                      = 7
	CQLParserBooleanLiteral            = 8
	CQLParserAND                       = 9
	CQLParserOR                        = 10
	CQLParserNOT                       = 11
	CQLParserLIKE                      = 12
	CQLParserILIKE                     = 13
	CQLParserBETWEEN                   = 14
	CQLParserIS                        = 15
	CQLParserNULL                      = 16
	CQLParserIN                        = 17
	CQLParserArithmeticOperator        = 18
	CQLParserSpatialOperator           = 19
	CQLParserDistanceOperator          = 20
	CQLParserPOINT                     = 21
	CQLParserLINESTRING                = 22
	CQLParserPOLYGON                   = 23
	CQLParserMULTIPOINT                = 24
	CQLParserMULTILINESTRING           = 25
	CQLParserMULTIPOLYGON              = 26
	CQLParserGEOMETRYCOLLECTION        = 27
	CQLParserENVELOPE                  = 28
	CQLParserNumericLiteral            = 29
	CQLParserIdentifier                = 30
	CQLParserIdentifierStart           = 31
	CQLParserIdentifierPart            = 32
	CQLParserALPHA                     = 33
	CQLParserDIGIT                     = 34
	CQLParserOCTOTHORP                 = 35
	CQLParserDOLLAR                    = 36
	CQLParserUNDERSCORE                = 37
	CQLParserDOUBLEQUOTE               = 38
	CQLParserPERCENT                   = 39
	CQLParserAMPERSAND                 = 40
	CQLParserQUOTE                     = 41
	CQLParserLEFTPAREN                 = 42
	CQLParserRIGHTPAREN                = 43
	CQLParserLEFTSQUAREBRACKET         = 44
	CQLParserRIGHTSQUAREBRACKET        = 45
	CQLParserASTERISK                  = 46
	CQLParserPLUS                      = 47
	CQLParserCOMMA                     = 48
	CQLParserMINUS                     = 49
	CQLParserPERIOD                    = 50
	CQLParserSOLIDUS                   = 51
	CQLParserCOLON                     = 52
	CQLParserSEMICOLON                 = 53
	CQLParserQUESTIONMARK              = 54
	CQLParserVERTICALBAR               = 55
	CQLParserBIT                       = 56
	CQLParserHEXIT                     = 57
	CQLParserUnsignedNumericLiteral    = 58
	CQLParserSignedNumericLiteral      = 59
	CQLParserExactNumericLiteral       = 60
	CQLParserApproximateNumericLiteral = 61
	CQLParserMantissa                  = 62
	CQLParserExponent                  = 63
	CQLParserSignedInteger             = 64
	CQLParserUnsignedInteger           = 65
	CQLParserSign                      = 66
	CQLParserWS                        = 67
	CQLParserCharacterStringLiteral    = 68
	CQLParserQuotedQuote               = 69
)

// CQLParser rules.
const (
	CQLParserRULE_cqlFilter                 = 0
	CQLParserRULE_booleanValueExpression    = 1
	CQLParserRULE_booleanTerm               = 2
	CQLParserRULE_booleanFactor             = 3
	CQLParserRULE_booleanPrimary            = 4
	CQLParserRULE_predicate                 = 5
	CQLParserRULE_binaryComparisonPredicate = 6
	CQLParserRULE_likePredicate             = 7
	CQLParserRULE_betweenPredicate          = 8
	CQLParserRULE_isNullPredicate           = 9
	CQLParserRULE_scalarExpression          = 10
	CQLParserRULE_scalarValue               = 11
	CQLParserRULE_propertyName              = 12
	CQLParserRULE_characterLiteral          = 13
	CQLParserRULE_numericLiteral            = 14
	CQLParserRULE_booleanLiteral            = 15
	CQLParserRULE_spatialPredicate          = 16
	CQLParserRULE_distancePredicate         = 17
	CQLParserRULE_geomExpression            = 18
	CQLParserRULE_geomLiteral               = 19
	CQLParserRULE_point                     = 20
	CQLParserRULE_pointList                 = 21
	CQLParserRULE_linestring                = 22
	CQLParserRULE_polygon                   = 23
	CQLParserRULE_polygonDef                = 24
	CQLParserRULE_multiPoint                = 25
	CQLParserRULE_multiLinestring           = 26
	CQLParserRULE_multiPolygon              = 27
	CQLParserRULE_geometryCollection        = 28
	CQLParserRULE_envelope                  = 29
	CQLParserRULE_coordList                 = 30
	CQLParserRULE_coordinate                = 31
	CQLParserRULE_inPredicate               = 32
)

// ICqlFilterContext is an interface to support dynamic dispatch.
type ICqlFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCqlFilterContext differentiates from other interfaces.
	IsCqlFilterContext()
}

type CqlFilterContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCqlFilterContext() *CqlFilterContext {
	var p = new(CqlFilterContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_cqlFilter
	return p
}

func (*CqlFilterContext) IsCqlFilterContext() {}

func NewCqlFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlFilterContext {
	var p = new(CqlFilterContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_cqlFilter

	return p
}

func (s *CqlFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlFilterContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *CqlFilterContext) EOF() antlr.TerminalNode {
	return s.GetToken(CQLParserEOF, 0)
}

func (s *CqlFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCqlFilter(s)
	}
}

func (s *CqlFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCqlFilter(s)
	}
}

func (p *CQLParser) CqlFilter() (localctx ICqlFilterContext) {
	localctx = NewCqlFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQLParserRULE_cqlFilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.booleanValueExpression(0)
	}
	{
		p.SetState(67)
		p.Match(CQLParserEOF)
	}

	return localctx
}

// IBooleanValueExpressionContext is an interface to support dynamic dispatch.
type IBooleanValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueExpressionContext differentiates from other interfaces.
	IsBooleanValueExpressionContext()
}

type BooleanValueExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanValueExpressionContext() *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanValueExpression
	return p
}

func (*BooleanValueExpressionContext) IsBooleanValueExpressionContext() {}

func NewBooleanValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanValueExpression

	return p
}

func (s *BooleanValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueExpressionContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanValueExpressionContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanValueExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(CQLParserOR, 0)
}

func (s *BooleanValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanValueExpression(s)
	}
}

func (s *BooleanValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanValueExpression(s)
	}
}

func (p *CQLParser) BooleanValueExpression() (localctx IBooleanValueExpressionContext) {
	return p.booleanValueExpression(0)
}

func (p *CQLParser) booleanValueExpression(_p int) (localctx IBooleanValueExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CQLParserRULE_booleanValueExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.booleanTerm(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanValueExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_booleanValueExpression)
			p.SetState(72)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(73)
				p.Match(CQLParserOR)
			}
			{
				p.SetState(74)
				p.booleanTerm(0)
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanTermContext is an interface to support dynamic dispatch.
type IBooleanTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanTermContext differentiates from other interfaces.
	IsBooleanTermContext()
}

type BooleanTermContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanTermContext() *BooleanTermContext {
	var p = new(BooleanTermContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanTerm
	return p
}

func (*BooleanTermContext) IsBooleanTermContext() {}

func NewBooleanTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTermContext {
	var p = new(BooleanTermContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanTerm

	return p
}

func (s *BooleanTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTermContext) BooleanFactor() IBooleanFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanFactorContext)
}

func (s *BooleanTermContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanTermContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *BooleanTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanTerm(s)
	}
}

func (s *BooleanTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanTerm(s)
	}
}

func (p *CQLParser) BooleanTerm() (localctx IBooleanTermContext) {
	return p.booleanTerm(0)
}

func (p *CQLParser) booleanTerm(_p int) (localctx IBooleanTermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanTermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CQLParserRULE_booleanTerm, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.BooleanFactor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanTermContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_booleanTerm)
			p.SetState(83)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(84)
				p.Match(CQLParserAND)
			}
			{
				p.SetState(85)
				p.BooleanFactor()
			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanFactorContext is an interface to support dynamic dispatch.
type IBooleanFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanFactorContext differentiates from other interfaces.
	IsBooleanFactorContext()
}

type BooleanFactorContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanFactorContext() *BooleanFactorContext {
	var p = new(BooleanFactorContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanFactor
	return p
}

func (*BooleanFactorContext) IsBooleanFactorContext() {}

func NewBooleanFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFactorContext {
	var p = new(BooleanFactorContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanFactor

	return p
}

func (s *BooleanFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFactorContext) BooleanPrimary() IBooleanPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanPrimaryContext)
}

func (s *BooleanFactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *BooleanFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanFactor(s)
	}
}

func (s *BooleanFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanFactor(s)
	}
}

func (p *CQLParser) BooleanFactor() (localctx IBooleanFactorContext) {
	localctx = NewBooleanFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQLParserRULE_booleanFactor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(91)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(94)
		p.BooleanPrimary()
	}

	return localctx
}

// IBooleanPrimaryContext is an interface to support dynamic dispatch.
type IBooleanPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanPrimaryContext differentiates from other interfaces.
	IsBooleanPrimaryContext()
}

type BooleanPrimaryContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanPrimaryContext() *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanPrimary
	return p
}

func (*BooleanPrimaryContext) IsBooleanPrimaryContext() {}

func NewBooleanPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanPrimary

	return p
}

func (s *BooleanPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanPrimaryContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BooleanPrimaryContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *BooleanPrimaryContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanPrimaryContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *BooleanPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanPrimary(s)
	}
}

func (s *BooleanPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanPrimary(s)
	}
}

func (p *CQLParser) BooleanPrimary() (localctx IBooleanPrimaryContext) {
	localctx = NewBooleanPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQLParserRULE_booleanPrimary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Predicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(CQLParserLEFTPAREN)
		}
		{
			p.SetState(98)
			p.booleanValueExpression(0)
		}
		{
			p.SetState(99)
			p.Match(CQLParserRIGHTPAREN)
		}

	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) BinaryComparisonPredicate() IBinaryComparisonPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryComparisonPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryComparisonPredicateContext)
}

func (s *PredicateContext) LikePredicate() ILikePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikePredicateContext)
}

func (s *PredicateContext) BetweenPredicate() IBetweenPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBetweenPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBetweenPredicateContext)
}

func (s *PredicateContext) IsNullPredicate() IIsNullPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsNullPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsNullPredicateContext)
}

func (s *PredicateContext) InPredicate() IInPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInPredicateContext)
}

func (s *PredicateContext) SpatialPredicate() ISpatialPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpatialPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpatialPredicateContext)
}

func (s *PredicateContext) DistancePredicate() IDistancePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDistancePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDistancePredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *CQLParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQLParserRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.BinaryComparisonPredicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.LikePredicate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.BetweenPredicate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.IsNullPredicate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.InPredicate()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.SpatialPredicate()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.DistancePredicate()
		}

	}

	return localctx
}

// IBinaryComparisonPredicateContext is an interface to support dynamic dispatch.
type IBinaryComparisonPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryComparisonPredicateContext differentiates from other interfaces.
	IsBinaryComparisonPredicateContext()
}

type BinaryComparisonPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBinaryComparisonPredicateContext() *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate
	return p
}

func (*BinaryComparisonPredicateContext) IsBinaryComparisonPredicateContext() {}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate

	return p
}

func (s *BinaryComparisonPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryComparisonPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *BinaryComparisonPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BinaryComparisonPredicateContext) ComparisonOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserComparisonOperator, 0)
}

func (s *BinaryComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (p *CQLParser) BinaryComparisonPredicate() (localctx IBinaryComparisonPredicateContext) {
	localctx = NewBinaryComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQLParserRULE_binaryComparisonPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.scalarExpression(0)
	}
	{
		p.SetState(113)
		p.Match(CQLParserComparisonOperator)
	}
	{
		p.SetState(114)
		p.scalarExpression(0)
	}

	return localctx
}

// ILikePredicateContext is an interface to support dynamic dispatch.
type ILikePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikePredicateContext differentiates from other interfaces.
	IsLikePredicateContext()
}

type LikePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyLikePredicateContext() *LikePredicateContext {
	var p = new(LikePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_likePredicate
	return p
}

func (*LikePredicateContext) IsLikePredicateContext() {}

func NewLikePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_likePredicate

	return p
}

func (s *LikePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *LikePredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *LikePredicateContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserLIKE, 0)
}

func (s *LikePredicateContext) ILIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserILIKE, 0)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (p *CQLParser) LikePredicate() (localctx ILikePredicateContext) {
	localctx = NewLikePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLParserRULE_likePredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.PropertyName()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(117)
			p.Match(CQLParserNOT)
		}

	}
	p.SetState(120)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLParserLIKE || _la == CQLParserILIKE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(121)
		p.CharacterLiteral()
	}

	return localctx
}

// IBetweenPredicateContext is an interface to support dynamic dispatch.
type IBetweenPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBetweenPredicateContext differentiates from other interfaces.
	IsBetweenPredicateContext()
}

type BetweenPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBetweenPredicateContext() *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_betweenPredicate
	return p
}

func (*BetweenPredicateContext) IsBetweenPredicateContext() {}

func NewBetweenPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_betweenPredicate

	return p
}

func (s *BetweenPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(CQLParserBETWEEN, 0)
}

func (s *BetweenPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *BetweenPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBetweenPredicate(s)
	}
}

func (p *CQLParser) BetweenPredicate() (localctx IBetweenPredicateContext) {
	localctx = NewBetweenPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQLParserRULE_betweenPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.PropertyName()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(124)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(127)
		p.Match(CQLParserBETWEEN)
	}
	{
		p.SetState(128)
		p.scalarExpression(0)
	}
	{
		p.SetState(129)
		p.Match(CQLParserAND)
	}
	{
		p.SetState(130)
		p.scalarExpression(0)
	}

	return localctx
}

// IIsNullPredicateContext is an interface to support dynamic dispatch.
type IIsNullPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsNullPredicateContext differentiates from other interfaces.
	IsIsNullPredicateContext()
}

type IsNullPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsNullPredicateContext() *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_isNullPredicate
	return p
}

func (*IsNullPredicateContext) IsIsNullPredicateContext() {}

func NewIsNullPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_isNullPredicate

	return p
}

func (s *IsNullPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsNullPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(CQLParserIS, 0)
}

func (s *IsNullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(CQLParserNULL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (p *CQLParser) IsNullPredicate() (localctx IIsNullPredicateContext) {
	localctx = NewIsNullPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLParserRULE_isNullPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.PropertyName()
	}
	{
		p.SetState(133)
		p.Match(CQLParserIS)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(134)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(137)
		p.Match(CQLParserNULL)
	}

	return localctx
}

// IScalarExpressionContext is an interface to support dynamic dispatch.
type IScalarExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarExpressionContext differentiates from other interfaces.
	IsScalarExpressionContext()
}

type ScalarExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarExpressionContext() *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_scalarExpression
	return p
}

func (*ScalarExpressionContext) IsScalarExpressionContext() {}

func NewScalarExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarExpression

	return p
}

func (s *ScalarExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarExpressionContext) ScalarValue() IScalarValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarValueContext)
}

func (s *ScalarExpressionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *ScalarExpressionContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *ScalarExpressionContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *ScalarExpressionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *ScalarExpressionContext) ArithmeticOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserArithmeticOperator, 0)
}

func (s *ScalarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarExpression(s)
	}
}

func (s *ScalarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarExpression(s)
	}
}

func (p *CQLParser) ScalarExpression() (localctx IScalarExpressionContext) {
	return p.scalarExpression(0)
}

func (p *CQLParser) scalarExpression(_p int) (localctx IScalarExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewScalarExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalarExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, CQLParserRULE_scalarExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserBooleanLiteral, CQLParserNumericLiteral, CQLParserIdentifier, CQLParserCharacterStringLiteral:
		{
			p.SetState(140)
			p.ScalarValue()
		}

	case CQLParserLEFTPAREN:
		{
			p.SetState(141)
			p.Match(CQLParserLEFTPAREN)
		}
		{
			p.SetState(142)
			p.scalarExpression(0)
		}
		{
			p.SetState(143)
			p.Match(CQLParserRIGHTPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewScalarExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_scalarExpression)
			p.SetState(147)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(148)
				p.Match(CQLParserArithmeticOperator)
			}
			{
				p.SetState(149)
				p.scalarExpression(2)
			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IScalarValueContext is an interface to support dynamic dispatch.
type IScalarValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarValueContext differentiates from other interfaces.
	IsScalarValueContext()
}

type ScalarValueContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarValueContext() *ScalarValueContext {
	var p = new(ScalarValueContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_scalarValue
	return p
}

func (*ScalarValueContext) IsScalarValueContext() {}

func NewScalarValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarValueContext {
	var p = new(ScalarValueContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarValue

	return p
}

func (s *ScalarValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarValueContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *ScalarValueContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *ScalarValueContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *ScalarValueContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ScalarValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarValue(s)
	}
}

func (s *ScalarValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarValue(s)
	}
}

func (p *CQLParser) ScalarValue() (localctx IScalarValueContext) {
	localctx = NewScalarValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CQLParserRULE_scalarValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.PropertyName()
		}

	case CQLParserCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.CharacterLiteral()
		}

	case CQLParserNumericLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.NumericLiteral()
		}

	case CQLParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.BooleanLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLParserIdentifier, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *CQLParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CQLParserRULE_propertyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(CQLParserIdentifier)
	}

	return localctx
}

// ICharacterLiteralContext is an interface to support dynamic dispatch.
type ICharacterLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterLiteralContext differentiates from other interfaces.
	IsCharacterLiteralContext()
}

type CharacterLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCharacterLiteralContext() *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_characterLiteral
	return p
}

func (*CharacterLiteralContext) IsCharacterLiteralContext() {}

func NewCharacterLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_characterLiteral

	return p
}

func (s *CharacterLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterLiteralContext) CharacterStringLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserCharacterStringLiteral, 0)
}

func (s *CharacterLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCharacterLiteral(s)
	}
}

func (s *CharacterLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCharacterLiteral(s)
	}
}

func (p *CQLParser) CharacterLiteral() (localctx ICharacterLiteralContext) {
	localctx = NewCharacterLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQLParserRULE_characterLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(CQLParserCharacterStringLiteral)
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *CQLParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQLParserRULE_numericLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(CQLParserNumericLiteral)
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *CQLParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQLParserRULE_booleanLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(CQLParserBooleanLiteral)
	}

	return localctx
}

// ISpatialPredicateContext is an interface to support dynamic dispatch.
type ISpatialPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpatialPredicateContext differentiates from other interfaces.
	IsSpatialPredicateContext()
}

type SpatialPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptySpatialPredicateContext() *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_spatialPredicate
	return p
}

func (*SpatialPredicateContext) IsSpatialPredicateContext() {}

func NewSpatialPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_spatialPredicate

	return p
}

func (s *SpatialPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpatialPredicateContext) SpatialOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserSpatialOperator, 0)
}

func (s *SpatialPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *SpatialPredicateContext) AllGeomExpression() []IGeomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem())
	var tst = make([]IGeomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomExpressionContext)
		}
	}

	return tst
}

func (s *SpatialPredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *SpatialPredicateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, 0)
}

func (s *SpatialPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *SpatialPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpatialPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpatialPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterSpatialPredicate(s)
	}
}

func (s *SpatialPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitSpatialPredicate(s)
	}
}

func (p *CQLParser) SpatialPredicate() (localctx ISpatialPredicateContext) {
	localctx = NewSpatialPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLParserRULE_spatialPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(CQLParserSpatialOperator)
	}
	{
		p.SetState(170)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(171)
		p.GeomExpression()
	}
	{
		p.SetState(172)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(173)
		p.GeomExpression()
	}
	{
		p.SetState(174)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IDistancePredicateContext is an interface to support dynamic dispatch.
type IDistancePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDistancePredicateContext differentiates from other interfaces.
	IsDistancePredicateContext()
}

type DistancePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyDistancePredicateContext() *DistancePredicateContext {
	var p = new(DistancePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_distancePredicate
	return p
}

func (*DistancePredicateContext) IsDistancePredicateContext() {}

func NewDistancePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistancePredicateContext {
	var p = new(DistancePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_distancePredicate

	return p
}

func (s *DistancePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DistancePredicateContext) DistanceOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserDistanceOperator, 0)
}

func (s *DistancePredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *DistancePredicateContext) AllGeomExpression() []IGeomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem())
	var tst = make([]IGeomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomExpressionContext)
		}
	}

	return tst
}

func (s *DistancePredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *DistancePredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *DistancePredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *DistancePredicateContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *DistancePredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *DistancePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistancePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistancePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterDistancePredicate(s)
	}
}

func (s *DistancePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitDistancePredicate(s)
	}
}

func (p *CQLParser) DistancePredicate() (localctx IDistancePredicateContext) {
	localctx = NewDistancePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQLParserRULE_distancePredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(CQLParserDistanceOperator)
	}
	{
		p.SetState(177)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(178)
		p.GeomExpression()
	}
	{
		p.SetState(179)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(180)
		p.GeomExpression()
	}
	{
		p.SetState(181)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(182)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(183)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IGeomExpressionContext is an interface to support dynamic dispatch.
type IGeomExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeomExpressionContext differentiates from other interfaces.
	IsGeomExpressionContext()
}

type GeomExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomExpressionContext() *GeomExpressionContext {
	var p = new(GeomExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geomExpression
	return p
}

func (*GeomExpressionContext) IsGeomExpressionContext() {}

func NewGeomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomExpressionContext {
	var p = new(GeomExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomExpression

	return p
}

func (s *GeomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomExpressionContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GeomExpressionContext) GeomLiteral() IGeomLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomExpression(s)
	}
}

func (s *GeomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomExpression(s)
	}
}

func (p *CQLParser) GeomExpression() (localctx IGeomExpressionContext) {
	localctx = NewGeomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQLParserRULE_geomExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.PropertyName()
		}

	case CQLParserPOINT, CQLParserLINESTRING, CQLParserPOLYGON, CQLParserMULTIPOINT, CQLParserMULTILINESTRING, CQLParserMULTIPOLYGON, CQLParserGEOMETRYCOLLECTION, CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.GeomLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGeomLiteralContext is an interface to support dynamic dispatch.
type IGeomLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeomLiteralContext differentiates from other interfaces.
	IsGeomLiteralContext()
}

type GeomLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomLiteralContext() *GeomLiteralContext {
	var p = new(GeomLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geomLiteral
	return p
}

func (*GeomLiteralContext) IsGeomLiteralContext() {}

func NewGeomLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomLiteralContext {
	var p = new(GeomLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomLiteral

	return p
}

func (s *GeomLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomLiteralContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *GeomLiteralContext) Linestring() ILinestringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinestringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinestringContext)
}

func (s *GeomLiteralContext) Polygon() IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *GeomLiteralContext) MultiPoint() IMultiPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPointContext)
}

func (s *GeomLiteralContext) MultiLinestring() IMultiLinestringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiLinestringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiLinestringContext)
}

func (s *GeomLiteralContext) MultiPolygon() IMultiPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPolygonContext)
}

func (s *GeomLiteralContext) GeometryCollection() IGeometryCollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeometryCollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeometryCollectionContext)
}

func (s *GeomLiteralContext) Envelope() IEnvelopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnvelopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnvelopeContext)
}

func (s *GeomLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomLiteral(s)
	}
}

func (s *GeomLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomLiteral(s)
	}
}

func (p *CQLParser) GeomLiteral() (localctx IGeomLiteralContext) {
	localctx = NewGeomLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CQLParserRULE_geomLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserPOINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Point()
		}

	case CQLParserLINESTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Linestring()
		}

	case CQLParserPOLYGON:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Polygon()
		}

	case CQLParserMULTIPOINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
			p.MultiPoint()
		}

	case CQLParserMULTILINESTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.MultiLinestring()
		}

	case CQLParserMULTIPOLYGON:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(194)
			p.MultiPolygon()
		}

	case CQLParserGEOMETRYCOLLECTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(195)
			p.GeometryCollection()
		}

	case CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(196)
			p.Envelope()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_point
	return p
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) POINT() antlr.TerminalNode {
	return s.GetToken(CQLParserPOINT, 0)
}

func (s *PointContext) PointList() IPointListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *CQLParser) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CQLParserRULE_point)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(CQLParserPOINT)
	}
	{
		p.SetState(200)
		p.PointList()
	}

	return localctx
}

// IPointListContext is an interface to support dynamic dispatch.
type IPointListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointListContext differentiates from other interfaces.
	IsPointListContext()
}

type PointListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointListContext() *PointListContext {
	var p = new(PointListContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_pointList
	return p
}

func (*PointListContext) IsPointListContext() {}

func NewPointListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointListContext {
	var p = new(PointListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_pointList

	return p
}

func (s *PointListContext) GetParser() antlr.Parser { return s.parser }

func (s *PointListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PointListContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *PointListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PointListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPointList(s)
	}
}

func (s *PointListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPointList(s)
	}
}

func (p *CQLParser) PointList() (localctx IPointListContext) {
	localctx = NewPointListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQLParserRULE_pointList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(203)
		p.Coordinate()
	}
	{
		p.SetState(204)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// ILinestringContext is an interface to support dynamic dispatch.
type ILinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinestringContext differentiates from other interfaces.
	IsLinestringContext()
}

type LinestringContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyLinestringContext() *LinestringContext {
	var p = new(LinestringContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_linestring
	return p
}

func (*LinestringContext) IsLinestringContext() {}

func NewLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinestringContext {
	var p = new(LinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_linestring

	return p
}

func (s *LinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *LinestringContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserLINESTRING, 0)
}

func (s *LinestringContext) CoordList() ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *LinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterLinestring(s)
	}
}

func (s *LinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitLinestring(s)
	}
}

func (p *CQLParser) Linestring() (localctx ILinestringContext) {
	localctx = NewLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CQLParserRULE_linestring)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(CQLParserLINESTRING)
	}
	{
		p.SetState(207)
		p.CoordList()
	}

	return localctx
}

// IPolygonContext is an interface to support dynamic dispatch.
type IPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonContext differentiates from other interfaces.
	IsPolygonContext()
}

type PolygonContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_polygon
	return p
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserPOLYGON, 0)
}

func (s *PolygonContext) PolygonDef() IPolygonDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *PolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *CQLParser) Polygon() (localctx IPolygonContext) {
	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CQLParserRULE_polygon)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(CQLParserPOLYGON)
	}
	{
		p.SetState(210)
		p.PolygonDef()
	}

	return localctx
}

// IPolygonDefContext is an interface to support dynamic dispatch.
type IPolygonDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonDefContext differentiates from other interfaces.
	IsPolygonDefContext()
}

type PolygonDefContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonDefContext() *PolygonDefContext {
	var p = new(PolygonDefContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_polygonDef
	return p
}

func (*PolygonDefContext) IsPolygonDefContext() {}

func NewPolygonDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonDefContext {
	var p = new(PolygonDefContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygonDef

	return p
}

func (s *PolygonDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonDefContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PolygonDefContext) AllCoordList() []ICoordListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordListContext)(nil)).Elem())
	var tst = make([]ICoordListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordListContext)
		}
	}

	return tst
}

func (s *PolygonDefContext) CoordList(i int) ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *PolygonDefContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PolygonDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *PolygonDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *PolygonDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygonDef(s)
	}
}

func (s *PolygonDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygonDef(s)
	}
}

func (p *CQLParser) PolygonDef() (localctx IPolygonDefContext) {
	localctx = NewPolygonDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CQLParserRULE_polygonDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(213)
		p.CoordList()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(214)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(215)
			p.CoordList()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IMultiPointContext is an interface to support dynamic dispatch.
type IMultiPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPointContext differentiates from other interfaces.
	IsMultiPointContext()
}

type MultiPointContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPointContext() *MultiPointContext {
	var p = new(MultiPointContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiPoint
	return p
}

func (*MultiPointContext) IsMultiPointContext() {}

func NewMultiPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointContext {
	var p = new(MultiPointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPoint

	return p
}

func (s *MultiPointContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOINT, 0)
}

func (s *MultiPointContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiPointContext) AllPointList() []IPointListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointListContext)(nil)).Elem())
	var tst = make([]IPointListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointListContext)
		}
	}

	return tst
}

func (s *MultiPointContext) PointList(i int) IPointListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *MultiPointContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPointContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPointContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPoint(s)
	}
}

func (s *MultiPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPoint(s)
	}
}

func (p *CQLParser) MultiPoint() (localctx IMultiPointContext) {
	localctx = NewMultiPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CQLParserRULE_multiPoint)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(CQLParserMULTIPOINT)
	}
	{
		p.SetState(224)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(225)
		p.PointList()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(226)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(227)
			p.PointList()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IMultiLinestringContext is an interface to support dynamic dispatch.
type IMultiLinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiLinestringContext differentiates from other interfaces.
	IsMultiLinestringContext()
}

type MultiLinestringContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiLinestringContext() *MultiLinestringContext {
	var p = new(MultiLinestringContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiLinestring
	return p
}

func (*MultiLinestringContext) IsMultiLinestringContext() {}

func NewMultiLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLinestringContext {
	var p = new(MultiLinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiLinestring

	return p
}

func (s *MultiLinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLinestringContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTILINESTRING, 0)
}

func (s *MultiLinestringContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiLinestringContext) AllCoordList() []ICoordListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordListContext)(nil)).Elem())
	var tst = make([]ICoordListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordListContext)
		}
	}

	return tst
}

func (s *MultiLinestringContext) CoordList(i int) ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *MultiLinestringContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiLinestringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiLinestringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiLinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiLinestring(s)
	}
}

func (s *MultiLinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiLinestring(s)
	}
}

func (p *CQLParser) MultiLinestring() (localctx IMultiLinestringContext) {
	localctx = NewMultiLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CQLParserRULE_multiLinestring)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(CQLParserMULTILINESTRING)
	}
	{
		p.SetState(236)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(237)
		p.CoordList()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(238)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(239)
			p.CoordList()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(245)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IMultiPolygonContext is an interface to support dynamic dispatch.
type IMultiPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPolygonContext differentiates from other interfaces.
	IsMultiPolygonContext()
}

type MultiPolygonContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonContext() *MultiPolygonContext {
	var p = new(MultiPolygonContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiPolygon
	return p
}

func (*MultiPolygonContext) IsMultiPolygonContext() {}

func NewMultiPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonContext {
	var p = new(MultiPolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPolygon

	return p
}

func (s *MultiPolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOLYGON, 0)
}

func (s *MultiPolygonContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiPolygonContext) AllPolygonDef() []IPolygonDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem())
	var tst = make([]IPolygonDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonDefContext)
		}
	}

	return tst
}

func (s *MultiPolygonContext) PolygonDef(i int) IPolygonDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *MultiPolygonContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPolygon(s)
	}
}

func (s *MultiPolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPolygon(s)
	}
}

func (p *CQLParser) MultiPolygon() (localctx IMultiPolygonContext) {
	localctx = NewMultiPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CQLParserRULE_multiPolygon)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(CQLParserMULTIPOLYGON)
	}
	{
		p.SetState(248)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(249)
		p.PolygonDef()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(250)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(251)
			p.PolygonDef()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IGeometryCollectionContext is an interface to support dynamic dispatch.
type IGeometryCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryCollectionContext differentiates from other interfaces.
	IsGeometryCollectionContext()
}

type GeometryCollectionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeometryCollectionContext() *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geometryCollection
	return p
}

func (*GeometryCollectionContext) IsGeometryCollectionContext() {}

func NewGeometryCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geometryCollection

	return p
}

func (s *GeometryCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryCollectionContext) GEOMETRYCOLLECTION() antlr.TerminalNode {
	return s.GetToken(CQLParserGEOMETRYCOLLECTION, 0)
}

func (s *GeometryCollectionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *GeometryCollectionContext) AllGeomLiteral() []IGeomLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem())
	var tst = make([]IGeomLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomLiteralContext)
		}
	}

	return tst
}

func (s *GeometryCollectionContext) GeomLiteral(i int) IGeomLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeometryCollectionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *GeometryCollectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *GeometryCollectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *GeometryCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeometryCollection(s)
	}
}

func (s *GeometryCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeometryCollection(s)
	}
}

func (p *CQLParser) GeometryCollection() (localctx IGeometryCollectionContext) {
	localctx = NewGeometryCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CQLParserRULE_geometryCollection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(CQLParserGEOMETRYCOLLECTION)
	}
	{
		p.SetState(260)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(261)
		p.GeomLiteral()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(262)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(263)
			p.GeomLiteral()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IEnvelopeContext is an interface to support dynamic dispatch.
type IEnvelopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnvelopeContext differentiates from other interfaces.
	IsEnvelopeContext()
}

type EnvelopeContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyEnvelopeContext() *EnvelopeContext {
	var p = new(EnvelopeContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_envelope
	return p
}

func (*EnvelopeContext) IsEnvelopeContext() {}

func NewEnvelopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvelopeContext {
	var p = new(EnvelopeContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_envelope

	return p
}

func (s *EnvelopeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvelopeContext) ENVELOPE() antlr.TerminalNode {
	return s.GetToken(CQLParserENVELOPE, 0)
}

func (s *EnvelopeContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *EnvelopeContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *EnvelopeContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *EnvelopeContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *EnvelopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvelopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvelopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterEnvelope(s)
	}
}

func (s *EnvelopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitEnvelope(s)
	}
}

func (p *CQLParser) Envelope() (localctx IEnvelopeContext) {
	localctx = NewEnvelopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CQLParserRULE_envelope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(CQLParserENVELOPE)
	}
	{
		p.SetState(272)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(273)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(274)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(275)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(276)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(277)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(278)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(279)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(280)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// ICoordListContext is an interface to support dynamic dispatch.
type ICoordListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordListContext differentiates from other interfaces.
	IsCoordListContext()
}

type CoordListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordListContext() *CoordListContext {
	var p = new(CoordListContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_coordList
	return p
}

func (*CoordListContext) IsCoordListContext() {}

func NewCoordListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordListContext {
	var p = new(CoordListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordList

	return p
}

func (s *CoordListContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *CoordListContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *CoordListContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *CoordListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *CoordListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *CoordListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *CoordListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordList(s)
	}
}

func (s *CoordListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordList(s)
	}
}

func (p *CQLParser) CoordList() (localctx ICoordListContext) {
	localctx = NewCoordListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CQLParserRULE_coordList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(283)
		p.Coordinate()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(284)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(285)
			p.Coordinate()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *CoordinateContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordinate(s)
	}
}

func (s *CoordinateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordinate(s)
	}
}

func (p *CQLParser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CQLParserRULE_coordinate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(294)
		p.Match(CQLParserNumericLiteral)
	}

	return localctx
}

// IInPredicateContext is an interface to support dynamic dispatch.
type IInPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInPredicateContext differentiates from other interfaces.
	IsInPredicateContext()
}

type InPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyInPredicateContext() *InPredicateContext {
	var p = new(InPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_inPredicate
	return p
}

func (*InPredicateContext) IsInPredicateContext() {}

func NewInPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InPredicateContext {
	var p = new(InPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_inPredicate

	return p
}

func (s *InPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *InPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(CQLParserIN, 0)
}

func (s *InPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *InPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *InPredicateContext) AllCharacterLiteral() []ICharacterLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem())
	var tst = make([]ICharacterLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterLiteralContext)
		}
	}

	return tst
}

func (s *InPredicateContext) CharacterLiteral(i int) ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *InPredicateContext) AllNumericLiteral() []INumericLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem())
	var tst = make([]INumericLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralContext)
		}
	}

	return tst
}

func (s *InPredicateContext) NumericLiteral(i int) INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *InPredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *InPredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterInPredicate(s)
	}
}

func (s *InPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitInPredicate(s)
	}
}

func (p *CQLParser) InPredicate() (localctx IInPredicateContext) {
	localctx = NewInPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CQLParserRULE_inPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.PropertyName()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(297)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(300)
		p.Match(CQLParserIN)
	}
	{
		p.SetState(301)
		p.Match(CQLParserLEFTPAREN)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserCharacterStringLiteral:
		{
			p.SetState(302)
			p.CharacterLiteral()
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(303)
				p.Match(CQLParserCOMMA)
			}
			{
				p.SetState(304)
				p.CharacterLiteral()
			}

			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case CQLParserNumericLiteral:
		{
			p.SetState(310)
			p.NumericLiteral()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(311)
				p.Match(CQLParserCOMMA)
			}
			{
				p.SetState(312)
				p.NumericLiteral()
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(320)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

func (p *CQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *BooleanValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanValueExpressionContext)
		}
		return p.BooleanValueExpression_Sempred(t, predIndex)

	case 2:
		var t *BooleanTermContext = nil
		if localctx != nil {
			t = localctx.(*BooleanTermContext)
		}
		return p.BooleanTerm_Sempred(t, predIndex)

	case 10:
		var t *ScalarExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ScalarExpressionContext)
		}
		return p.ScalarExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CQLParser) BooleanValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CQLParser) BooleanTerm_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CQLParser) ScalarExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
