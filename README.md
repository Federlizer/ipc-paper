# IPC Paper
My IPC paper repository containing all files relevant to my IPC paper, including the paper itself.

### Algorithm

##### Go
A matrix multiplication algorithm, nothing too fancy. Create a two dimentional matrix that holds big.Int values
and pass it to an array multiplication function that goes through the array and multiplies them, returning the product.

###### Findings
In golang, the flat function (one that doesn't create any go routines and makes the multiplication on its own) is actually
significantly faster than the one that does spread out the multiplication in multiple go routines. That's possibly because
of the overhead that comes with actually creating the routines. Here are the stats that I've found so far:

Matrix length: 10 inner arrays (each array contains 9 or 10 values, from 1 to 99)
Routines took 66.191µs
Flat took 17.277µs

Matrix length: 50 inner arrays (each array contains 9 or 10 values, from 1 to 99)
Routines took 187.754µs
Flat took 50.677µs

### Notes
I had the realization that the way I was benchmarking wasn't as robust. I was multiplying the numbers into a single number,
instead a better way to benchmark concurrency is to have two matrices and multiply the cell of each matrix into a third matrix.
