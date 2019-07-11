package main

import (
	"strings"
	"testing"

	"github.com/rendon/testcli"
)

func TestInputRequired(t *testing.T) {
	testcli.Run("./dechar", "-o", "table")
	if testcli.Success() {
		t.Fatalf("Expected to fail, but succeeded: %s", testcli.Error())
	}
}

func TestSimpleByDefault(t *testing.T) {
	testcli.Run("./dechar", "CHAR(81,120),CHAR(120,81)")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutMatches("Qx,xQ") {
		t.Fatalf("Expected decoded output, but failed: %s", testcli.Stdout())
	}

	if testcli.StdoutMatches("[+|\\-=]") {
		t.Fatalf("Expected simple output, but got table: %s", testcli.Stdout())
	}
	length := len(strings.Split(testcli.Stdout(), "\\n"))
	if length != 1 {
		t.Fatalf("Expected one line of output, instead got %d", length)
	}

	// if testcli.StdoutMatches("\\n*.*\\n.*\\n*") {
	// 	t.Fatalf("Expected simple output, but got lines: %s", testcli.Stdout())
	// }
}

func TestTableOutput(t *testing.T) {
	testcli.Run("./dechar", "-o", "table", "CHAR(81)")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
	if !testcli.StdoutMatches("^\\+\\-+\\+\\-+\\+") {
		t.Fatalf("Expected formatted table output, but did not get header: %s", testcli.Stdout())
	}
}

func TestLineOutput(t *testing.T) {
	testcli.Run("./dechar", "-o", "lines", "CHAR(81,120),CHAR(120,81)")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
	if testcli.StdoutContains(",") {
		t.Fatalf("Expected lines output, but found comma separated output: %s", testcli.Stdout())
	}
	length := len(strings.Split(testcli.Stdout(), "\n"))
	if length != 3 {
		t.Fatalf("Expected two lines of output, instead got %d", length)
	}
}

func TestFullDecode(t *testing.T) {
	testcli.Run("./dechar", "CHAR(45,120,49,45,81,45),CHAR(45,120,50,45,81,45),CHAR(45,120,51,45,81,45),CHAR(45,120,52,45,81,45)")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
	if !testcli.StdoutContains("-x1-Q-,-x2-Q-,-x3-Q-,-x4-Q-") {
		t.Fatalf("Expected known decoding, instead got: %s", testcli.Stdout())
	}
}
