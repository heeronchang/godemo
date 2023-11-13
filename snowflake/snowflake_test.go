package snowflake

import (
	"testing"
)

func TestSnowflake_NextID(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		s := &Snowflake{
			workerid:     1,
			datacenterid: 1,
		}

		for i := 0; i < 100; i++ {
			got := s.NextID()
			t.Log(got)
		}
	})
}
