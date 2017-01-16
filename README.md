# go-hex

[![GoDoc](https://godoc.org/github.com/tmthrgd/go-bitwise?status.svg)](https://godoc.org/github.com/tmthrgd/go-bitwise)
[![Build Status](https://travis-ci.org/tmthrgd/go-bitwise.svg?branch=master)](https://travis-ci.org/tmthrgd/go-bitwise)

Efficient bitwise (xor/and/and-not/or) implementations for Golang.

go-bitwise provides bitwise operations using SSE/AVX instructions on x86-64.

## Download

```
go get github.com/tmthrgd/go-bitwise
```

## Benchmark

```
BenchmarkXOR/15-8     	100000000	        15.6 ns/op	 958.78 MB/s
BenchmarkXOR/32-8     	200000000	         9.23 ns/op	3467.88 MB/s
BenchmarkXOR/128-8    	100000000	        11.7 ns/op	10895.13 MB/s
BenchmarkXOR/1K-8     	50000000	        34.2 ns/op	29899.36 MB/s
BenchmarkXOR/16K-8    	 2000000	       787 ns/op	20811.37 MB/s
BenchmarkXOR/128K-8   	  200000	      9936 ns/op	13190.61 MB/s
BenchmarkXOR/1M-8     	   20000	     89205 ns/op	11754.67 MB/s
BenchmarkXOR/16M-8    	     500	   3056743 ns/op	5488.59 MB/s
BenchmarkXOR/128M-8   	      50	  24597236 ns/op	5456.62 MB/s
BenchmarkAnd/15-8     	100000000	        18.7 ns/op	 800.10 MB/s
BenchmarkAnd/32-8     	200000000	         9.23 ns/op	3467.38 MB/s
BenchmarkAnd/128-8    	100000000	        11.8 ns/op	10840.26 MB/s
BenchmarkAnd/1K-8     	50000000	        34.2 ns/op	29922.25 MB/s
BenchmarkAnd/16K-8    	 2000000	       787 ns/op	20804.43 MB/s
BenchmarkAnd/128K-8   	  200000	      9938 ns/op	13188.18 MB/s
BenchmarkAnd/1M-8     	   20000	     91050 ns/op	11516.47 MB/s
BenchmarkAnd/16M-8    	     500	   3044681 ns/op	5510.34 MB/s
BenchmarkAnd/128M-8   	      50	  24351110 ns/op	5511.77 MB/s
BenchmarkAndNot/15-8  	100000000	        18.8 ns/op	 799.65 MB/s
BenchmarkAndNot/32-8  	200000000	         9.26 ns/op	3456.80 MB/s
BenchmarkAndNot/128-8 	100000000	        11.9 ns/op	10799.33 MB/s
BenchmarkAndNot/1K-8  	50000000	        34.4 ns/op	29806.72 MB/s
BenchmarkAndNot/16K-8 	 2000000	       791 ns/op	20692.48 MB/s
BenchmarkAndNot/128K-8         	  200000	     10043 ns/op	13050.53 MB/s
BenchmarkAndNot/1M-8           	   20000	     90389 ns/op	11600.61 MB/s
BenchmarkAndNot/16M-8          	     500	   3060622 ns/op	5481.63 MB/s
BenchmarkAndNot/128M-8         	      50	  24505583 ns/op	5477.03 MB/s
BenchmarkOr/15-8               	100000000	        18.8 ns/op	 797.68 MB/s
BenchmarkOr/32-8               	200000000	         9.29 ns/op	3444.37 MB/s
BenchmarkOr/128-8              	100000000	        11.9 ns/op	10796.04 MB/s
BenchmarkOr/1K-8               	50000000	        34.8 ns/op	29403.06 MB/s
BenchmarkOr/16K-8              	 2000000	       790 ns/op	20724.55 MB/s
BenchmarkOr/128K-8             	  200000	      9995 ns/op	13112.48 MB/s
BenchmarkOr/1M-8               	   20000	     90165 ns/op	11629.42 MB/s
BenchmarkOr/16M-8              	     500	   3054965 ns/op	5491.79 MB/s
BenchmarkOr/128M-8             	      50	  24489454 ns/op	5480.63 MB/s
BenchmarkNot/15-8  	100000000	        18.2 ns/op	 825.06 MB/s
BenchmarkNot/32-8  	100000000	        10.2 ns/op	3147.15 MB/s
BenchmarkNot/128-8 	100000000	        13.1 ns/op	9793.64 MB/s
BenchmarkNot/1K-8  	50000000	        31.3 ns/op	32753.75 MB/s
BenchmarkNot/16K-8 	 3000000	       425 ns/op	38519.25 MB/s
BenchmarkNot/128K-8         	  300000	      5092 ns/op	25738.26 MB/s
BenchmarkNot/1M-8           	   20000	     63538 ns/op	16503.00 MB/s
BenchmarkNot/16M-8          	    1000	   2070027 ns/op	8104.83 MB/s
BenchmarkNot/128M-8         	     100	  18581626 ns/op	7223.14 MB/s
```

## License

Unless otherwise noted, the go-bitwise source files are distributed under the Modified BSD License
found in the LICENSE file.
