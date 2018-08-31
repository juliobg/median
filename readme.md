median
======

This is a Go package that provides functions to build histogram of a values list and calculate its median using this histogram. I was just playing around with goroutines and wanted to do something like this... 

The algorithm chooses the bin where median should be located and creates a new histogram with narrower bins. After a few iterations the median is found (as soon as the bin where the median must be located as only one element).

A limitation of the algorith is that it requires upper and lower bounds for the values whose median we are trying to find. In practice the numeric limits of
the variable type you are using may be given (but this would slow down the calculation a bit).

As I already said, it is possible to parallelize the calculation (just paralellizing the creation of histograms and accumulating the resuls thereafter).

Some benchmarks comparing the one thread implementation with the parallel one and sorting.

Note that the functions do not use builtins, but median.Interface, with probably slows down things a bit. In any case, parallel histogram median calculation is found to be ~15x faster, for large sets of numbers. 

```
BenchmarkParallelMedian/2-4 	           10000	    155405 ns/op
BenchmarkParallelMedian/4-4 	           10000	    154123 ns/op
BenchmarkParallelMedian/8-4 	           10000	    154083 ns/op
BenchmarkParallelMedian/16-4         	   10000	    159452 ns/op
BenchmarkParallelMedian/32-4         	   10000	    159494 ns/op
BenchmarkParallelMedian/64-4         	   10000	    159270 ns/op
BenchmarkParallelMedian/128-4        	   10000	    160430 ns/op
BenchmarkParallelMedian/256-4        	   10000	    161299 ns/op
BenchmarkParallelMedian/512-4        	   10000	    160473 ns/op
BenchmarkParallelMedian/1024-4       	   10000	    171110 ns/op
BenchmarkParallelMedian/2048-4       	   10000	    194728 ns/op
BenchmarkParallelMedian/4096-4       	   10000	    192923 ns/op
BenchmarkParallelMedian/8192-4       	    5000	    288694 ns/op
BenchmarkParallelMedian/16384-4      	    5000	    295693 ns/op
BenchmarkParallelMedian/32768-4      	    2000	    796096 ns/op
BenchmarkParallelMedian/65536-4      	    1000	   1750518 ns/op
BenchmarkParallelMedian/131072-4     	    1000	   2003998 ns/op
BenchmarkParallelMedian/262144-4     	     300	   3898309 ns/op
BenchmarkParallelMedian/524288-4     	     100	  12733657 ns/op
BenchmarkParallelMedian/1048576-4    	      50	  23573297 ns/op
BenchmarkParallelMedian/2097152-4    	      50	  27509265 ns/op
BenchmarkParallelMedian/4194304-4    	      20	  95360383 ns/op
BenchmarkParallelMedian/8388608-4    	      10	 161270545 ns/op
BenchmarkParallelMedian/16777216-4   	       3	 360463878 ns/op

BenchmarkMedian/2-4                  	   50000	     26534 ns/op
BenchmarkMedian/4-4                  	   50000	     27027 ns/op
BenchmarkMedian/8-4                  	   50000	     36279 ns/op
BenchmarkMedian/16-4                 	   50000	     32538 ns/op
BenchmarkMedian/32-4                 	   50000	     30423 ns/op
BenchmarkMedian/64-4                 	   50000	     31239 ns/op
BenchmarkMedian/128-4                	   50000	     32136 ns/op
BenchmarkMedian/256-4                	   50000	     38467 ns/op
BenchmarkMedian/512-4                	   30000	     39656 ns/op
BenchmarkMedian/1024-4               	   30000	     47983 ns/op
BenchmarkMedian/2048-4               	   20000	     77231 ns/op
BenchmarkMedian/4096-4               	   20000	     83415 ns/op
BenchmarkMedian/8192-4               	   10000	    200232 ns/op
BenchmarkMedian/16384-4              	    5000	    248011 ns/op
BenchmarkMedian/32768-4              	    2000	   1012838 ns/op
BenchmarkMedian/65536-4              	    1000	   1887459 ns/op
BenchmarkMedian/131072-4             	     500	   3646129 ns/op
BenchmarkMedian/262144-4             	     200	   7469552 ns/op
BenchmarkMedian/524288-4             	     100	  19677482 ns/op
BenchmarkMedian/1048576-4            	      50	  38387869 ns/op
BenchmarkMedian/2097152-4            	      30	  54862659 ns/op
BenchmarkMedian/4194304-4            	      10	 153436800 ns/op
BenchmarkMedian/8388608-4            	       5	 276651754 ns/op
BenchmarkMedian/16777216-4           	       2	 589933200 ns/op

BenchmarkMedianSorting/2-4           	20000000	        99.5 ns/op
BenchmarkMedianSorting/4-4           	10000000	       138 ns/op
BenchmarkMedianSorting/8-4           	10000000	       225 ns/op
BenchmarkMedianSorting/16-4          	 3000000	       435 ns/op
BenchmarkMedianSorting/32-4          	 1000000	      1224 ns/op
BenchmarkMedianSorting/64-4          	  500000	      2306 ns/op
BenchmarkMedianSorting/128-4         	  300000	      5808 ns/op
BenchmarkMedianSorting/256-4         	  100000	     16851 ns/op
BenchmarkMedianSorting/512-4         	   30000	     49687 ns/op
BenchmarkMedianSorting/1024-4        	   10000	    122662 ns/op
BenchmarkMedianSorting/2048-4        	    5000	    275332 ns/op
BenchmarkMedianSorting/4096-4        	    2000	    604867 ns/op
BenchmarkMedianSorting/8192-4        	    1000	   1306330 ns/op
BenchmarkMedianSorting/16384-4       	     500	   2735971 ns/op
BenchmarkMedianSorting/32768-4       	     300	   5946819 ns/op
BenchmarkMedianSorting/65536-4       	     100	  12532066 ns/op
BenchmarkMedianSorting/131072-4      	      50	  26587496 ns/op
BenchmarkMedianSorting/262144-4      	      20	  56387899 ns/op
BenchmarkMedianSorting/524288-4      	      10	 120013874 ns/op
BenchmarkMedianSorting/1048576-4     	       5	 253278860 ns/op
BenchmarkMedianSorting/2097152-4     	       2	 525855975 ns/op
BenchmarkMedianSorting/4194304-4     	       1	1097517806 ns/op
BenchmarkMedianSorting/8388608-4     	       1	2358328327 ns/op
BenchmarkMedianSorting/16777216-4    	       1	4927177885 ns/op
```
