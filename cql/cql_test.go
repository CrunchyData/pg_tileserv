package cql

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestPredicate(t *testing.T) {
	checkCQL(t, "", "")
	checkCQL(t, "id > tt", "\"id\" > \"tt\"")
	checkCQL(t, "id > 1", "\"id\" > 1")
	checkCQL(t, "id >= 1", "\"id\" >= 1")
	checkCQL(t, "id < 1", "\"id\" < 1")
	checkCQL(t, "id <= 1", "\"id\" <= 1")
	checkCQL(t, "id = 1", "\"id\" = 1")
	checkCQL(t, "id <> 1", "\"id\" <> 1")

	checkCQL(t, "id = -1.2345", "\"id\" = -1.2345")
	checkCQL(t, "id = id2", "\"id\" = \"id2\"")
	checkCQL(t, "id = 'foo'", "\"id\" = 'foo'")

	checkCQL(t, "id LIKE 'foo'", "\"id\" LIKE 'foo'")
	checkCQL(t, "id ILIKE 'foo'", "\"id\" ILIKE 'foo'")
	checkCQL(t, "id ILIKE '%Ca%'", "\"id\" ILIKE '%Ca%'")

	checkCQL(t, "id BETWEEN 1 and 2", "\"id\" BETWEEN 1 AND 2")
	checkCQL(t, "id NOT BETWEEN 1 and 2", "\"id\" NOT BETWEEN 1 AND 2")

	checkCQL(t, "id IN (1,2,3)", "\"id\" IN (1,2,3)")
	checkCQL(t, "id NOT IN (1,2,3)", "\"id\" NOT IN (1,2,3)")
	checkCQL(t, "id IN ('a','b','c')", "\"id\" IN ('a','b','c')")

	checkCQL(t, "id IS NULL", "\"id\" IS NULL")
	checkCQL(t, "id IS NOT NULL", "\"id\" IS NOT NULL")
}
func TestSpatialPredicate(t *testing.T) {
	checkCQL(t, "crosses(geom, POINT(0 0))", "ST_Crosses(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "Contains(geom, POINT(0 0))", "ST_Contains(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "DISJOINT(geom, POINT(0 0))", "ST_Disjoint(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "EQUALS(geom, POINT(0 0))", "ST_Equals(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "INTERSECTS(geom, POINT(0 0))", "ST_Intersects(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "OVERLAPS(geom, POINT(0 0))", "ST_Overlaps(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "TOUCHES(geom, POINT(0 0))", "ST_Touches(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "within(geom, POINT(0 0))", "ST_Within(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")

	checkCQL(t, "Dwithin(geom, POINT(0 0), 100)", "ST_DWithin(\"geom\",'SRID=4326;POINT(0 0)'::geometry,100)")
}

func TestArithmetic(t *testing.T) {
	checkCQL(t, "p > 1 + x", "\"p\" > 1 + \"x\"")
	checkCQL(t, "p > 2 * 3 + x", "\"p\" > 2 * 3 + \"x\"")
	checkCQL(t, "p > 2 * (3 + x)", "\"p\" > 2 * (3 + \"x\")")
	checkCQL(t, "p > (y + 5) / (3 - x)", "\"p\" > (\"y\" + 5) / (3 - \"x\")")
	checkCQL(t, "p = x % 10", "\"p\" = \"x\" % 10")
	checkCQL(t, "p BETWEEN x + 10 AND x * 2", "\"p\" BETWEEN \"x\" + 10 AND \"x\" * 2")
	checkCQL(t, "p BETWEEN 2 * (1 + 1000000) AND 900000", "\"p\" BETWEEN 2 * (1 + 1000000) AND 900000")
}

func TestGeometryLiteral(t *testing.T) {
	checkCQL(t, "equals(geom, POINT(0 0))",
		"ST_Equals(\"geom\",'SRID=4326;POINT(0 0)'::geometry)")
	checkCQL(t, "equals(geom, LINESTRING(0 0, 1 1))",
		"ST_Equals(\"geom\",'SRID=4326;LINESTRING(0 0,1 1)'::geometry)")
	checkCQL(t, "equals(geom, POLYGON((0 0, 0 9, 9 0, 0 0)))",
		"ST_Equals(\"geom\",'SRID=4326;POLYGON((0 0,0 9,9 0,0 0))'::geometry)")
	checkCQL(t, "equals(geom, POLYGON((0 0, 0 9, 9 0, 0 0),(1 1, 1 8, 8 1, 1 1)))",
		"ST_Equals(\"geom\",'SRID=4326;POLYGON((0 0,0 9,9 0,0 0),(1 1,1 8,8 1,1 1))'::geometry)")
	checkCQL(t, "equals(geom, MULTIPOINT((0 0), (0 9)))",
		"ST_Equals(\"geom\",'SRID=4326;MULTIPOINT((0 0),(0 9))'::geometry)")
	checkCQL(t, "equals(geom, MULTILINESTRING((0 0, 1 1),(1 1, 2 2)))",
		"ST_Equals(\"geom\",'SRID=4326;MULTILINESTRING((0 0,1 1),(1 1,2 2))'::geometry)")
	checkCQL(t, "equals(geom, MULTIPOLYGON(((1 4, 4 1, 1 1, 1 4)), ((1 9, 4 9, 1 6, 1 9))))",
		"ST_Equals(\"geom\",'SRID=4326;MULTIPOLYGON(((1 4,4 1,1 1,1 4)),((1 9,4 9,1 6,1 9)))'::geometry)")
	checkCQL(t, "equals(geom, GEOMETRYCOLLECTION(POLYGON((1 4, 4 1, 1 1, 1 4)),LINESTRING (3 3, 5 5), POINT (1 5)))",
		"ST_Equals(\"geom\",'SRID=4326;GEOMETRYCOLLECTION(POLYGON((1 4,4 1,1 1,1 4)),LINESTRING(3 3,5 5),POINT(1 5))'::geometry)")
	checkCQL(t, "equals(geom, ENVELOPE(1,2,3,4))",
		"ST_Equals(\"geom\",ST_MakeEnvelope(1,2,3,4,4326))")
}

func TestGeometryLiteralWithSRID(t *testing.T) {
	checkCQLWithSRID(t, "equals(geom, POINT(0 0))", 1111, 2222,
		"ST_Equals(\"geom\",ST_Transform('SRID=1111;POINT(0 0)'::geometry,2222))")
	checkCQLWithSRID(t, "equals(geom, ENVELOPE(1,2,3,4))", 1111, 2222,
		"ST_Equals(\"geom\",ST_Transform(ST_MakeEnvelope(1,2,3,4,1111),2222))")
}

func TestBooleanExpression(t *testing.T) {
	checkCQL(t, "x > 1 AND x < 9", "\"x\" > 1 AND \"x\" < 9")
	checkCQL(t, "x = 1 OR x = 2", "\"x\" = 1 OR \"x\" = 2")
	checkCQL(t, "(x = 1 OR x = 2) AND y < 4", "(\"x\" = 1 OR \"x\" = 2) AND \"y\" < 4")
	checkCQL(t, "NOT x IS NOT NULL", "NOT  \"x\" IS NOT NULL")
}

func TestTemporal(t *testing.T) {
	checkCQL(t, "p BETWEEN 1991-01-01 AND 2000-12-31T01:59:59",
		"\"p\" BETWEEN timestamp '1991-01-01' AND timestamp '2000-12-31T01:59:59'")
	checkCQL(t, "1991-01-01 > p", "timestamp '1991-01-01' > \"p\"")
	checkCQL(t, "p > 1991-01-01T01:23:45.678", "\"p\" > timestamp '1991-01-01T01:23:45.678'")
	checkCQL(t, "p > 1991-01-01T01:23", "\"p\" > timestamp '1991-01-01T01:23'")
	checkCQL(t, "p > NOW()", "\"p\" > timestamp 'NOW'")
}

func TestSyntaxErrors(t *testing.T) {
	checkCQLError(t, "x y")
	checkCQLError(t, "x == y")
	checkCQLError(t, "x > 10y")
	checkCQLError(t, "NOT x IS > 3")
	// extra paren
	checkCQLError(t, "equals(geom, ENVELOPE(1,2,3,4)))")
	// comma between ordinates
	checkCQLError(t, "equals(geom, POINT(0,0))")
	// bad temporal values
	checkCQLError(t, "p > 200-01")
	checkCQLError(t, "p > 2000-01")
	checkCQLError(t, "p > 2000-01-01T01")
}

func checkCQL(t *testing.T, cqlStr string, sql string) {
	actual, err := TranspileToSQL(cqlStr, 4326, 4326)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	actual = strings.TrimSpace(actual)
	equals(t, sql, actual, "")
}

func checkCQLWithSRID(t *testing.T, cqlStr string, filterSRID int, sourceSRID int, sql string) {
	actual, err := TranspileToSQL(cqlStr, filterSRID, sourceSRID)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	actual = strings.TrimSpace(actual)
	equals(t, sql, actual, "")
}

func checkCQLError(t *testing.T, cqlStr string) {
	_, err := TranspileToSQL(cqlStr, 4326, 4326)
	isError(t, err, "")
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected: %#v; got: %#v\n", filepath.Base(file), line, msg, exp, act)
		tb.FailNow()
	}
}

// isError fails the test if err is nil
func isError(tb testing.TB, err error, msg string) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s - expected error\n", filepath.Base(file), line, msg)
		tb.FailNow()
	}
}
