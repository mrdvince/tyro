type TimeMap struct {
    KvStore map[string][]interface{}
}


func Constructor() TimeMap {
    return TimeMap{}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if this.KvStore == nil {
        this.KvStore = make(map[string][]interface{})
    }
    this.KvStore[key] = append(this.KvStore[key], []interface{}{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    res := ""
    values, ok := this.KvStore[key]
    if !ok {
        values = make([]interface{},0)
    }
    l, r := 0, len(values) -1
    for l <= r {
        m := (l + r) / 2
        curr_slice, _ := values[m].([]interface{})
        curr_val := curr_slice[0].(string)
        curr_ts := curr_slice[1].(int)
        if curr_ts <= timestamp {
            res = curr_val
            l = m + 1
        } else {
            r = m -1
        }
    }
    return res
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */