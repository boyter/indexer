# indexer

Code for GopherConSyd 2023

So please clone this, and start interacting!

It's a small portion of the caisson index that powers searchcode.com with no dependencies.

Compile it and then run `indexer` in a directory with some code (not too much as it takes a while to index).

Run against this repository you should see something like the below,

```
$ indexer
currentBlockDocumentCount:40 currentDocumentCount:40 currentBlockStartDocumentCount:0
enter search term: 
test
--------------
11 index result(s)

queryTime:2 matchTime:0
enter search term: 
.git/hooks/pre-rebase.sample
20. if test "$#" = 2
47. if test -z "$not_in_master"
56. if test "$only_next_1" = "$only_next_2"
59. 	if test -z "$not_in_topic"
107.  * Whenever you need to test or publish your changes to topic
```
