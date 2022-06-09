package coverage

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

const referenceShaSum = "7763355c8763f1ce2ff5583afd2687ba123934052e577031af4ee106d5546521"

func TestStudents(t *testing.T) {
	err := os.Rename("./autocode/students_test", "students_test.go")
	if err != nil {
		t.Fatalf("Failed to move test file: %v", err)
	}

	toBeTested, err := os.ReadFile("toBeTested.go")
	if err != nil {
		t.Fatalf("Failed to read toBeTested.go file: %v", err)
	}
	gotShaSum := fmt.Sprintf("%x", sha256.Sum256(toBeTested))
	if gotShaSum != referenceShaSum {
		t.Fatalf("toBeTested.go file was changed, do not edit or rename it\nexpected SHA256: %v\ngot SHA256: %v", referenceShaSum, gotShaSum)
	}

	cmd := exec.Command("go", "test", "-test.timeout=3m", "-cover", "students_test.go", "toBeTested.go")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		t.Errorf("Error during tests execution: %v", err)
		t.Fatalf("Execution output: \n%s", string(bytes.ReplaceAll(stdout.Bytes(), []byte("---"), []byte("--"))))
	}
	execOutput := stdout.Bytes()
	execOutput = execOutput[bytes.Index(execOutput, []byte("coverage: "))+10 : bytes.IndexByte(execOutput, '%')]
	coverage, err := strconv.ParseFloat(string(execOutput), 64)
	if err != nil {
		t.Fatalf("Can't determine coverage level: %v", err)
	}

	var skipRequired bool
	coverageTiers := []float64{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0}
	for _, coverageTier := range coverageTiers {
		t.Run(fmt.Sprintf("Coverage >= %v", coverageTier), func(t *testing.T) {
			if skipRequired {
				t.FailNow()
			}
			if coverage < coverageTier {
				skipRequired = true
				t.Fatalf("Coverage %v is less than %v", coverage, coverageTier)
			}
		})
	}
}
