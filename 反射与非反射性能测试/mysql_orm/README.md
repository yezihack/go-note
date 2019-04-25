[TOC]

# 基本信息
> 目的测试一下反射的性能,拿原生的SQL查询和Gorm查询对比


# 基准测试
```
go test -bench=. -cpu=1,2,3,4 -count=1 -benchmem
```

# 基准结果

```
BenchmarkGetBlackListByOri             3         360152671 ns/op        126615634 B/op   1947154 allocs/op
BenchmarkGetBlackListByOri-2           3         388286756 ns/op        126615714 B/op   1947154 allocs/op
BenchmarkGetBlackListByOri-3           3         337075252 ns/op        126616333 B/op   1947164 allocs/op
BenchmarkGetBlackListByOri-4           2         638743885 ns/op        126616384 B/op   1947165 allocs/op
BenchmarkGetBlackListByOrm             1        1784319825 ns/op        429050424 B/op   8671340 allocs/op
BenchmarkGetBlackListByOrm-2           1        1525479786 ns/op        429032952 B/op   8671085 allocs/op
BenchmarkGetBlackListByOrm-3           1        1733162526 ns/op        429032584 B/op   8671077 allocs/op
BenchmarkGetBlackListByOrm-4           1        1670718181 ns/op        429033016 B/op   8671078 allocs/op
```