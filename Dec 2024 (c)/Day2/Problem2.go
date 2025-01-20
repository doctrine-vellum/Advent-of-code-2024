package main

func Problem2() int {
	numSafeReports := 0;
	reports := ReadFile();
	for _, report := range reports {
		if len(report) < 2 {
			continue;
		}
		unsafeNum := checkReport(report);
		if unsafeNum == -1 {
			numSafeReports += 1;
			continue;
		}
		subReport1 := make([]int, len(report) - 1);
		subReport2 := make([]int, len(report) - 1);

		for i := 0; i < unsafeNum; i++ {
			subReport1[i] = report[i]; 
			subReport2[i] = report[i]; 
		}
		
		subReport1[unsafeNum] = report[unsafeNum]; 
		subReport2[unsafeNum] = report[unsafeNum + 1]; 

		for i := unsafeNum + 1; i < len(report) - 1; i++ {
			subReport1[i] = report[i + 1]; 
			subReport2[i] = report[i + 1]; 
		}

		if checkReport(subReport1) == -1 {
			numSafeReports += 1;
		} else if checkReport(subReport2) == -1 {
			numSafeReports += 1;
		} else if unsafeNum == 1 && checkReport(report[1:]) == -1 {
			numSafeReports += 1;	
		}
	}
	return numSafeReports;
}

func checkReport(report []int) int {
	asc := 1;
	if report[1] - report[0] < 0 {
		asc = -1;
	}
	for i := 0; i < len(report) - 1; i++ {
		if diff := report[i + 1] - report[i]; diff * asc < 1 || diff * asc > 3 {
			return i;
		}
	}
	return -1;
}
