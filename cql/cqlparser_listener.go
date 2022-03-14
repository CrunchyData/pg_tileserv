// Generated from CQLParser.g4 by ANTLR 4.7.

package cql // CQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CQLParserListener is a complete listener for a parse tree produced by CQLParser.
type CQLParserListener interface {
	antlr.ParseTreeListener

	// EnterCqlFilter is called when entering the cqlFilter production.
	EnterCqlFilter(c *CqlFilterContext)

	// EnterBooleanValueExpression is called when entering the booleanValueExpression production.
	EnterBooleanValueExpression(c *BooleanValueExpressionContext)

	// EnterBooleanTerm is called when entering the booleanTerm production.
	EnterBooleanTerm(c *BooleanTermContext)

	// EnterBooleanFactor is called when entering the booleanFactor production.
	EnterBooleanFactor(c *BooleanFactorContext)

	// EnterBooleanPrimary is called when entering the booleanPrimary production.
	EnterBooleanPrimary(c *BooleanPrimaryContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterBinaryComparisonPredicate is called when entering the binaryComparisonPredicate production.
	EnterBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterScalarExpression is called when entering the scalarExpression production.
	EnterScalarExpression(c *ScalarExpressionContext)

	// EnterScalarValue is called when entering the scalarValue production.
	EnterScalarValue(c *ScalarValueContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterCharacterLiteral is called when entering the characterLiteral production.
	EnterCharacterLiteral(c *CharacterLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterTemporalLiteral is called when entering the temporalLiteral production.
	EnterTemporalLiteral(c *TemporalLiteralContext)

	// EnterSpatialPredicate is called when entering the spatialPredicate production.
	EnterSpatialPredicate(c *SpatialPredicateContext)

	// EnterDistancePredicate is called when entering the distancePredicate production.
	EnterDistancePredicate(c *DistancePredicateContext)

	// EnterGeomExpression is called when entering the geomExpression production.
	EnterGeomExpression(c *GeomExpressionContext)

	// EnterGeomLiteral is called when entering the geomLiteral production.
	EnterGeomLiteral(c *GeomLiteralContext)

	// EnterPoint is called when entering the point production.
	EnterPoint(c *PointContext)

	// EnterPointList is called when entering the pointList production.
	EnterPointList(c *PointListContext)

	// EnterLinestring is called when entering the linestring production.
	EnterLinestring(c *LinestringContext)

	// EnterPolygon is called when entering the polygon production.
	EnterPolygon(c *PolygonContext)

	// EnterPolygonDef is called when entering the polygonDef production.
	EnterPolygonDef(c *PolygonDefContext)

	// EnterMultiPoint is called when entering the multiPoint production.
	EnterMultiPoint(c *MultiPointContext)

	// EnterMultiLinestring is called when entering the multiLinestring production.
	EnterMultiLinestring(c *MultiLinestringContext)

	// EnterMultiPolygon is called when entering the multiPolygon production.
	EnterMultiPolygon(c *MultiPolygonContext)

	// EnterGeometryCollection is called when entering the geometryCollection production.
	EnterGeometryCollection(c *GeometryCollectionContext)

	// EnterEnvelope is called when entering the envelope production.
	EnterEnvelope(c *EnvelopeContext)

	// EnterCoordList is called when entering the coordList production.
	EnterCoordList(c *CoordListContext)

	// EnterCoordinate is called when entering the coordinate production.
	EnterCoordinate(c *CoordinateContext)

	// ExitCqlFilter is called when exiting the cqlFilter production.
	ExitCqlFilter(c *CqlFilterContext)

	// ExitBooleanValueExpression is called when exiting the booleanValueExpression production.
	ExitBooleanValueExpression(c *BooleanValueExpressionContext)

	// ExitBooleanTerm is called when exiting the booleanTerm production.
	ExitBooleanTerm(c *BooleanTermContext)

	// ExitBooleanFactor is called when exiting the booleanFactor production.
	ExitBooleanFactor(c *BooleanFactorContext)

	// ExitBooleanPrimary is called when exiting the booleanPrimary production.
	ExitBooleanPrimary(c *BooleanPrimaryContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitBinaryComparisonPredicate is called when exiting the binaryComparisonPredicate production.
	ExitBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitScalarExpression is called when exiting the scalarExpression production.
	ExitScalarExpression(c *ScalarExpressionContext)

	// ExitScalarValue is called when exiting the scalarValue production.
	ExitScalarValue(c *ScalarValueContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitCharacterLiteral is called when exiting the characterLiteral production.
	ExitCharacterLiteral(c *CharacterLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitTemporalLiteral is called when exiting the temporalLiteral production.
	ExitTemporalLiteral(c *TemporalLiteralContext)

	// ExitSpatialPredicate is called when exiting the spatialPredicate production.
	ExitSpatialPredicate(c *SpatialPredicateContext)

	// ExitDistancePredicate is called when exiting the distancePredicate production.
	ExitDistancePredicate(c *DistancePredicateContext)

	// ExitGeomExpression is called when exiting the geomExpression production.
	ExitGeomExpression(c *GeomExpressionContext)

	// ExitGeomLiteral is called when exiting the geomLiteral production.
	ExitGeomLiteral(c *GeomLiteralContext)

	// ExitPoint is called when exiting the point production.
	ExitPoint(c *PointContext)

	// ExitPointList is called when exiting the pointList production.
	ExitPointList(c *PointListContext)

	// ExitLinestring is called when exiting the linestring production.
	ExitLinestring(c *LinestringContext)

	// ExitPolygon is called when exiting the polygon production.
	ExitPolygon(c *PolygonContext)

	// ExitPolygonDef is called when exiting the polygonDef production.
	ExitPolygonDef(c *PolygonDefContext)

	// ExitMultiPoint is called when exiting the multiPoint production.
	ExitMultiPoint(c *MultiPointContext)

	// ExitMultiLinestring is called when exiting the multiLinestring production.
	ExitMultiLinestring(c *MultiLinestringContext)

	// ExitMultiPolygon is called when exiting the multiPolygon production.
	ExitMultiPolygon(c *MultiPolygonContext)

	// ExitGeometryCollection is called when exiting the geometryCollection production.
	ExitGeometryCollection(c *GeometryCollectionContext)

	// ExitEnvelope is called when exiting the envelope production.
	ExitEnvelope(c *EnvelopeContext)

	// ExitCoordList is called when exiting the coordList production.
	ExitCoordList(c *CoordListContext)

	// ExitCoordinate is called when exiting the coordinate production.
	ExitCoordinate(c *CoordinateContext)
}
