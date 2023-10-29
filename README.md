# concurrent_stats
Package to explore concurrency in Go and how to use it for statistical analysis

### Running the Tests
To run all the tests use `go test ./...` from the root directory of the project.
Alternatively use `go test ./<package_name>` to run an individual package's
tests. For example to run the tests for the package `csvparser`:
    `go test ./csvparser`

### Benchmarking
Need to figure out how this is going to work...

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

The shared disadvantage of both approachs used are that they are computationally
expensive, but in this modern era computational power is plentiful. These
techniques are often selected for use with slower interpreted languages. If this
is considered good enough for use with those slower languages it will only be
faster using a compiled language such as Go.

### Resources Referenced
- [Wikipedia - Bootstrapping](https://en.wikipedia.org/wiki/Bootstrapping_(statistics))
- [Wikipedia - Cross Validation](https://en.wikipedia.org/wiki/Cross-validation_(statistics))
- [A Gentle Introduction to the Bootstrap Method](https://machinelearningmastery.com/a-gentle-introduction-to-the-bootstrap-method/)
- [Gonum/stat Package Documentation](https://pkg.go.dev/gonum.org/v1/gonum@v0.14.0/stat)
- [Investopedia: Multiple Linear Regression](https://www.investopedia.com/terms/m/mlr.asp)
- [California State University Long Beach Web Page](https://home.csulb.edu/~msaintg/ppa696/696regmx.htm)
