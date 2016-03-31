ev
==

This is an **e**xecution **v**isualizer for Go.

Usage
-----

You underscore import the package into your main program, and do no further
changes there.

```go
import _ "github.com/calmh/ev"
```

This does nothing unless the environment variable `EV_INTERVAL` is set at
startup. If it is, resource usage statistics are printed to stdout every
`EV_INTERVAL` milliseconds together with your program's normal output.

```
jb@syno:~/s/g/c/ev $ EV_INTERVAL=1000 testprog
Hello, world!
ev pid 11866 @1459409325892: 5 gr, 1 ms user, 3 ms sys, 299 KiB alloc, 4472 objs, 299 KiB totalloc, 1636 KiB sys, 736 KiB heap, 288 KiB stack, 4096 KiB nextgc, 0 ms gcpause, 0 gcs
ev pid 11866 @1459409326896: 5 gr, 2 ms user, 3 ms sys, 308 KiB alloc, 4500 objs, 308 KiB totalloc, 1892 KiB sys, 640 KiB heap, 384 KiB stack, 4096 KiB nextgc, 0 ms gcpause, 0 gcs
ev pid 11866 @1459409327897: 5 gr, 2 ms user, 4 ms sys, 310 KiB alloc, 4515 objs, 310 KiB totalloc, 1892 KiB sys, 608 KiB heap, 416 KiB stack, 4096 KiB nextgc, 0 ms gcpause, 0 gcs
```

You can interpret this data manually if you like but it's more powerful to
use the program `ev` to visualize the trace. To do this, simply pipe the
output through `ev`. It will filter out the execution trace lines but pass
on other output.

```
jb@syno:~/s/g/c/ev $ EV_INTERVAL=1000 testprog | ev
Hello, world!
```

You can of course also pipe a saved log to `ev`.

`ev` starts a web server, by default on localhost port 8080. Visit that with
your browser and you'll get a nice overview of what's going on.