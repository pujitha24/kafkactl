package consume

import (
	"bytes"
	"testing"

	"github.com/deviceinsight/kafkactl/v5/internal/output"
)

func TestPrintMessage_DefaultLineSeparator(t *testing.T) {
	var buf bytes.Buffer
	output.IoStreams = output.IOStreams{Out: &buf}

	value := "hello"
	msg := &message{Value: &value}

	if err := printMessage(msg, Flags{Separator: "#"}); err != nil {
		t.Fatalf("printMessage failed: %v", err)
	}

	if buf.String() != "hello\n" {
		t.Fatalf("expected %q, got %q", "hello\n", buf.String())
	}
}

func TestPrintMessage_CustomLineSeparator(t *testing.T) {
	var buf bytes.Buffer
	output.IoStreams = output.IOStreams{Out: &buf}

	key := "line1\nline2"
	value := "line3\nline4"
	msg := &message{Key: &key, Value: &value}

	flags := Flags{Separator: "#", PrintKeys: true, LineSeparator: "~~"}

	if err := printMessage(msg, flags); err != nil {
		t.Fatalf("printMessage failed: %v", err)
	}

	expected := "line1\nline2#line3\nline4~~"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}

func TestPrintMessage_ControlCharLineSeparator(t *testing.T) {
	var buf bytes.Buffer
	output.IoStreams = output.IOStreams{Out: &buf}

	value := "hello"
	msg := &message{Value: &value}

	flags := Flags{Separator: "#", LineSeparator: "\\t"}

	if err := printMessage(msg, flags); err != nil {
		t.Fatalf("printMessage failed: %v", err)
	}

	expected := "hello\t"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}
