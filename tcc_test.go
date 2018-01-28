package tcc

import "testing"
import "fmt"

func TestCompileString(t *testing.T) {
	program := `
	extern int printf(char* c,...);
	int main (int argc, char** argv) {
		printf("Hello %s!\n", argv[0]);
		return 0;
	}
	`
	tcc := NewTcc()
	defer tcc.Delete()
	tcc.SetOutputType(OUTPUT_MEMORY)
	tcc.SetLibPath("./tcc-0.9.27")
	err := tcc.CompileString(program)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		t.Fail()
		return
	}
	i, err := tcc.Run([]string{"world"})
	if i != 0 || err != nil {
		fmt.Printf("I: %d\n", i)
		fmt.Printf("ERROR: %v\n", err)
		t.Fail()
		return
	}
}
