package domain_enums_test

import (
	"testing"

	domain_enums "github.com/alaa-aqeel/looply-app/src/core/Domain/enums"
)

func TestCommandStatus_Label(t *testing.T) {
	if domain_enums.Pending.Label() != "pending" {
		t.Fatal("Pending label incorrect")
	}
	if domain_enums.Done.Label() != "done" {
		t.Fatal("Done label incorrect")
	}
}

func TestCommandStatusForm_Int(t *testing.T) {
	status, err := domain_enums.NewCommandStatus(1)
	if err != nil || status != domain_enums.Process {
		t.Fatal("Failed to parse process")
	}
}
