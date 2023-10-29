# concurrent_stats
Package to explore concurrency in Go and how to use it for statistical analysis

### Running the Tests
To run all the tests use `go test ./...` from the root directory of the project.
Alternatively use `go test ./<package_name>` to run an individual package's
tests. For example to run the tests for the package `csvparser`:
    `go test ./csvparser`

### Running Benchmarks
The benchmarks for the serial and concurrent programs can be run with
`go test -bench=.`. The benchmarks are within the file
`regression_benchmark_test.go`.

If you're using a computer that has a shell that supports Bash scripts you can
make use of the `script run_benchmarks.sh` to repeatedly run the benchmark. The
benchmark itself will average the runtime over 100 trials, but this script
repeats the benchmarking process in it's entirety. This can be helpful in case
there are benchmarks that happen to run at a time when unkillable background
processes are particularly busy and increasing the runtime.

### Data and Techniques Used

#### Bootstrapping
Bootstrapping was chosen because of the relatively simple implementation and
because it can be parallelized relatively easily. Many of the substeps required
for bootstrapping are ["embarrassingly" parallel](https://en.wikipedia.org/wiki/Embarrassingly_parallel) and because of this
do not require complicated memory sharing patterns. The nature of the algorithm
also creates obvious points to stop and collect the parallel data before
continuing.

#### Cross Validation
Similarly to bootstrapping, the cross validation method chosen because of it's
relatively simple implementation and the relatively low complexity of
parallelizing the process.

#### Comparing Methods
The shared advantages of both approaches used are that they are relatively
simple algorithms to implement and that they have clear points of
parallelization.

The shared disadvantage of both approaches used are that they are computationally
expensive, but in this modern era computational power is plentiful. These
techniques are often selected for use with slower interpreted languages. If this
is considered good enough for use with those slower languages it will only be
faster using a compiled language such as Go.

#### Comparing Serial and Concurrent Runtimes
The concurrent/parallelized regressions consistently runs faster than the
serialized regressions and this is with the simplest possible parallelization.
Because concurrency is built into Go there are no additional dependencies to
make use of this and the implementation itself, routines and channels, have an
easy to understand syntax. The simplicity of concurrency with Go and the faster
runtimes from even the most basic implementation leave me wondering why anyone
would not use concurrency for heavy computational tasks such as the techniques
used in this program. It is worth noting that while the implementation of
concurrency in Go is simple, any code that uses concurrency will add complexity
that would make further refactoring more difficult. Reiterating an earlier point
, these techniques have clear points of parallelization which mitigates the
potential added complexity. Additionally, these techniques are well defined and
realistically refactors would be less structural and more in the realm of
changing data types or allowing additional options to be passed.

The most obvious areas that concurrency could be used to improve the runtime of
this particular program are within the bootstrapping method which uses
significantly more iteration than the cross validation method. Some of this
iteration could be removed by taking time to refactor some of the more
questionable design decisions that arose from time constraints, but as it stands
this would be where the most runtime gains would be main because it builds each
new sample by accessing the initial bootstrap sample over and over. Cross
validation would similarly benefit from parallelization, but because it simply
omits a value rather than constructing entirely new samples would not see as
much of an improvement in runtime.

### Resources Referenced
- [Wikipedia - Bootstrapping](https://en.wikipedia.org/wiki/Bootstrapping_(statistics))
- [Wikipedia - Cross Validation](https://en.wikipedia.org/wiki/Cross-validation_(statistics))
- [A Gentle Introduction to the Bootstrap Method](https://machinelearningmastery.com/a-gentle-introduction-to-the-bootstrap-method/)
- [Gonum/stat Package Documentation](https://pkg.go.dev/gonum.org/v1/gonum@v0.14.0/stat)
- [Investopedia: Multiple Linear Regression](https://www.investopedia.com/terms/m/mlr.asp)
- [California State University Long Beach Web Page](https://home.csulb.edu/~msaintg/ppa696/696regmx.htm)
