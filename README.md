[![Go Report Card](https://goreportcard.com/badge/github.com/cookedsteak/timemix)](https://goreportcard.com/report/github.com/cookedsteak/timemix)

# timemix/intersection :alarm_clock:

## 干嘛的
#### 时间并集
```
t.Union()
```
对于零散的拥有 `开始值（时间`）和`结束值（时间）`的对象，
比如工单、操作记录，合并该对象的值并去重。

示例，可能有一个操作他会有多个离散的时间记录，
合并前：
```
开始时间：6:30:00 结束时间：9:15:00
开始时间：7:30:00 结束时间：10:30:00
开始时间：12:00:00 结束时间：15:06:00
开始时间：12:00:01 结束时间：19:13:06
```

合并后：
```
开始时间：6:30:00 结束时间：10:30:00
开始时间：12:00:00 结束时间：19:13:06
```

#### 时间交集
```
t.Intersect()
```
对 `开始值（时间`）和`结束值（时间）` 做交集运算。
返回值为只有一个元素的`TSlice`。

合并前：
```
开始时间：6:30:00 结束时间：9:15:00
开始时间：7:30:00 结束时间：10:30:00
开始时间：8:00:00 结束时间：11:06:00
```

合并后：
```
开始时间：8:00:00 结束时间：9:15:00
```

## What for

#### Union
```
t.Union()
```
To get the union from some time (or number) periods.
For example: working tickets, action records

Before:
```
start：6:30:00 end：9:15:00
start：7:30:00 end：10:30:00
start：12:00:00 end：15:06:00
start：12:00:01 end：19:13:06
```

After:
```
start：6:30:00 end：10:30:00
start：12:00:00 end：19:13:06
```

#### Intersection
```
t.Intersect()
```
To get the intersection of some time (or number) periods.
The return value is a `TSlice` with only one element.

Before：
```
start：6:30:00 end：9:15:00
start：7:30:00 end：10:30:00
start：8:00:00 end：11:06:00
```

After：
```
start：8:00:00 end：9:15:00
```
