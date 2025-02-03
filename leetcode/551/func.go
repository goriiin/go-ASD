package main

func checkRecord(s string) bool {
	mas := []byte(s)

	var (
		ACounter, LCounter, maxLCounter int
	)

	for _, b := range mas {
		switch b {
		case 'A':
			{
				ACounter++
				if LCounter > maxLCounter {
					maxLCounter = LCounter
				}

				LCounter = 0
			}
		case 'L':
			{
				LCounter++
			}
		default:
			{
				if LCounter > maxLCounter {
					maxLCounter = LCounter
				}

				LCounter = 0
			}
		}
	}

	if LCounter > maxLCounter {
		maxLCounter = LCounter
	}

	return !(ACounter > 1 || maxLCounter > 2)
}
