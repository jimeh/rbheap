# rbheap [![Build Status](https://travis-ci.org/jimeh/rbheap.png)](https://travis-ci.org/jimeh/rbheap)

A tool for working with and analyzing Ruby processes' heap dumps created via
`ObjectSpace.dump_all`.

## How to dump the heap

Dumping the heap from a running Ruby process is not very difficult as long as
your app has required the [`rbtrace`](https://github.com/tmm1/rbtrace) gem. The
steps are:

1. Add the `rbtrace` gem to your application, and require it in your code when
   the application boots.
2. Find the PID of your running Ruby process.
3. Run the following to dump your process' heap to a file called `heap.jsonl`:
    ```bash
    rbtrace -p $PID -e 'Thread.new{GC.start;require "objspace";File.open("heap.jsonl","w"){|f|ObjectSpace.dump_all(output: f)}}'
    ```

## Commands

### `leak`

The leak command is intended to help track down memory leaks. By requiring three
heap dumps as input, it attempts to find memory that showed up in dump #2, and
is still there in #3.

The idea is to take a heap dump shortly after the application starts and before
it's had much of a chance to leak memory. Then take another heap dump after it's
been running for a while and leaked memory. And finally take a third heap dump
after it's been running for a while longer and leaked even more.

But comparing these three dumps and extracting only the objects which are held
in memory during heap dumps #2 and #3, we should mostly be left with objects
which are leaked memory.

```
Usage:
  rbheap leak [flags] <dump-A> <dump-B> <dump-C>

Flags:
  -f, --format string   output format: "hex" / "full" (default "hex")
  -h, --help            help for leak
  -v, --verbose         print verbose information
```

## Credits

This project is mostly based on the ideas and concepts from the following
articles:

- [What I Learned About Hunting Memory Leaks in Ruby
2.1](http://blog.skylight.io/hunting-for-leaks-in-ruby/) by Peter Wagenet
- [How I spent two weeks hunting a memory leak in
  Ruby](http://www.be9.io/2015/09/21/memory-leak/) by Oleg Dashevskii
- [Debugging memory leaks in
  Ruby](https://samsaffron.com/archive/2015/03/31/debugging-memory-leaks-in-ruby)
  by Sam Saffron
