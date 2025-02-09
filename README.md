# GoStats
Statistical Package using Go vs R and Python

Github Repository: https://github.com/mwood8881/GoStats.git

## Project Description

## Look at branch copilot suggestions for the Readme.md file for the automated ai readme file.

This program uses Go to perform linear regression on the Anscombe Quartet dataset which is then compared to the results of running it through Python and R. It executes the time and memory to determine which program is best used for this linear regression task. 

The Anscombe Quartet is four datasets with different distributions. 

## Features
- Linear regression on Anscombe Quartet with Go
- Compares with Python and R versions
- Measures execution times for all programs
- Unit tests to check benchmarks

## Prerequisites
Install 
-Go
-Go testing tools
- Python and R

## Getting Started
1) Clone respository:
```bash
git clone https://github.com/yourusername/anscombe-go-stats.git
cd anscombe-go-stats ```

2) Install Go package stats:
```bash
go get github.com/montanaflynn/stats
```

3) 
```bash
go run main.go
```

4) Test and Benchmark
```bash
go run main.go linear_regression.go
go test -bench
```

5) Compare with results from python and R

## Results
The results for all of the programs all look the same withe the same resulting coefficients and slopes:
Go: 
Regression model coefficients for Set 1: Intercept: 3.00, Slope: 0.50
Regression model coefficients for Set 2: Intercept: 3.00, Slope: 0.50
Regression model coefficients for Set 3: Intercept: 3.00, Slope: 0.50
Regression model coefficients for Set 4: Intercept: 3.00, Slope: 0.50
Total time: 270.542µs

R: Coefficients:
Coefficients:
            Estimate Std. Error t value Pr(>|t|)   
(Intercept)   3.0001     1.1247   2.667  0.02573 * 
x1            0.5001     0.1179   4.241  0.00217 **
Coefficients:
            Estimate Std. Error t value Pr(>|t|)   
(Intercept)    3.001      1.125   2.667  0.02576 * 
x2             0.500      0.118   4.239  0.00218 **
            Estimate Std. Error t value Pr(>|t|)   
(Intercept)   3.0025     1.1245   2.670  0.02562 * 
x3            0.4997     0.1179   4.239  0.00218 **
Coefficients:
            Estimate Std. Error t value Pr(>|t|)   
(Intercept)   3.0017     1.1239   2.671  0.02559 * 
x4            0.4999     0.1178   4.243  0.00216 **
(64-bit)

Python:
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0001      1.125      2.667      0.026       0.456       5.544
x1             0.5001      0.118      4.241      0.002       0.233       0.767
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0009      1.125      2.667      0.026       0.455       5.547
x2             0.5000      0.118      4.239      0.002       0.233       0.767
const          3.0025      1.124      2.670      0.026       0.459       5.546
x3             0.4997      0.118      4.239      0.002       0.233       0.766
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0017      1.124      2.671      0.026       0.459       5.544
x4             0.4999      0.118      4.243      0.002       0.233       0.766

## Concerns:
People might not want to work with Go because R and Python have more libraries they can use in the packages. They aslo might be easier to use than Go for more complex problems. 

## Recommendation to Management:
I would recommend that data scientists use Go since it produces the same results but only takes Total time: 270.542µs. It shows that Go can handle basic statistical computations and efficiently. I would recommend they use Python or R however for visualizations or if the problem becomes more complex and requires more packages to be used. 

   
