package designundergroundsystem

type UndergroundSystem struct {
	TicketIn  map[int]Ticket1
	TicketOut map[string]Ticket2
}

type Ticket1 struct {
	Start_Station string
	Time          int
}

type Ticket2 struct {
	Time  int
	Count int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[int]Ticket1),
		make(map[string]Ticket2),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.TicketIn[id] = Ticket1{
		Start_Station: stationName,
		Time:          t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	thick1 := this.TicketIn[id]
	str := thick1.Start_Station + "_" + stationName
	if _, found := this.TicketOut[str]; !found {
		this.TicketOut[str] = Ticket2{
			Time:  t - thick1.Time,
			Count: 1,
		}
	} else {
		record := this.TicketOut[str]
		this.TicketOut[str] = Ticket2{
			Time:  record.Time + (t - thick1.Time),
			Count: record.Count + 1,
		}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	var sum float64
	var count int
	str := startStation + "_" + endStation

	sum = float64(this.TicketOut[str].Time)
	count = this.TicketOut[str].Count

	return sum / float64(count)
}
