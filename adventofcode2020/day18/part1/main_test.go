package main

import "testing"

func TestNumberExpr(t *testing.T) {
	Run(&numberExpr{10}, 10, t)
}

func TestBinaryExpr(t *testing.T) {
	// 10 + 5
	expr := &binaryExpr{&numberExpr{10}, &numberExpr{5}, func(l, r int) int { return l + r }}
	Run(expr, 15, t)
}

func TestParenExpr(t *testing.T) {
	// 10 + (10*5)
	expr := &binaryExpr{
		&numberExpr{10},
		&parensExpr{
			&binaryExpr{
				&numberExpr{10},
				&numberExpr{5},
				func(l, r int) int { return l * r },
			},
		},
		func(l, r int) int { return l + r },
	}
	Run(expr, 60, t)
}

func TestCases(t *testing.T) {
	tables := []struct {
		expr string
		e    int
	}{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, table := range tables {
		p := newParser(table.expr)
		e := p.parseExpression()
		v := e.evaluate()
		if v != table.e {
			t.Errorf("Unexpected result from expression %s: expected: %d got: %d", table.expr, table.e, v)
		}
	}
}

func Run(expr expr, exp int, t *testing.T) {
	result := expr.evaluate()
	if result != exp {
		t.Errorf("Result of %T was incorrect, got : %d, expected: %d", expr, result, exp)
	}
}
