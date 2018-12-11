package pwm

import (
	"os"
	"testing"
)

func TestPWM(t *testing.T) {
	writeFile := func(f string, c []byte, p os.FileMode) error {
		return nil
	}
	New()
	d := &driver{
		sysfs:     SysFS,
		writeFile: writeFile,
	}
	if err := d.Export(0); err != nil {
		t.Error(err)
	}
	if err := d.Frequency(0, 100); err != nil {
		t.Error(err)
	}
	if err := d.Enable(0); err != nil {
		t.Error(err)
	}
	if err := d.DutyCycle(0, 50); err != nil {
		t.Error(err)
	}
	if err := d.Disable(0); err != nil {
		t.Error(err)
	}
	if err := d.Unexport(0); err != nil {
		t.Error(err)
	}
}
