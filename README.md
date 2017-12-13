# go-cp
Golang version of the command line utility `cp(1)`.  
Made just for fun to benchmark performance (and take a first stab at Golang).

## Usage
```
./go-cp SOURCE DEST
```
**NOTE:** Recursive copying and clobbering is enabled by default.

## Benchmarks
The following benchmarks were recorded for a file of size 1.5G on a 4-core machine, coreutils version `8.26`.

* Destination file does not exist -

  |       | user (s) | system (s) | CPU (%) | total (s) |
  |:-----:|:--------:|:----------:|:-------:|:---------:|
  |   cp  |   0.00   |    0.93    |    99   |   0.933   |
  | go-cp |   0.02   |    0.94    |    99   |   0.957   |

* Destination file exists _(I'm interested in figuring out the cause of these stats!)_ -

  |       | user (s) | system (s) | CPU (%) | total (s) |
  |:-----:|:--------:|:----------:|:-------:|:---------:|
  |   cp  |   0.00   |    1.34    |    10   |  12.809   |
  | go-cp |   0.01   |    0.63    |    99   |   0.642   |

## Notes
The following points are noted for the second case (destination file exists) -
* Performance gain for `go-cp` and the corresponding loss for `cp(1)` is quite noticeable.
* CPU utilization for `cp(1)` is quite low as compared to first.
* It's worth mentioning that before running either of the utilities, `sync` was run to synchronize cached writes. Moreover, the difference in the total durations is still as significant as above if the caches are dropped via `echo 3 > /proc/sys/vim/drop_caches`.
