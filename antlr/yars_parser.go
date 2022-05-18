// Code generated from antlr/Yars.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Yars
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type YarsParser struct {
	*antlr.BaseParser
}

var yarsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yarsParserInit() {
	staticData := &yarsParserStaticData
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
		"project", "actionStatement", "recommendStatement", "paramsDecl", "paramDecl",
		"expression", "lvalue", "binaryOps", "unaryOps", "unitStatement", "fieldsDecl",
		"fieldDecl", "type_", "arrayType", "baseType", "eos",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 135, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 39, 8, 0, 10, 0, 12, 0, 42, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 64, 8, 3, 10,
		3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 5, 5, 90, 8, 5, 10, 5, 12, 5, 93, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 99, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 5, 10, 113, 8, 10, 10, 10, 12, 10, 116, 9, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 125, 8, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 1, 10, 16, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3, 1, 0, 7, 12,
		2, 0, 6, 6, 10, 11, 1, 0, 1, 5, 128, 0, 32, 1, 0, 0, 0, 2, 45, 1, 0, 0,
		0, 4, 52, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 83, 1,
		0, 0, 0, 12, 98, 1, 0, 0, 0, 14, 100, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0,
		18, 104, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 124,
		1, 0, 0, 0, 26, 126, 1, 0, 0, 0, 28, 130, 1, 0, 0, 0, 30, 132, 1, 0, 0,
		0, 32, 33, 5, 15, 0, 0, 33, 34, 5, 17, 0, 0, 34, 40, 5, 24, 0, 0, 35, 39,
		3, 18, 9, 0, 36, 39, 3, 2, 1, 0, 37, 39, 3, 4, 2, 0, 38, 35, 1, 0, 0, 0,
		38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1,
		0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43,
		44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 46, 5, 14, 0, 0, 46, 47, 5, 17,
		0, 0, 47, 48, 5, 17, 0, 0, 48, 49, 5, 26, 0, 0, 49, 50, 5, 17, 0, 0, 50,
		51, 5, 24, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 5, 16, 0, 0, 53, 54, 5, 17,
		0, 0, 54, 55, 5, 17, 0, 0, 55, 56, 5, 26, 0, 0, 56, 57, 5, 17, 0, 0, 57,
		58, 3, 6, 3, 0, 58, 5, 1, 0, 0, 0, 59, 65, 5, 18, 0, 0, 60, 61, 3, 8, 4,
		0, 61, 62, 3, 30, 15, 0, 62, 64, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 64, 67,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0,
		67, 65, 1, 0, 0, 0, 68, 69, 5, 19, 0, 0, 69, 7, 1, 0, 0, 0, 70, 71, 5,
		17, 0, 0, 71, 72, 5, 25, 0, 0, 72, 73, 3, 10, 5, 0, 73, 9, 1, 0, 0, 0,
		74, 75, 6, 5, -1, 0, 75, 76, 5, 22, 0, 0, 76, 77, 3, 10, 5, 0, 77, 78,
		5, 23, 0, 0, 78, 84, 1, 0, 0, 0, 79, 80, 3, 16, 8, 0, 80, 81, 3, 10, 5,
		3, 81, 84, 1, 0, 0, 0, 82, 84, 3, 12, 6, 0, 83, 74, 1, 0, 0, 0, 83, 79,
		1, 0, 0, 0, 83, 82, 1, 0, 0, 0, 84, 91, 1, 0, 0, 0, 85, 86, 10, 2, 0, 0,
		86, 87, 3, 14, 7, 0, 87, 88, 3, 10, 5, 3, 88, 90, 1, 0, 0, 0, 89, 85, 1,
		0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		11, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 99, 5, 17, 0, 0, 95, 96, 5, 17,
		0, 0, 96, 97, 5, 28, 0, 0, 97, 99, 5, 17, 0, 0, 98, 94, 1, 0, 0, 0, 98,
		95, 1, 0, 0, 0, 99, 13, 1, 0, 0, 0, 100, 101, 7, 0, 0, 0, 101, 15, 1, 0,
		0, 0, 102, 103, 7, 1, 0, 0, 103, 17, 1, 0, 0, 0, 104, 105, 5, 13, 0, 0,
		105, 106, 5, 17, 0, 0, 106, 107, 3, 20, 10, 0, 107, 19, 1, 0, 0, 0, 108,
		114, 5, 18, 0, 0, 109, 110, 3, 22, 11, 0, 110, 111, 3, 30, 15, 0, 111,
		113, 1, 0, 0, 0, 112, 109, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 117, 118, 5, 19, 0, 0, 118, 21, 1, 0, 0, 0, 119, 120, 5, 17, 0, 0,
		120, 121, 3, 24, 12, 0, 121, 23, 1, 0, 0, 0, 122, 125, 3, 26, 13, 0, 123,
		125, 3, 28, 14, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 25,
		1, 0, 0, 0, 126, 127, 5, 20, 0, 0, 127, 128, 5, 21, 0, 0, 128, 129, 3,
		28, 14, 0, 129, 27, 1, 0, 0, 0, 130, 131, 7, 2, 0, 0, 131, 29, 1, 0, 0,
		0, 132, 133, 5, 24, 0, 0, 133, 31, 1, 0, 0, 0, 8, 38, 40, 65, 83, 91, 98,
		114, 124,
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

// YarsParserInit initializes any static state used to implement YarsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYarsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YarsParserInit() {
	staticData := &yarsParserStaticData
	staticData.once.Do(yarsParserInit)
}

// NewYarsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYarsParser(input antlr.TokenStream) *YarsParser {
	YarsParserInit()
	this := new(YarsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &yarsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Yars.g4"

	return this
}

// YarsParser tokens.
const (
	YarsParserEOF        = antlr.TokenEOF
	YarsParserINT32      = 1
	YarsParserUINT32     = 2
	YarsParserFLOAT32    = 3
	YarsParserBOOL       = 4
	YarsParserSTRING     = 5
	YarsParserNOT        = 6
	YarsParserAND        = 7
	YarsParserOR         = 8
	YarsParserIN         = 9
	YarsParserPLUS       = 10
	YarsParserMINUS      = 11
	YarsParserMUL        = 12
	YarsParserUNIT       = 13
	YarsParserACTION     = 14
	YarsParserPROJECT    = 15
	YarsParserRECOMMEND  = 16
	YarsParserIDENTIFIER = 17
	YarsParserL_CURLY    = 18
	YarsParserR_CURLY    = 19
	YarsParserL_BRACKET  = 20
	YarsParserR_BRACKET  = 21
	YarsParserL_RBRACKET = 22
	YarsParserR_RBRACKET = 23
	YarsParserSEMI       = 24
	YarsParserCOLON      = 25
	YarsParserAT         = 26
	YarsParserCOMMA      = 27
	YarsParserDOT        = 28
	YarsParserWHITESPACE = 29
)

// YarsParser rules.
const (
	YarsParserRULE_project            = 0
	YarsParserRULE_actionStatement    = 1
	YarsParserRULE_recommendStatement = 2
	YarsParserRULE_paramsDecl         = 3
	YarsParserRULE_paramDecl          = 4
	YarsParserRULE_expression         = 5
	YarsParserRULE_lvalue             = 6
	YarsParserRULE_binaryOps          = 7
	YarsParserRULE_unaryOps           = 8
	YarsParserRULE_unitStatement      = 9
	YarsParserRULE_fieldsDecl         = 10
	YarsParserRULE_fieldDecl          = 11
	YarsParserRULE_type_              = 12
	YarsParserRULE_arrayType          = 13
	YarsParserRULE_baseType           = 14
	YarsParserRULE_eos                = 15
)

// IProjectContext is an interface to support dynamic dispatch.
type IProjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectContext differentiates from other interfaces.
	IsProjectContext()
}

type ProjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectContext() *ProjectContext {
	var p = new(ProjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_project
	return p
}

func (*ProjectContext) IsProjectContext() {}

func NewProjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectContext {
	var p = new(ProjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_project

	return p
}

func (s *ProjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(YarsParserPROJECT, 0)
}

func (s *ProjectContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, 0)
}

func (s *ProjectContext) SEMI() antlr.TerminalNode {
	return s.GetToken(YarsParserSEMI, 0)
}

func (s *ProjectContext) EOF() antlr.TerminalNode {
	return s.GetToken(YarsParserEOF, 0)
}

func (s *ProjectContext) AllUnitStatement() []IUnitStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitStatementContext); ok {
			len++
		}
	}

	tst := make([]IUnitStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitStatementContext); ok {
			tst[i] = t.(IUnitStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProjectContext) UnitStatement(i int) IUnitStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitStatementContext); ok {
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

	return t.(IUnitStatementContext)
}

func (s *ProjectContext) AllActionStatement() []IActionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionStatementContext); ok {
			len++
		}
	}

	tst := make([]IActionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionStatementContext); ok {
			tst[i] = t.(IActionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProjectContext) ActionStatement(i int) IActionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionStatementContext); ok {
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

	return t.(IActionStatementContext)
}

func (s *ProjectContext) AllRecommendStatement() []IRecommendStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecommendStatementContext); ok {
			len++
		}
	}

	tst := make([]IRecommendStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecommendStatementContext); ok {
			tst[i] = t.(IRecommendStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProjectContext) RecommendStatement(i int) IRecommendStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecommendStatementContext); ok {
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

	return t.(IRecommendStatementContext)
}

func (s *ProjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterProject(s)
	}
}

func (s *ProjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitProject(s)
	}
}

func (p *YarsParser) Project() (localctx IProjectContext) {
	this := p
	_ = this

	localctx = NewProjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YarsParserRULE_project)
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
		p.SetState(32)
		p.Match(YarsParserPROJECT)
	}
	{
		p.SetState(33)
		p.Match(YarsParserIDENTIFIER)
	}
	{
		p.SetState(34)
		p.Match(YarsParserSEMI)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YarsParserUNIT)|(1<<YarsParserACTION)|(1<<YarsParserRECOMMEND))) != 0 {
		p.SetState(38)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case YarsParserUNIT:
			{
				p.SetState(35)
				p.UnitStatement()
			}

		case YarsParserACTION:
			{
				p.SetState(36)
				p.ActionStatement()
			}

		case YarsParserRECOMMEND:
			{
				p.SetState(37)
				p.RecommendStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(YarsParserEOF)
	}

	return localctx
}

// IActionStatementContext is an interface to support dynamic dispatch.
type IActionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetActor returns the actor token.
	GetActor() antlr.Token

	// GetObject returns the object token.
	GetObject() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetActor sets the actor token.
	SetActor(antlr.Token)

	// SetObject sets the object token.
	SetObject(antlr.Token)

	// IsActionStatementContext differentiates from other interfaces.
	IsActionStatementContext()
}

type ActionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	actor  antlr.Token
	object antlr.Token
}

func NewEmptyActionStatementContext() *ActionStatementContext {
	var p = new(ActionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_actionStatement
	return p
}

func (*ActionStatementContext) IsActionStatementContext() {}

func NewActionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionStatementContext {
	var p = new(ActionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_actionStatement

	return p
}

func (s *ActionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionStatementContext) GetName() antlr.Token { return s.name }

func (s *ActionStatementContext) GetActor() antlr.Token { return s.actor }

func (s *ActionStatementContext) GetObject() antlr.Token { return s.object }

func (s *ActionStatementContext) SetName(v antlr.Token) { s.name = v }

func (s *ActionStatementContext) SetActor(v antlr.Token) { s.actor = v }

func (s *ActionStatementContext) SetObject(v antlr.Token) { s.object = v }

func (s *ActionStatementContext) ACTION() antlr.TerminalNode {
	return s.GetToken(YarsParserACTION, 0)
}

func (s *ActionStatementContext) AT() antlr.TerminalNode {
	return s.GetToken(YarsParserAT, 0)
}

func (s *ActionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(YarsParserSEMI, 0)
}

func (s *ActionStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(YarsParserIDENTIFIER)
}

func (s *ActionStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, i)
}

func (s *ActionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterActionStatement(s)
	}
}

func (s *ActionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitActionStatement(s)
	}
}

func (p *YarsParser) ActionStatement() (localctx IActionStatementContext) {
	this := p
	_ = this

	localctx = NewActionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YarsParserRULE_actionStatement)

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
		p.SetState(45)
		p.Match(YarsParserACTION)
	}
	{
		p.SetState(46)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*ActionStatementContext).name = _m
	}
	{
		p.SetState(47)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*ActionStatementContext).actor = _m
	}
	{
		p.SetState(48)
		p.Match(YarsParserAT)
	}
	{
		p.SetState(49)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*ActionStatementContext).object = _m
	}
	{
		p.SetState(50)
		p.Match(YarsParserSEMI)
	}

	return localctx
}

// IRecommendStatementContext is an interface to support dynamic dispatch.
type IRecommendStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetActor returns the actor token.
	GetActor() antlr.Token

	// GetObject returns the object token.
	GetObject() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetActor sets the actor token.
	SetActor(antlr.Token)

	// SetObject sets the object token.
	SetObject(antlr.Token)

	// IsRecommendStatementContext differentiates from other interfaces.
	IsRecommendStatementContext()
}

type RecommendStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	actor  antlr.Token
	object antlr.Token
}

func NewEmptyRecommendStatementContext() *RecommendStatementContext {
	var p = new(RecommendStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_recommendStatement
	return p
}

func (*RecommendStatementContext) IsRecommendStatementContext() {}

func NewRecommendStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecommendStatementContext {
	var p = new(RecommendStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_recommendStatement

	return p
}

func (s *RecommendStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RecommendStatementContext) GetName() antlr.Token { return s.name }

func (s *RecommendStatementContext) GetActor() antlr.Token { return s.actor }

func (s *RecommendStatementContext) GetObject() antlr.Token { return s.object }

func (s *RecommendStatementContext) SetName(v antlr.Token) { s.name = v }

func (s *RecommendStatementContext) SetActor(v antlr.Token) { s.actor = v }

func (s *RecommendStatementContext) SetObject(v antlr.Token) { s.object = v }

func (s *RecommendStatementContext) RECOMMEND() antlr.TerminalNode {
	return s.GetToken(YarsParserRECOMMEND, 0)
}

func (s *RecommendStatementContext) AT() antlr.TerminalNode {
	return s.GetToken(YarsParserAT, 0)
}

func (s *RecommendStatementContext) ParamsDecl() IParamsDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsDeclContext)
}

func (s *RecommendStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(YarsParserIDENTIFIER)
}

func (s *RecommendStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, i)
}

func (s *RecommendStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecommendStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecommendStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterRecommendStatement(s)
	}
}

func (s *RecommendStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitRecommendStatement(s)
	}
}

func (p *YarsParser) RecommendStatement() (localctx IRecommendStatementContext) {
	this := p
	_ = this

	localctx = NewRecommendStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YarsParserRULE_recommendStatement)

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
		p.SetState(52)
		p.Match(YarsParserRECOMMEND)
	}
	{
		p.SetState(53)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*RecommendStatementContext).name = _m
	}
	{
		p.SetState(54)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*RecommendStatementContext).actor = _m
	}
	{
		p.SetState(55)
		p.Match(YarsParserAT)
	}
	{
		p.SetState(56)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*RecommendStatementContext).object = _m
	}
	{
		p.SetState(57)
		p.ParamsDecl()
	}

	return localctx
}

// IParamsDeclContext is an interface to support dynamic dispatch.
type IParamsDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsDeclContext differentiates from other interfaces.
	IsParamsDeclContext()
}

type ParamsDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsDeclContext() *ParamsDeclContext {
	var p = new(ParamsDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_paramsDecl
	return p
}

func (*ParamsDeclContext) IsParamsDeclContext() {}

func NewParamsDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsDeclContext {
	var p = new(ParamsDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_paramsDecl

	return p
}

func (s *ParamsDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsDeclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(YarsParserL_CURLY, 0)
}

func (s *ParamsDeclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(YarsParserR_CURLY, 0)
}

func (s *ParamsDeclContext) AllParamDecl() []IParamDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclContext); ok {
			tst[i] = t.(IParamDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParamsDeclContext) ParamDecl(i int) IParamDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclContext); ok {
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

	return t.(IParamDeclContext)
}

func (s *ParamsDeclContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *ParamsDeclContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
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

	return t.(IEosContext)
}

func (s *ParamsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterParamsDecl(s)
	}
}

func (s *ParamsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitParamsDecl(s)
	}
}

func (p *YarsParser) ParamsDecl() (localctx IParamsDeclContext) {
	this := p
	_ = this

	localctx = NewParamsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YarsParserRULE_paramsDecl)
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
		p.SetState(59)
		p.Match(YarsParserL_CURLY)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == YarsParserIDENTIFIER {
		{
			p.SetState(60)
			p.ParamDecl()
		}
		{
			p.SetState(61)
			p.Eos()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(YarsParserR_CURLY)
	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) GetName() antlr.Token { return s.name }

func (s *ParamDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ParamDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(YarsParserCOLON, 0)
}

func (s *ParamDeclContext) Expression() IExpressionContext {
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

func (s *ParamDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, 0)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *YarsParser) ParamDecl() (localctx IParamDeclContext) {
	this := p
	_ = this

	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YarsParserRULE_paramDecl)

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
		p.SetState(70)

		var _m = p.Match(YarsParserIDENTIFIER)

		localctx.(*ParamDeclContext).name = _m
	}
	{
		p.SetState(71)
		p.Match(YarsParserCOLON)
	}
	{
		p.SetState(72)
		p.expression(0)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryOpContext struct {
	*ExpressionContext
	op IBinaryOpsContext
}

func NewBinaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryOpContext) GetOp() IBinaryOpsContext { return s.op }

func (s *BinaryOpContext) SetOp(v IBinaryOpsContext) { s.op = v }

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) AllExpression() []IExpressionContext {
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

func (s *BinaryOpContext) Expression(i int) IExpressionContext {
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

func (s *BinaryOpContext) BinaryOps() IBinaryOpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOpsContext)
}

func (s *BinaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterBinaryOp(s)
	}
}

func (s *BinaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitBinaryOp(s)
	}
}

type LvalueExpressionContext struct {
	*ExpressionContext
}

func NewLvalueExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LvalueExpressionContext {
	var p = new(LvalueExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LvalueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueExpressionContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LvalueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterLvalueExpression(s)
	}
}

func (s *LvalueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitLvalueExpression(s)
	}
}

type UnaryOpContext struct {
	*ExpressionContext
	op IUnaryOpsContext
}

func NewUnaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryOpContext) GetOp() IUnaryOpsContext { return s.op }

func (s *UnaryOpContext) SetOp(v IUnaryOpsContext) { s.op = v }

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) Expression() IExpressionContext {
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

func (s *UnaryOpContext) UnaryOps() IUnaryOpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpsContext)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

type BracketsContext struct {
	*ExpressionContext
}

func NewBracketsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketsContext {
	var p = new(BracketsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BracketsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketsContext) L_RBRACKET() antlr.TerminalNode {
	return s.GetToken(YarsParserL_RBRACKET, 0)
}

func (s *BracketsContext) Expression() IExpressionContext {
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

func (s *BracketsContext) R_RBRACKET() antlr.TerminalNode {
	return s.GetToken(YarsParserR_RBRACKET, 0)
}

func (s *BracketsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterBrackets(s)
	}
}

func (s *BracketsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitBrackets(s)
	}
}

func (p *YarsParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *YarsParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, YarsParserRULE_expression, _p)

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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YarsParserL_RBRACKET:
		localctx = NewBracketsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(75)
			p.Match(YarsParserL_RBRACKET)
		}
		{
			p.SetState(76)
			p.expression(0)
		}
		{
			p.SetState(77)
			p.Match(YarsParserR_RBRACKET)
		}

	case YarsParserNOT, YarsParserPLUS, YarsParserMINUS:
		localctx = NewUnaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)

			var _x = p.UnaryOps()

			localctx.(*UnaryOpContext).op = _x
		}
		{
			p.SetState(80)
			p.expression(3)
		}

	case YarsParserIDENTIFIER:
		localctx = NewLvalueExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Lvalue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBinaryOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, YarsParserRULE_expression)
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(86)

				var _x = p.BinaryOps()

				localctx.(*BinaryOpContext).op = _x
			}
			{
				p.SetState(87)
				p.expression(3)
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetField returns the field token.
	GetField() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetField sets the field token.
	SetField(antlr.Token)

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	field  antlr.Token
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) GetName() antlr.Token { return s.name }

func (s *LvalueContext) GetField() antlr.Token { return s.field }

func (s *LvalueContext) SetName(v antlr.Token) { s.name = v }

func (s *LvalueContext) SetField(v antlr.Token) { s.field = v }

func (s *LvalueContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(YarsParserIDENTIFIER)
}

func (s *LvalueContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, i)
}

func (s *LvalueContext) DOT() antlr.TerminalNode {
	return s.GetToken(YarsParserDOT, 0)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *YarsParser) Lvalue() (localctx ILvalueContext) {
	this := p
	_ = this

	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YarsParserRULE_lvalue)

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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)

			var _m = p.Match(YarsParserIDENTIFIER)

			localctx.(*LvalueContext).name = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)

			var _m = p.Match(YarsParserIDENTIFIER)

			localctx.(*LvalueContext).name = _m
		}
		{
			p.SetState(96)
			p.Match(YarsParserDOT)
		}
		{
			p.SetState(97)

			var _m = p.Match(YarsParserIDENTIFIER)

			localctx.(*LvalueContext).field = _m
		}

	}

	return localctx
}

// IBinaryOpsContext is an interface to support dynamic dispatch.
type IBinaryOpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOpsContext differentiates from other interfaces.
	IsBinaryOpsContext()
}

type BinaryOpsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpsContext() *BinaryOpsContext {
	var p = new(BinaryOpsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_binaryOps
	return p
}

func (*BinaryOpsContext) IsBinaryOpsContext() {}

func NewBinaryOpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpsContext {
	var p = new(BinaryOpsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_binaryOps

	return p
}

func (s *BinaryOpsContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpsContext) MUL() antlr.TerminalNode {
	return s.GetToken(YarsParserMUL, 0)
}

func (s *BinaryOpsContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YarsParserPLUS, 0)
}

func (s *BinaryOpsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YarsParserMINUS, 0)
}

func (s *BinaryOpsContext) AND() antlr.TerminalNode {
	return s.GetToken(YarsParserAND, 0)
}

func (s *BinaryOpsContext) OR() antlr.TerminalNode {
	return s.GetToken(YarsParserOR, 0)
}

func (s *BinaryOpsContext) IN() antlr.TerminalNode {
	return s.GetToken(YarsParserIN, 0)
}

func (s *BinaryOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterBinaryOps(s)
	}
}

func (s *BinaryOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitBinaryOps(s)
	}
}

func (p *YarsParser) BinaryOps() (localctx IBinaryOpsContext) {
	this := p
	_ = this

	localctx = NewBinaryOpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YarsParserRULE_binaryOps)
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
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YarsParserAND)|(1<<YarsParserOR)|(1<<YarsParserIN)|(1<<YarsParserPLUS)|(1<<YarsParserMINUS)|(1<<YarsParserMUL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpsContext is an interface to support dynamic dispatch.
type IUnaryOpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpsContext differentiates from other interfaces.
	IsUnaryOpsContext()
}

type UnaryOpsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpsContext() *UnaryOpsContext {
	var p = new(UnaryOpsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_unaryOps
	return p
}

func (*UnaryOpsContext) IsUnaryOpsContext() {}

func NewUnaryOpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpsContext {
	var p = new(UnaryOpsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_unaryOps

	return p
}

func (s *UnaryOpsContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpsContext) NOT() antlr.TerminalNode {
	return s.GetToken(YarsParserNOT, 0)
}

func (s *UnaryOpsContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YarsParserPLUS, 0)
}

func (s *UnaryOpsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YarsParserMINUS, 0)
}

func (s *UnaryOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterUnaryOps(s)
	}
}

func (s *UnaryOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitUnaryOps(s)
	}
}

func (p *YarsParser) UnaryOps() (localctx IUnaryOpsContext) {
	this := p
	_ = this

	localctx = NewUnaryOpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, YarsParserRULE_unaryOps)
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
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YarsParserNOT)|(1<<YarsParserPLUS)|(1<<YarsParserMINUS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnitStatementContext is an interface to support dynamic dispatch.
type IUnitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitStatementContext differentiates from other interfaces.
	IsUnitStatementContext()
}

type UnitStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitStatementContext() *UnitStatementContext {
	var p = new(UnitStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_unitStatement
	return p
}

func (*UnitStatementContext) IsUnitStatementContext() {}

func NewUnitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitStatementContext {
	var p = new(UnitStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_unitStatement

	return p
}

func (s *UnitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitStatementContext) UNIT() antlr.TerminalNode {
	return s.GetToken(YarsParserUNIT, 0)
}

func (s *UnitStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, 0)
}

func (s *UnitStatementContext) FieldsDecl() IFieldsDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsDeclContext)
}

func (s *UnitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterUnitStatement(s)
	}
}

func (s *UnitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitUnitStatement(s)
	}
}

func (p *YarsParser) UnitStatement() (localctx IUnitStatementContext) {
	this := p
	_ = this

	localctx = NewUnitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YarsParserRULE_unitStatement)

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
		p.SetState(104)
		p.Match(YarsParserUNIT)
	}
	{
		p.SetState(105)
		p.Match(YarsParserIDENTIFIER)
	}
	{
		p.SetState(106)
		p.FieldsDecl()
	}

	return localctx
}

// IFieldsDeclContext is an interface to support dynamic dispatch.
type IFieldsDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsDeclContext differentiates from other interfaces.
	IsFieldsDeclContext()
}

type FieldsDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDeclContext() *FieldsDeclContext {
	var p = new(FieldsDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_fieldsDecl
	return p
}

func (*FieldsDeclContext) IsFieldsDeclContext() {}

func NewFieldsDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDeclContext {
	var p = new(FieldsDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_fieldsDecl

	return p
}

func (s *FieldsDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDeclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(YarsParserL_CURLY, 0)
}

func (s *FieldsDeclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(YarsParserR_CURLY, 0)
}

func (s *FieldsDeclContext) AllFieldDecl() []IFieldDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDeclContext); ok {
			len++
		}
	}

	tst := make([]IFieldDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDeclContext); ok {
			tst[i] = t.(IFieldDeclContext)
			i++
		}
	}

	return tst
}

func (s *FieldsDeclContext) FieldDecl(i int) IFieldDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDeclContext); ok {
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

	return t.(IFieldDeclContext)
}

func (s *FieldsDeclContext) AllEos() []IEosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEosContext); ok {
			len++
		}
	}

	tst := make([]IEosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEosContext); ok {
			tst[i] = t.(IEosContext)
			i++
		}
	}

	return tst
}

func (s *FieldsDeclContext) Eos(i int) IEosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
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

	return t.(IEosContext)
}

func (s *FieldsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterFieldsDecl(s)
	}
}

func (s *FieldsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitFieldsDecl(s)
	}
}

func (p *YarsParser) FieldsDecl() (localctx IFieldsDeclContext) {
	this := p
	_ = this

	localctx = NewFieldsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, YarsParserRULE_fieldsDecl)
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
		p.SetState(108)
		p.Match(YarsParserL_CURLY)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == YarsParserIDENTIFIER {
		{
			p.SetState(109)
			p.FieldDecl()
		}
		{
			p.SetState(110)
			p.Eos()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(YarsParserR_CURLY)
	}

	return localctx
}

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YarsParserIDENTIFIER, 0)
}

func (s *FieldDeclContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (p *YarsParser) FieldDecl() (localctx IFieldDeclContext) {
	this := p
	_ = this

	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YarsParserRULE_fieldDecl)

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
		p.SetState(119)
		p.Match(YarsParserIDENTIFIER)
	}
	{
		p.SetState(120)
		p.Type_()
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *Type_Context) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *YarsParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, YarsParserRULE_type_)

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
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case YarsParserL_BRACKET:
		{
			p.SetState(122)
			p.ArrayType()
		}

	case YarsParserINT32, YarsParserUINT32, YarsParserFLOAT32, YarsParserBOOL, YarsParserSTRING:
		{
			p.SetState(123)
			p.BaseType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(YarsParserL_BRACKET, 0)
}

func (s *ArrayTypeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(YarsParserR_BRACKET, 0)
}

func (s *ArrayTypeContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *YarsParser) ArrayType() (localctx IArrayTypeContext) {
	this := p
	_ = this

	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, YarsParserRULE_arrayType)

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
		p.SetState(126)
		p.Match(YarsParserL_BRACKET)
	}
	{
		p.SetState(127)
		p.Match(YarsParserR_BRACKET)
	}
	{
		p.SetState(128)
		p.BaseType()
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(YarsParserINT32, 0)
}

func (s *BaseTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(YarsParserUINT32, 0)
}

func (s *BaseTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(YarsParserFLOAT32, 0)
}

func (s *BaseTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(YarsParserBOOL, 0)
}

func (s *BaseTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(YarsParserSTRING, 0)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (p *YarsParser) BaseType() (localctx IBaseTypeContext) {
	this := p
	_ = this

	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, YarsParserRULE_baseType)
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
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<YarsParserINT32)|(1<<YarsParserUINT32)|(1<<YarsParserFLOAT32)|(1<<YarsParserBOOL)|(1<<YarsParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YarsParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarsParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(YarsParserSEMI, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarsListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *YarsParser) Eos() (localctx IEosContext) {
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, YarsParserRULE_eos)

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
		p.Match(YarsParserSEMI)
	}

	return localctx
}

func (p *YarsParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *YarsParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
