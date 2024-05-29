package main

type entry struct {
	value     string
	timestamp int
}
type TimeMap struct {
	tm map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{tm: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.tm[key] = append(this.tm[key], entry{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	stamps, ok := this.tm[key]
	if !ok {
		return ""
	}
	l, r := 0, len(stamps)
	for l < r {
		mid := (r-l)/2 + l
		if stamps[mid].timestamp == timestamp {
			return stamps[mid].value
		} else if stamps[mid].timestamp > timestamp {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r-1 < 0 {
		return ""
	}
	return stamps[r-1].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
