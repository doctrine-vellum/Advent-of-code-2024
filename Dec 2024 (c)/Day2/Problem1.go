package main

func Problem1() int {
	numSafeReports := 0;
	reports := ReadFile();
	for _, report := range reports {
		asc := 1;
		if len(report) < 2 {
			continue;
		}
		if report[1] - report[0] < 0 {
			asc = -1;
		}
		for i := 0; i < len(report) - 1; i++ {
			if diff := report[i + 1] - report[i]; diff * asc < 1 || diff * asc > 3 {
				numSafeReports -= 1;
				break;
			}
		}
		numSafeReports += 1;
		
	}
	return numSafeReports;
}
