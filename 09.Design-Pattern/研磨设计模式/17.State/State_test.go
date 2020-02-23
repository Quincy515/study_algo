package State

func ExampleWeek() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next(ctx)
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}

	// Output:
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
	// Sunday
}
