package rulang

// NodeType identifies various AST nodes
type NodeType int

// Node is an interface for AST nodes
type Node interface {
	NodeType() NodeType
	String() string
}

// Package represents a package name .
//
// Example: package rulang
type Package struct {
	Name string
}

// Import represents an import . Imports can also be aliased.
//
// Examples:
// 		import "os"
// 		import ru "github.com/eliquious/rulang"
type Import struct {
	Alias  string
	Module string
}

// TypeDefinition defines a type.
//
// Examples:
// 		type Person struct
// 		    Name string 		// FieldDeclaration
// 		    Age  uint8
// 		end
//
//		type List [T] struct
//		    Array []T
//		end
type TypeDefinition struct {
	Name     string
	Extends  []string
	Generics []string
	Fields   []FieldDeclaration
}

// FieldDeclaration represents a field in a TypeDefinition.
//
// Fields are public if they are capitalized. If they are lowercase, the fields are
// only available inside the same package.
type FieldDeclaration struct {
	Name     string
	IsPublic bool
	Type     TypeDeclaration
}

// TypeDeclaration represents the type of a field, argument or return value.
// Examples:
// 		uint, float, string, etc
// 		[]string
// 		[]List[T]
//
// 		func(Error)
type TypeDeclaration struct {
	IsArray  bool
	Name     string
	Generics []GenericTypeDeclaration

	IsFunc    bool
	Arguments []TypeDeclaration
}

// GenericTypeDeclaration represents a generic type.
// Examples:
// 		T, []T
type GenericTypeDeclaration struct {
	Name    string
	IsArray bool
}

// EnumDefinition defines an enumeration. Enums can implement an interface as well.
// Example:
// 		enum StatusCode
//			Ok, InternalServerError, NotFound, Redirect
// 		end
type EnumDefinition struct {
	Name         string
	Enumerations []string
}

type InterfaceDefinition struct {
	Name      string
	Generics  []GenericTypeDeclaration
	Functions []FunctionSignature
}

type FunctionSignature struct {
	Name       string
	Generics   []GenericTypeDeclaration
	Arguments  []ArgumentDeclaration
	ReturnType ReturnType
}

// FunctionDefinition defines a function. Functions gan also have
// generics. Functions can also return a tuple of values.
type FunctionDefinition struct {
	FunctionSignature
	Body []Node
}

// Argument represents an argument in a function or macro.
type ArgumentDeclaration struct {
	Variable string
	Type     TypeDeclaration
}

// ReturnType represents the types a function returns.
type ReturnType struct {
	Types []TypeDeclaration
}

// VariableAssignment represents a variable assignment.
// Example:
// 		var i int
// 		var list List[int]
type VariableAssignment struct {
	Variable string
	Type     TypeDeclaration
	Value    Node
}

// InferredVariableAssignment represents a variable assignment with an inferred type.
// Example:
// 		router := http.Default()
type InferredVariableAssignment struct {
	VariableAssignment
}

// FunctionCall represents a function call.
// Examples:
// 		router := http.Default()
//		router.GET("/ping", func(w http.ResponseWriter, req http.Request) {
//		    w.String(200, "pong")
//		})
//		router.Run(":8080")
type FunctionCall struct {
	Package   string
	Variable  string
	Method    string
	Arguments []Node
	Generics  []TypeDeclaration
}

type LiteralValue struct {
	Value string
}

type VariableLiteral struct {
	LiteralValue
}

type StringLiteral struct {
	LiteralValue
}

type IntegerLiteral struct {
	LiteralValue
}

type FloatingPointLiteral struct {
	LiteralValue
}

type CharacterLiteral struct {
	Value rune
}

type StringInterpolationLiteral struct {
	Value string
}

type StructLiteral struct {
	Type     string
	Generics []TypeDeclaration
	Values   []Node
}

type StructFieldLiteral struct {
	FieldName string
	Value     Node
}

type FunctionLiteral struct {
	Arguments  []ArgumentDeclaration
	ReturnType ReturnType
	Body       []Node
}

type InterfaceImplementation struct {
	Type      string
	Generics  []string
	Functions []FunctionDefinition
}

type MacroDefinition struct {
	Name          string
	Generics      []GenericTypeDeclaration
	ArgumentTypes []TypeDeclaration
}

type MacroImplementation struct {
	Name      string
	Generics  []TypeDeclaration
	Arguments []ArgumentDeclaration
	Body      []Node
}

type BinaryOperationDefinition struct {
	Operation  string
	Arguments  [2]ArgumentDeclaration
	ReturnType []TypeDeclaration
	Body       []Node
}

type UnaryOperationDefinition struct {
	Operation  string
	Arguments  ArgumentDeclaration
	ReturnType []TypeDeclaration
	Body       []Node
}

type SymbolLiteral struct {
	Body []Node
}

type SymbolDefinition struct {
	Name string
	Body []Node
}

type SymbolArgument struct {
	Name string
}

type SymbolEvaluation struct {
	Name string
}

type ChannelDeclaration struct {
	Type TypeDeclaration
}

type IfElseStatement struct {
	Blocks []ConditionalBlock
}

type ConditionalBlock struct {
	HasCondition bool
	Condition    Node
	Body         []Node
}

type WhileStatement struct {
	Condition Node
	Body      []Node
}

type ForRangeStatement struct {
	Variables []VariableLiteral
	Iterator  Node
	Body      []Node
}
