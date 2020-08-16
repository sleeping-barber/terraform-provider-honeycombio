package honeycombio

import (
	"context"
	"testing"

	"github.com/kvrhdn/go-honeycombio"
)

func createTriggerWithRecipient(t *testing.T, c *honeycombio.Client, dataset string, recipient honeycombio.TriggerRecipient) (trigger *honeycombio.Trigger, deleteFn func()) {
	ctx := context.Background()

	trigger = &honeycombio.Trigger{
		Name: "Terraform provider - acc test trigger recipient",
		Query: &honeycombio.QuerySpec{
			Calculations: []honeycombio.CalculationSpec{
				{
					Op: honeycombio.CalculateOpCount,
				},
			},
		},
		Threshold: &honeycombio.TriggerThreshold{
			Op:    honeycombio.TriggerThresholdOpGreaterThan,
			Value: &[]float64{100}[0],
		},
		Recipients: []honeycombio.TriggerRecipient{recipient},
	}
	trigger, err := c.Triggers.Create(ctx, dataset, trigger)
	if err != nil {
		t.Error(err)
	}

	return trigger, func() {
		err := c.Triggers.Delete(ctx, dataset, trigger.ID)
		if err != nil {
			t.Error(err)
		}
	}
}
