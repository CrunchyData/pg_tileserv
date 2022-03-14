// Generated from CQLParser.g4 by ANTLR 4.7.

package cql // CQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCQLParserListener is a complete listener for a parse tree produced by CQLParser.
type BaseCQLParserListener struct{}

var _ CQLParserListener = &BaseCQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCqlFilter is called when production cqlFilter is entered.
func (s *BaseCQLParserListener) EnterCqlFilter(ctx *CqlFilterContext) {}

// ExitCqlFilter is called when production cqlFilter is exited.
func (s *BaseCQLParserListener) ExitCqlFilter(ctx *CqlFilterContext) {}

// EnterBooleanValueExpression is called when production booleanValueExpression is entered.
func (s *BaseCQLParserListener) EnterBooleanValueExpression(ctx *BooleanValueExpressionContext) {}

// ExitBooleanValueExpression is called when production booleanValueExpression is exited.
func (s *BaseCQLParserListener) ExitBooleanValueExpression(ctx *BooleanValueExpressionContext) {}

// EnterBooleanTerm is called when production booleanTerm is entered.
func (s *BaseCQLParserListener) EnterBooleanTerm(ctx *BooleanTermContext) {}

// ExitBooleanTerm is called when production booleanTerm is exited.
func (s *BaseCQLParserListener) ExitBooleanTerm(ctx *BooleanTermContext) {}

// EnterBooleanFactor is called when production booleanFactor is entered.
func (s *BaseCQLParserListener) EnterBooleanFactor(ctx *BooleanFactorContext) {}

// ExitBooleanFactor is called when production booleanFactor is exited.
func (s *BaseCQLParserListener) ExitBooleanFactor(ctx *BooleanFactorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseCQLParserListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseCQLParserListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseCQLParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseCQLParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBinaryComparisonPredicate is called when production binaryComparisonPredicate is entered.
func (s *BaseCQLParserListener) EnterBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// ExitBinaryComparisonPredicate is called when production binaryComparisonPredicate is exited.
func (s *BaseCQLParserListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseCQLParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseCQLParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseCQLParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseCQLParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseCQLParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseCQLParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseCQLParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseCQLParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterScalarExpression is called when production scalarExpression is entered.
func (s *BaseCQLParserListener) EnterScalarExpression(ctx *ScalarExpressionContext) {}

// ExitScalarExpression is called when production scalarExpression is exited.
func (s *BaseCQLParserListener) ExitScalarExpression(ctx *ScalarExpressionContext) {}

// EnterScalarValue is called when production scalarValue is entered.
func (s *BaseCQLParserListener) EnterScalarValue(ctx *ScalarValueContext) {}

// ExitScalarValue is called when production scalarValue is exited.
func (s *BaseCQLParserListener) ExitScalarValue(ctx *ScalarValueContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseCQLParserListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseCQLParserListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterCharacterLiteral is called when production characterLiteral is entered.
func (s *BaseCQLParserListener) EnterCharacterLiteral(ctx *CharacterLiteralContext) {}

// ExitCharacterLiteral is called when production characterLiteral is exited.
func (s *BaseCQLParserListener) ExitCharacterLiteral(ctx *CharacterLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseCQLParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseCQLParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCQLParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCQLParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterTemporalLiteral is called when production temporalLiteral is entered.
func (s *BaseCQLParserListener) EnterTemporalLiteral(ctx *TemporalLiteralContext) {}

// ExitTemporalLiteral is called when production temporalLiteral is exited.
func (s *BaseCQLParserListener) ExitTemporalLiteral(ctx *TemporalLiteralContext) {}

// EnterSpatialPredicate is called when production spatialPredicate is entered.
func (s *BaseCQLParserListener) EnterSpatialPredicate(ctx *SpatialPredicateContext) {}

// ExitSpatialPredicate is called when production spatialPredicate is exited.
func (s *BaseCQLParserListener) ExitSpatialPredicate(ctx *SpatialPredicateContext) {}

// EnterDistancePredicate is called when production distancePredicate is entered.
func (s *BaseCQLParserListener) EnterDistancePredicate(ctx *DistancePredicateContext) {}

// ExitDistancePredicate is called when production distancePredicate is exited.
func (s *BaseCQLParserListener) ExitDistancePredicate(ctx *DistancePredicateContext) {}

// EnterGeomExpression is called when production geomExpression is entered.
func (s *BaseCQLParserListener) EnterGeomExpression(ctx *GeomExpressionContext) {}

// ExitGeomExpression is called when production geomExpression is exited.
func (s *BaseCQLParserListener) ExitGeomExpression(ctx *GeomExpressionContext) {}

// EnterGeomLiteral is called when production geomLiteral is entered.
func (s *BaseCQLParserListener) EnterGeomLiteral(ctx *GeomLiteralContext) {}

// ExitGeomLiteral is called when production geomLiteral is exited.
func (s *BaseCQLParserListener) ExitGeomLiteral(ctx *GeomLiteralContext) {}

// EnterPoint is called when production point is entered.
func (s *BaseCQLParserListener) EnterPoint(ctx *PointContext) {}

// ExitPoint is called when production point is exited.
func (s *BaseCQLParserListener) ExitPoint(ctx *PointContext) {}

// EnterPointList is called when production pointList is entered.
func (s *BaseCQLParserListener) EnterPointList(ctx *PointListContext) {}

// ExitPointList is called when production pointList is exited.
func (s *BaseCQLParserListener) ExitPointList(ctx *PointListContext) {}

// EnterLinestring is called when production linestring is entered.
func (s *BaseCQLParserListener) EnterLinestring(ctx *LinestringContext) {}

// ExitLinestring is called when production linestring is exited.
func (s *BaseCQLParserListener) ExitLinestring(ctx *LinestringContext) {}

// EnterPolygon is called when production polygon is entered.
func (s *BaseCQLParserListener) EnterPolygon(ctx *PolygonContext) {}

// ExitPolygon is called when production polygon is exited.
func (s *BaseCQLParserListener) ExitPolygon(ctx *PolygonContext) {}

// EnterPolygonDef is called when production polygonDef is entered.
func (s *BaseCQLParserListener) EnterPolygonDef(ctx *PolygonDefContext) {}

// ExitPolygonDef is called when production polygonDef is exited.
func (s *BaseCQLParserListener) ExitPolygonDef(ctx *PolygonDefContext) {}

// EnterMultiPoint is called when production multiPoint is entered.
func (s *BaseCQLParserListener) EnterMultiPoint(ctx *MultiPointContext) {}

// ExitMultiPoint is called when production multiPoint is exited.
func (s *BaseCQLParserListener) ExitMultiPoint(ctx *MultiPointContext) {}

// EnterMultiLinestring is called when production multiLinestring is entered.
func (s *BaseCQLParserListener) EnterMultiLinestring(ctx *MultiLinestringContext) {}

// ExitMultiLinestring is called when production multiLinestring is exited.
func (s *BaseCQLParserListener) ExitMultiLinestring(ctx *MultiLinestringContext) {}

// EnterMultiPolygon is called when production multiPolygon is entered.
func (s *BaseCQLParserListener) EnterMultiPolygon(ctx *MultiPolygonContext) {}

// ExitMultiPolygon is called when production multiPolygon is exited.
func (s *BaseCQLParserListener) ExitMultiPolygon(ctx *MultiPolygonContext) {}

// EnterGeometryCollection is called when production geometryCollection is entered.
func (s *BaseCQLParserListener) EnterGeometryCollection(ctx *GeometryCollectionContext) {}

// ExitGeometryCollection is called when production geometryCollection is exited.
func (s *BaseCQLParserListener) ExitGeometryCollection(ctx *GeometryCollectionContext) {}

// EnterEnvelope is called when production envelope is entered.
func (s *BaseCQLParserListener) EnterEnvelope(ctx *EnvelopeContext) {}

// ExitEnvelope is called when production envelope is exited.
func (s *BaseCQLParserListener) ExitEnvelope(ctx *EnvelopeContext) {}

// EnterCoordList is called when production coordList is entered.
func (s *BaseCQLParserListener) EnterCoordList(ctx *CoordListContext) {}

// ExitCoordList is called when production coordList is exited.
func (s *BaseCQLParserListener) ExitCoordList(ctx *CoordListContext) {}

// EnterCoordinate is called when production coordinate is entered.
func (s *BaseCQLParserListener) EnterCoordinate(ctx *CoordinateContext) {}

// ExitCoordinate is called when production coordinate is exited.
func (s *BaseCQLParserListener) ExitCoordinate(ctx *CoordinateContext) {}
