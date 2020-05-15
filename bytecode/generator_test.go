package bytecode

import (
	"fmt"
	"testing"

	"github.com/guilhermeg2k/dusklang/ast"
	"github.com/guilhermeg2k/dusklang/vm"
)

func TestFullVarDeclaration(t *testing.T) {
	f := Function{
		Consts:          make(vm.Consts),
		Labels:          make(vm.Labels),
		VariablesOffset: make(VariablesOffset),
		bytecode:        []byte{},
	}
	fullVarDeclaration := ast.FullVarDeclaration{
		Variables: []ast.AutoVarDeclaration{
			ast.AutoVarDeclaration{
				Identifier: "x",
				Type:       "int",
			},
			ast.AutoVarDeclaration{
				Identifier: "y",
				Type:       "int",
				Expression: &ast.BinaryOperation{
					Operator: "+",
					Left: &ast.Literal{
						Type:  "number",
						Value: "10",
					},
					Right: &ast.Literal{
						Type:  "number",
						Value: "10",
					},
				},
			},
		},
	}
	generateFullVarDeclaration(&f, fullVarDeclaration)
	fmt.Println(f.bytecode)
}
func TestExpressionGeneration(t *testing.T) {
	function := Function{
		Consts: make(vm.Consts),
		Labels: make(vm.Labels),
	}
	var expression ast.Expression
	expression = &ast.BinaryOperation{
		Operator: "+",
		Left: &ast.Literal{
			Type:  "number",
			Value: "10",
		},
		Right: &ast.BinaryOperation{
			Operator: "/",
			Left: &ast.Literal{
				Type:  "number",
				Value: "4",
			},
			Right: &ast.BinaryOperation{
				Operator: "*",
				Left: &ast.Literal{
					Type:  "number",
					Value: "4",
				},
				Right: &ast.Literal{
					Type:  "number",
					Value: "2",
				},
			},
		},
	}
	generateExpression(&function, expression)
	virtualMachine := vm.VirtualMachine{
		Stack: &vm.Stack{},
	}
	main := vm.Function{
		Storage: &vm.Storage{},
	}
	main.Consts = function.Consts
	main.Bytecode = function.bytecode
	main.Bytecode = append(main.Bytecode, 255)
	returned := vm.Evaluate(&virtualMachine, &main)
	if returned[0] != 12 {
		t.Errorf(fmt.Sprintf("want %d, got %d", 12, returned[0]))
	}
}
