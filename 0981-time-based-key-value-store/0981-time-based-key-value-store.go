type TimeMap struct {
    KvStore map[string][]KVPair
}

type KVPair struct {
    value string
    timestamp int
}

func Constructor() TimeMap {
    return TimeMap{}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if this.KvStore == nil {
        this.KvStore = make(map[string][]KVPair)
    }
    this.KvStore[key] = append(this.KvStore[key], KVPair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    values, ok := this.KvStore[key]
    if !ok {
        return ""
    }
    l, r := 0, len(values) -1
    for l <= r {
        m := (l + r) / 2
        if values[m].timestamp <= timestamp {
            if m == len(values) -1 || values[m+1].timestamp > timestamp {
                return values[m].value
            }
            l = m + 1
        } else {
            r = m -1
        }
    }
    return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */