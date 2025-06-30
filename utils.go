package main

const (
	MAX_STACK_SIZE  int    = 100
	OP_SHOW         string = "show"
	OP_PUSH         string = "push"
	OP_LVALUE       string = "lvalue"
	OP_RVALUE       string = "rvalue"
	OP_POP          string = "pop"
	OP_ASSIGN       string = ":="
	OP_PRINT        string = "print"
	OP_HALT         string = "halt"
	OP_LABEL        string = "label"
	OP_GOTO         string = "goto"
	OP_END          string = "end"
	OP_BEGIN        string = "begin"
	OP_CALL         string = "call"
	OP_RETURN       string = "return"
	OP_ADD          string = "+"
	OP_SUBRACT      string = "-"
	OP_MULTIPLY     string = "*"
	OP_DIVIDE       string = "/"
	OP_MODULO       string = "div"
	OP_AND          string = "&"
	OP_OR           string = "|"
	OP_NOT          string = "!"
	OP_NOT_EQUAL    string = "<>"
	OP_LESS_EQUAL   string = "<="
	OP_GREAT_EQUAL  string = ">="
	OP_LESS_THAN    string = "<"
	OP_GREATER_THAN string = ">"
)

func GetValue(item interface{}) int {
	switch val := item.(type) {
	case int:
		return val
	case *int:
		return *val
	default:
		log.Fatalf("Expected int or *int, got %T", item)
		return 0 // unreachable
	}
}
