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
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
)

func TranspileToSQL(cqlStr string, filterSRID int, sourceSRID int) (string, error) {
	if len(cqlStr) < 1 {
		return "", nil
	}
	// Setup the input
	is := antlr.NewInputStream(cqlStr)

	parseErrors := &CqlErrorListener{}

	// Create the Lexer
	lexer := NewCqlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(parseErrors)

	//-- create parser
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewCQLParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parseErrors)

	tree := parser.CqlFilter()
	//-- parse the CQL expression
	listener := NewCqlListener(filterSRID, sourceSRID)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if parseErrors.errorCount > 0 {
		log.Debug("CQL parser error = " + parseErrors.msg)
		msg := syntaxErrorMsg(cqlStr, parseErrors.col)
		err := fmt.Errorf("CQL syntax error: %s", msg)
		return "", err
	}
	return listener.GetSQL(), nil
}

func syntaxErrorMsg(input string, col int) string {
	start := 0
	dots1 := ""
	if col > 30 {
		start = col - 20
		dots1 = "..."
	}

	end := len(input)
	dots2 := ""
	if col < end-30 {
		end = col + 20
		dots2 = "..."
	}

	//-- extract parts and insert error flag
	before := input[start:col]
	after := input[col:end]
	msg := "\"" + dots1 + before + " !!>> " + after + dots2 + "\""
	return msg
}

//======================================
type CqlErrorListener struct {
	*antlr.DefaultErrorListener
	errorCount int
	line       int
	col        int
	msg        string
}

func (l *CqlErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//-- only report first error (but entire input is scanned by parser)
	if l.errorCount >= 1 {
		return
	}
	l.errorCount += 1
	l.line = line
	l.col = column
	l.msg = msg
}

func (l *CqlErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}
func (l *CqlErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}
func (l *CqlErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}

//----------------------------------

type cqlListener struct {
	*BaseCQLParserListener
	// SRID for filter CRS
	filterSRID int
	// SRID for source CRS
	sourceSRID int

	// final result SQL
	sql string
}

func NewCqlListener(filterSRID int, sourceSRID int) *cqlListener {
	this := new(cqlListener)
	this.filterSRID = filterSRID
	this.sourceSRID = sourceSRID
	return this
}
func (l *cqlListener) GetSQL() string {
	return l.sql
}

func (l *cqlListener) sqlGeometryLiteral(wkt string) string {
	sql := fmt.Sprintf("'SRID=%d;%s'::geometry", l.filterSRID, wkt)
	return sql
}

func (l *cqlListener) sqlEnvelopeLiteral(xmin string, ymin string, xmax string, ymax string) string {
	return fmt.Sprintf("ST_MakeEnvelope(%s,%s,%s,%s,%d)", xmin, ymin, xmax, ymax, l.filterSRID)
}

func (l *cqlListener) sqlTransformCrs(sql string) string {
	if l.sourceSRID == l.filterSRID {
		return sql
	}
	return fmt.Sprintf("ST_Transform(%s,%d)", sql, l.sourceSRID)
}

// helper function to avoid nil pointer problems
func getText(ctx antlr.ParserRuleContext) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetText()
}

// helper function to avoid nil pointer problems
func getNodeText(node antlr.TerminalNode) string {
	if node == nil {
		return ""
	}
	return node.GetText()
}

//========================================
type CqlContext struct {
	*antlr.BaseParserRuleContext
	// SQL fragment for the context subtree
	sql string
}

type SqlHolder interface {
	GetSql() string
}

func NewCqlContext(parent antlr.ParserRuleContext, invokingStateNumber int) *CqlContext {
	var p = new(CqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingStateNumber)
	return p
}

func (c *CqlContext) SetSql(sql string) {
	c.sql = sql
	//	fmt.Printf("setting sql for %s = %s\n", c.GetText(), sql)
}

func (c *CqlContext) GetSql() string {
	return c.sql
}

// helper function to get SQL from any context
func sqlFor(ctx antlr.ParserRuleContext) string {
	cc, ok := ctx.(SqlHolder)
	if ok {
		return cc.GetSql()
	}
	return "!ERROR!"
}

//========================================

/*
func (l *cqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}
*/

func (l *cqlListener) ExitCqlFilter(ctx *CqlFilterContext) {
	l.sql = sqlFor(ctx.BooleanValueExpression())
}

func (l *cqlListener) ExitBooleanValueExpression(ctx *BooleanValueExpressionContext) {
	sql := sqlFor(ctx.BooleanTerm())
	if ctx.OR() != nil {
		expr := sqlFor(ctx.BooleanValueExpression())
		sql = expr + " OR " + sql
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitBooleanTerm(ctx *BooleanTermContext) {
	sql := sqlFor(ctx.BooleanFactor())
	if ctx.AND() != nil {
		expr := sqlFor(ctx.BooleanTerm())
		sql = expr + " AND " + sql
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitBooleanFactor(ctx *BooleanFactorContext) {
	sql := sqlFor(ctx.BooleanPrimary())
	if ctx.NOT() != nil {
		sql = " NOT " + sql
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {
	var sql string
	if ctx.LEFTPAREN() == nil {
		sql = sqlFor(ctx.Predicate())
	} else {
		sql = "(" + sqlFor(ctx.BooleanValueExpression()) + ")"
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitPredicate(ctx *PredicateContext) {
	var sql string
	if ctx.BinaryComparisonPredicate() != nil {
		sql = sqlFor(ctx.BinaryComparisonPredicate())
	} else if ctx.LikePredicate() != nil {
		sql = sqlFor(ctx.LikePredicate())
	} else if ctx.BetweenPredicate() != nil {
		sql = sqlFor(ctx.BetweenPredicate())
	} else if ctx.IsNullPredicate() != nil {
		sql = sqlFor(ctx.IsNullPredicate())
	} else if ctx.InPredicate() != nil {
		sql = sqlFor(ctx.InPredicate())
	} else if ctx.SpatialPredicate() != nil {
		sql = sqlFor(ctx.SpatialPredicate())
	} else if ctx.DistancePredicate() != nil {
		sql = sqlFor(ctx.DistancePredicate())
		//	} else if ctx.TemporalPredicate() != nil {
		//		sql = sqlFor(ctx.TemporalPredicate())
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
	expr1 := sqlFor(ctx.ScalarExpression(0))
	expr2 := sqlFor(ctx.ScalarExpression(1))
	op := getNodeText(ctx.ComparisonOperator())
	sql := expr1 + " " + op + " " + expr2
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitScalarValue(ctx *ScalarValueContext) {
	var sql string
	if ctx.PropertyName() != nil {
		sql = quotedName(getText(ctx.PropertyName()))
	} else if ctx.CharacterLiteral() != nil {
		sql = quotedText(getText(ctx.CharacterLiteral()))
	} else if ctx.NumericLiteral() != nil {
		sql = getText(ctx.NumericLiteral())
	} else if ctx.BooleanLiteral() != nil {
		sql = getText(ctx.BooleanLiteral())
	} else if ctx.TemporalLiteral() != nil {
		sql = sqlFor(ctx.TemporalLiteral())
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitScalarExpression(ctx *ScalarExpressionContext) {
	var sql string
	if ctx.LEFTPAREN() != nil {
		sql = "(" + sqlFor(ctx.ScalarExpression(0)) + ")"
	} else if ctx.ArithmeticOperator() != nil {
		expr1 := sqlFor(ctx.ScalarExpression(0))
		expr2 := sqlFor(ctx.ScalarExpression(1))
		op := getNodeText(ctx.ArithmeticOperator())
		sql = expr1 + " " + op + " " + expr2
	} else {
		sql = sqlFor(ctx.ScalarValue())
	}
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitLikePredicate(ctx *LikePredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(getText(ctx.PropertyName())))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	op := " LIKE "
	if ctx.ILIKE() != nil {
		op = " ILIKE "
	}
	sb.WriteString(op)
	sb.WriteString(quotedText(getText(ctx.CharacterLiteral())))
	ctx.SetSql(sb.String())
}

func (l *cqlListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {
	prop := quotedName(getText(ctx.PropertyName()))
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	var expr1 string
	if ctx.ScalarExpression(0) != nil {
		expr1 = sqlFor(ctx.ScalarExpression(0))
	}
	/*
		if ctx.TemporalExpression(0) != nil {
			expr1 = sqlFor(ctx.TemporalExpression(0))
		}
	*/
	var expr2 string
	if ctx.ScalarExpression(0) != nil {
		expr2 = sqlFor(ctx.ScalarExpression(1))
	}
	/*
		if ctx.TemporalExpression(0) != nil {
			expr2 = sqlFor(ctx.TemporalExpression(1))
		}
	*/
	sql := " " + prop + not + " BETWEEN " + expr1 + " AND " + expr2
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {
	prop := quotedName(getText(ctx.PropertyName()))
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	sql := " " + prop + " IS" + not + " NULL"
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitInPredicate(ctx *InPredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(getText(ctx.PropertyName())))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	sb.WriteString(" IN (")
	inPredValueList(ctx, &sb)
	sb.WriteString(") ")
	sql := sb.String()
	ctx.SetSql(sql)
}

func inPredValueList(ctx *InPredicateContext, sb *strings.Builder) {
	//-- numeric literal list?
	nums := ctx.AllNumericLiteral()
	if len(nums) > 0 {
		for i, num := range nums {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(num.GetText())
		}
		return
	}
	//-- must be string literal list
	strs := ctx.AllCharacterLiteral()
	if len(strs) > 0 {
		for i, s := range strs {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(quotedText(s.GetText()))
		}
	}
}

func (l *cqlListener) ExitSpatialPredicate(ctx *SpatialPredicateContext) {
	var sb strings.Builder
	sb.WriteString(toPostGISFunction(ctx.SpatialOperator().GetText()))
	sb.WriteString("(")
	sb.WriteString(sqlFor(ctx.GeomExpression(0)))
	sb.WriteString(",")
	sb.WriteString(sqlFor(ctx.GeomExpression(1)))
	sb.WriteString(")")
	ctx.SetSql(sb.String())
}

func (l *cqlListener) ExitDistancePredicate(ctx *DistancePredicateContext) {
	var sb strings.Builder
	sb.WriteString(toPostGISFunction(ctx.DistanceOperator().GetText()))
	sb.WriteString("(")
	sb.WriteString(sqlFor(ctx.GeomExpression(0)))
	sb.WriteString(",")
	sb.WriteString(sqlFor(ctx.GeomExpression(1)))
	sb.WriteString(",")
	sb.WriteString(ctx.NumericLiteral().GetText())
	sb.WriteString(")")
	ctx.SetSql(sb.String())
}

func (l *cqlListener) ExitGeomExpression(ctx *GeomExpressionContext) {
	var sb strings.Builder
	if ctx.PropertyName() != nil {
		sb.WriteString(quotedName(getText(ctx.PropertyName())))
	} else {
		sb.WriteString(sqlFor(ctx.GeomLiteral()))
	}
	ctx.SetSql(sb.String())
}

func (l *cqlListener) ExitTemporalLiteral(ctx *TemporalLiteralContext) {
	val := strings.ToUpper(ctx.GetText())
	//var sql string
	if strings.HasPrefix(val, "NOW") {
		val = "NOW"
	}
	sql := fmt.Sprintf("timestamp '%s'", val)
	//TODO: handle NOW()
	ctx.SetSql(sql)
}

func (l *cqlListener) ExitGeomLiteral(ctx *GeomLiteralContext) {
	envCtx, ok := ctx.GetChild(0).(*EnvelopeContext)
	var sql string
	if ok {
		nums := envCtx.AllNumericLiteral()
		b1 := nums[0].GetText()
		b2 := nums[1].GetText()
		b3 := nums[2].GetText()
		b4 := nums[3].GetText()
		sql = l.sqlEnvelopeLiteral(b1, b2, b3, b4)
	} else {
		wkt := getGeomText(ctx)
		sql = l.sqlGeometryLiteral(wkt)
	}
	sql = l.sqlTransformCrs(sql)
	ctx.SetSql(sql)
}

func getGeomText(ctx *GeomLiteralContext) string {
	trees := ctx.GetChildren()
	var sb strings.Builder
	extractGeomText(trees, &sb)
	return sb.String()
}

func extractGeomText(trees []antlr.Tree, sb *strings.Builder) {
	isPrevNumeric := false
	for _, t := range trees {
		tn, ok := t.(antlr.TerminalNode)
		if ok {
			//-- add a blank between consecutive numbers to separate them
			if tn.GetSymbol().GetTokenType() == CQLParserNumericLiteral {
				if isPrevNumeric {
					sb.WriteString(" ")
				}
				isPrevNumeric = true
			} else {
				isPrevNumeric = false
			}
			sb.WriteString(strings.ToUpper(tn.GetText()))
		} else {
			ch := t.GetChildren()
			extractGeomText(ch, sb)
		}
	}
}

var pgFunctionForCql = map[string]string{
	"crosses":    "ST_Crosses",
	"contains":   "ST_Contains",
	"disjoint":   "ST_Disjoint",
	"equals":     "ST_Equals",
	"intersects": "ST_Intersects",
	"overlaps":   "ST_Overlaps",
	"touches":    "ST_Touches",
	"within":     "ST_Within",

	"dwithin": "ST_DWithin",
}

func toPostGISFunction(cqlFunName string) string {
	cqlNameLow := strings.ToLower(cqlFunName)
	if fun, ok := pgFunctionForCql[cqlNameLow]; ok {
		return fun
	}
	//-- this will trigger a SQL error
	return "UNKNOWN_" + cqlFunName
}

func quotedName(name string) string {
	//-- CQL property names can be quoted
	if strings.HasPrefix(name, "\"") {
		return name
	}
	return "\"" + name + "\""
}

func quotedText(s string) string {
	//TODO: make this better (escape quotes, etc)
	//TODO: is SQL injection a risk here?
	return s
}
