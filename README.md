
# Procolli

Procolli displays files in the proc filesystem as machine-readable JSON.

## Usage

```bash
❯ procolli /proc/pressure/cpu --watch 1000
```

```json
{"Some":{"Avg10":2.85,"Avg60":2.83,"Avg300":1.83,"Total":162275990},"Full":{"Avg10":0.56,"Avg60":0.74,"Avg300":0.86,"Total":107276495}}
{"Some":{"Avg10":2.85,"Avg60":2.83,"Avg300":1.83,"Total":162290720},"Full":{"Avg10":0.56,"Avg60":0.74,"Avg300":0.86,"Total":107286969}}
{"Some":{"Avg10":2.51,"Avg60":2.77,"Avg300":1.82,"Total":162309222},"Full":{"Avg10":0.64,"Avg60":0.75,"Avg300":0.86,"Total":107297297}}
{"Some":{"Avg10":2.51,"Avg60":2.77,"Avg300":1.82,"Total":162335117},"Full":{"Avg10":0.64,"Avg60":0.75,"Avg300":0.86,"Total":107308164}}
{"Some":{"Avg10":2.24,"Avg60":2.71,"Avg300":1.82,"Total":162349376},"Full":{"Avg10":0.71,"Avg60":0.76,"Avg300":0.86,"Total":107318717}}
```

## Supported Files

```
/proc/buddyinfo
/proc/cmdline
/proc/cpuinfo
/proc/loadavg
/proc/mdstat
/proc/meminfo
/proc/net/dev
/proc/net/protocols
/proc/net/sockstat
/proc/net/sockstat6
/proc/net/stat
/proc/net/tcp
/proc/net/tcp6
/proc/net/udp
/proc/net/udp6
/proc/net/unix
/proc/pressure
/proc/pressure/<resource>
/proc/schedstat
/proc/slabinfo
/proc/stat
/proc/swaps
/proc/vmstat
/proc/zoneinfo
```

```
/proc/<pid>/cgroup
/proc/<pid>/cmdline
/proc/<pid>/comm
/proc/<pid>/cwd
/proc/<pid>/environ
/proc/<pid>/limits
/proc/<pid>/maps
/proc/<pid>/net/dev
/proc/<pid>/schedstat
/proc/<pid>/stat
```

## License

The MIT License

Copyright (c) 2022 Róisín Grannell

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
