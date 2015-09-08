package ov
import (
	"os"
	"testing"

	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
)

// testing power state type
func TestPowerState(t *testing.T) {
	var (
		c *OVClient
		testData    string = "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57"
		expectsData string = "2M25090RMW"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// get a blade object
		blade, err := c.GetServerHardware(testData)
		assert.NoError(t, err, "GetServerHardware threw error -> %s", err)
		assert.Equal(t, expectsData, blade.SerialNumber)
		// get a power state object
		var pt *PowerTask
		pt = pt.NewPowerTask(blade)
		pt.Timeout = 46 // timeout is 20 sec
		assert.Equal(t, expectsData, pt.Blade.SerialNumber)

		// Test the power state executor to off
		log.Info("------- Setting Power to Off")
		err = pt.PowerExecutor(P_OFF)
		assert.NoError(t, err, "PowerExecutor threw no errors -> %s", err)

		// Test the power state executor to on
		log.Info("------- Setting Power to On")
		err = pt.PowerExecutor(P_ON)
		assert.NoError(t, err, "PowerExecutor threw no errors -> %s", err)

		// Test the power state executor to off and leave off
		log.Info("------- Setting Power to Off")
		err = pt.PowerExecutor(P_OFF)
		assert.NoError(t, err, "PowerExecutor threw no errors -> %s", err)

	}
}