# Calmet

Calmet is a Go library created by a student for self-learning purposes in the fields of numerical methods
and
linear algebra. The library is divided into several packages, each focusing on a specific area of
mathematics.

## Packages

- [eigen](#eigen)
- [equations](#equations)
- [integral](#integral)
- [interpoly](#interpoly)
- [node](#node)
- [spline](#spline)
- [tools](#tools)

Calmet is an open-source project, and contributions from the community are welcome. By using and contributing to this
library, you can help create a valuable resource for students and developers interested in numerical analysis, linear
algebra, and interpolation.

Please note that while the library has been developed with care, it may still contain errors or inaccuracies. Users are
advised to verify the results independently and contribute to the project by reporting issues or submitting pull
requests.

# eigen

This package provides functions for finding eigenvalues and eigenvectors of a given matrix using different methods.

## Functions

1. **PowerMethod(A [][]float64, e float64) ([]float64, float64, float64, int)**: This function implements the Power
   Method for finding the maximum eigenvalue and its corresponding eigenvector of a given matrix `A`. It takes the
   matrix `A` and a precision `e` as input, and returns the eigenvector, eigenvalue, error, and the number of
   iterations.
2. **JacobiMethod(A [][]float64, eps float64) ([][]float64, []float64, []float64, int)**: This function implements the
   Jacobi Method for finding all eigenvalues and their corresponding eigenvectors of a given matrix `A`. It takes the
   matrix `A` and a precision `eps` as input, and returns the matrix of eigenvectors, eigenvalues, errors, and the
   number of iterations.

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/foreverNP/calmet/pkg/eigen"
)

func main() {
	A := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	e := 0.001

	// Power Method
	eigenvector, eigenvalue, error, iterations := eigen.PowerMethod(A, e)
	fmt.Println("Power Method:")
	fmt.Println("Eigenvector:", eigenvector)
	fmt.Println("Eigenvalue:", eigenvalue)
	fmt.Println("Error:", error)
	fmt.Println("Iterations:", iterations)

	// Jacobi Method
	eigenvectors, eigenvalues, errors, iterations := eigen.JacobiMethod(A, e)
	fmt.Println("\nJacobi Method:")
	fmt.Println("Eigenvectors:\n", eigenvectors)
	fmt.Println("Eigenvalues:", eigenvalues)
	fmt.Println("Errors:", errors)
	fmt.Println("Iterations:", iterations)
}
```

This example demonstrates how to use the functions provided by the `eigen` package to find eigenvalues and eigenvectors
of a given matrix using the Power Method and the Jacobi Method. The results, including the errors and the number of
iterations, are printed to the console.

# equations

This package provides functions to solve systems of linear algebraic equations (SLAE) using various methods.

## Functions

1. **GaussMethod(A [][]float64, B []float64) []float64**: Solves SLAE using the Gaussian elimination method. It takes a
   matrix of coefficients `A` and a vector of free terms `B`, and returns a vector of solutions `X`.
2. **JacobiMethod(A [][]float64, B []float64, e float64) ([]float64, int)**: Solves SLAE using the iterative Jacobi
   method. It takes a matrix of coefficients `A`, a vector of free terms `B`, and a precision `e`, and returns a vector
   of solutions `X` and the number of iterations `K`.
3. **ReflectionMethod(A [][]float64, B []float64) ([]float64, [][]float64, [][]float64)**: Solves SLAE using the
   reflection method. It takes a matrix of coefficients `A` and a vector of free terms `B`, and returns a vector of
   solutions `X`, an upper triangular matrix `R`, and an orthogonal matrix `Q`.
4. **RelaxationMethod(A [][]float64, B []float64, w float64, e float64) ([]float64, int)**: Solves SLAE using the
   relaxation method (with the Gauss-Seidel method as a special case when w == 1). It takes a matrix of
   coefficients `A`, a vector of free terms `B`, a relaxation factor `w`, and a precision `e`, and returns a vector of
   solutions `X` and the number of iterations `K`.
5. **SolveTridiagonal(A, B [][]float64) []float64**: Solves a system of linear equations with a tridiagonal matrix using
   the Thomas algorithm (also known as the tridiagonal matrix algorithm). It takes a matrix of coefficients `A` and a
   vector of free terms `B`, and returns a vector of solutions `X`.

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/foreverNP/calmet/pkg/equations"
)

func main() {
	A := [][]float64{
		{2, 1, 0},
		{1, 2, 1},
		{0, 1, 2},
	}
	B := []float64{1, 2, 3}

	X := equations.GaussMethod(A, B)
	fmt.Println("GaussMethod:", X)

	X, K := equations.JacobiMethod(A, B, 0.001)
	fmt.Println("JacobiMethod:", X, "Number of iterations:", K)

	X, R, Q := equations.ReflectionMethod(A, B)
	fmt.Println("ReflectionMethod:", X)
	fmt.Println("Upper triangular matrix R:", R)
	fmt.Println("Orthogonal matrix Q:", Q)

	X, K := equations.RelaxationMethod(A, B, 1.5, 0.001)
	fmt.Println("RelaxationMethod:", X, "Number of iterations:", K)

	A = [][]float64{
		{2, 1, 0},
		{1, 2, 1},
		{0, 1, 2},
	}
	B = [][]float64{
		{1},
		{2},
		{3},
	}
	X = equations.SolveTridiagonal(A, B)
	fmt.Println("SolveTridiagonal:", X)
}

```

This example demonstrates how to use the functions provided by the `equations` package to solve systems of linear
algebraic equations using various methods. It solves the same system of equations using the Gaussian elimination method,
the iterative Jacobi method, the reflection method, the relaxation method, and the Thomas algorithm for tridiagonal
matrices. The results, including the number of iterations for iterative methods, are printed to the console.

# integral

Package provides functions for numerical integration using various methods.

It is useful for calculating approximate values of integrals using various numerical integration methods. It
also provides error estimation using the Runge method for both Simpson's rule and the trapezoidal rule.

## Types

- **integrand**: Represents a function to be integrated. It is a function type that takes a `float64` as input and
  returns a `float64` as output.

## Functions

1. **IntegrateGaussLegendre(f integrand, a, b float64, n int) float64**: This function calculates the approximate value
   of the integral of the function `f` from `a` to `b` using the Gauss-Legendre quadrature formula with `n` nodes.

2. **IntegrateSimpson(f integrand, a, b float64, e float64, logFile \*os.File) float64**: This function calculates the
   approximate value of the integral of the function `f` from `a` to `b` with a given accuracy `e` using Simpson's rule
   and the Runge method for error estimation. It also logs the results to a file if `logFile` is not `nil`.

3. **IntegrateTrapezoidal(f integrand, a, b float64, e float64, logFile \*os.File) float64**: This function calculates
   the approximate value of the integral of the function `f` from `a` to `b` with a given accuracy `e` using the
   trapezoidal rule and the Runge method for error estimation. It also logs the results to a file if `logFile` is
   not `nil`.

# interpoly

Package provides structures and functions for Newton interpolation polynomials.
It is useful for creating and working with Newton interpolation polynomials. It allows you to calculate the
value of the polynomial at a given point and obtain a string representation of the polynomial.

## Structs

- **NewtonPoly**: Represents a Newton interpolation polynomial. It has two fields: `nodes` and `separatedDif`. `nodes`
  is a slice of `node.Node` structs representing the interpolation nodes, and `separatedDif` is a slice of `float64`
  representing the divided differences.

## Functions

1. **Solve(x float64) float64**: This method of the `NewtonPoly` struct calculates the value of the Newton interpolation
   polynomial at the point `x`.

2. **String() string**: This method of the `NewtonPoly` struct returns a string representation of the Newton
   interpolation polynomial.

3. **NewtonPolyBuilder(nodes []node.Node) NewtonPoly**: This function builds a Newton interpolation polynomial based on
   the given interpolation nodes. It takes a slice of `node.Node` structs representing the interpolation nodes and
   returns a `NewtonPoly` struct representing the built polynomial.

## Example Usage

Here's an example of how to use the `interpoly` package to create a Newton interpolation polynomial and calculate the
value of the polynomial at a given point.

```go
package main

import (
	"fmt"

	"github.com/foreverNP/calmet/pkg/interpoly"
	"github.com/foreverNP/calmet/pkg/node"
)

func main() {
	// Define interpolation nodes
	nodes := []node.Node{
		{X: 0.0, Y: 1.0},
		{X: 1.0, Y: 2.0},
		{X: 2.0, Y: 4.0},
		{X: 3.0, Y: 8.0},
		{X: 4.0, Y: 16.0},
	}

	// Build Newton interpolation polynomial
	newtonPoly := interpoly.NewtonPolyBuilder(nodes)

	// Print the polynomial
	fmt.Println("Newton interpolation polynomial:")
	fmt.Println(newtonPoly.String())

	// Calculate the value of the polynomial at x = 2.5
	x := 2.5
	result := newtonPoly.Solve(x)
	fmt.Printf("Value of the polynomial at x = %.1f: %.3f\n", x, result)
}
```

In this example, we first define the interpolation nodes as a slice of `node.Node` structs. We then use
the `interpoly.NewtonPolyBuilder` function to build a Newton interpolation polynomial based on these nodes. We print the
string representation of the polynomial using the `String` method. Finally, we calculate the value of the polynomial
at `x = 2.5` using the `Solve` method and print the result.

# node

Package provides structures and functions for nodes used in two-dimensional interpolation.

## Structs

- **Node**: Represents a node for interpolation in a two-dimensional space. It has two fields: `X` and `Y`, both of
  type `float64`, representing the coordinates of the node in the space for interpolation.

## Variables

- **gaussLegendreNodes**: A map containing nodes and weights for the Gauss-Legendre quadrature formula for n = 3, 4, 5.
  Each key in the map is an integer representing the number of nodes, and the corresponding value is a slice of `Node`
  structs.

## Functions

1. **BuildEquidistantNodes(f func(float64) float64, a float64, b float64, n int) []Node**: This function builds
   equidistant nodes for interpolation. It takes a function `f` to be interpolated, lower and upper bounds `a` and `b`,
   and the degree of the polynomial `n`. It returns a slice of `Node` structs representing the interpolation points.

2. **BuildChebyshevNodes(f func(float64) float64, a float64, b float64, n int) []Node**: This function builds Chebyshev
   nodes for interpolation. It takes a function `f` to be interpolated, lower and upper bounds `a` and `b`, and the
   degree of the polynomial `n`. It returns a slice of `Node` structs representing the interpolation points.

3. **BuildGaussLegendreNodes(a, b float64, n int) []Node**: This function builds nodes and weights for the
   Gauss-Legendre quadrature formula. It takes lower and upper bounds `a` and `b`, and the number of nodes `n`. It
   returns a slice of `Node` structs representing the interpolation points.

# spline

Package provides a function for creating a cubic spline interpolation based on a given derivative and set of
nodes.

It is useful for creating and working with cubic spline interpolations based on a given derivative and set of
nodes. It allows you to calculate the value of the cubic spline at a given point.

## Types

- **CubicSpline**: Represents a cubic spline interpolation for a given set of nodes. It has two fields: `nodes` (a slice
  of `node.Node` for interpolation) and `coeffs` (a slice of `float64` representing the coefficients for the cubic
  spline interpolation).

## Functions

1. **New(Df func(float64) float64, nodes []node.Node) CubicSpline**: This function creates a new cubic spline
   interpolation based on the given derivative `Df` and set of nodes `nodes`. It solves a system of linear algebraic
   equations (SLAE) using the tridiagonal matrix algorithm and returns the resulting cubic spline interpolation.

2. **Solve(x float64) (float64, error)**: This method of the `CubicSpline` struct calculates the value of the cubic
   spline at the point `x`. It returns the calculated value and an error if the argument is out of range.

## Example Usage

Here's an example of how to use the `spline` package to create a cubic spline interpolation and calculate the value of
the spline at a given point:

```go
package main

import (
	"fmt"
	"github.com/foreverNP/calmet/pkg/node"
	"github.com/foreverNP/calmet/pkg/spline"
)

func main() {
	// Define the derivative function Df(x) = 2x
	Df := func(x float64) float64 {
		return 2 * x
	}

	// Create a slice of nodes for interpolation
	nodes := []node.Node{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 4},
		{X: 3, Y: 9},
	}

	// Create a new cubic spline interpolation based on the derivative and nodes
	cs := spline.New(Df, nodes)

	// Calculate the value of the cubic spline at the point x = 2.5
	result, err := cs.Solve(2.5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The value of the cubic spline at the point x = 2.5 is: %.2f\n", result)
}
```

# tools

Package provides various mathematical operations and functions for vectors and matrices. This package is useful
for performing various mathematical operations in the field of linear algebra

1. **DotProduct(vector1, vector2 []float64) float64**: This function calculates the dot product of two vectors. It
   panics if the lengths of the vectors do not match.

2. **EuclideanNorm(vector []float64) float64**: This function calculates the Euclidean norm of a vector.

3. **SubtractVectors(vector1, vector2 []float64) []float64**: This function subtracts two vectors. It panics if the
   lengths of the vectors do not match.

4. **MultiplyMatrices(matrix1, matrix2 [][]float64) [][]float64**: This function multiplies two matrices. It panics if
   the number of columns in the first matrix does not match the number of rows in the second matrix.

5. **MatrixNorm(matrix [][]float64) float64**: This function calculates the cubic/row norm of a matrix.

6. **IsDiagonallyDominant(matrix [][]float64) bool**: This function checks if a matrix is diagonally dominant.

7. **MaxAbsoluteDifference(vec1, vec2 []float64) float64**: This function calculates the maximum absolute difference
   between the elements of two vectors.

8. **UniformNorm(vector []float64) float64**: This function calculates the uniform norm of a vector.

9. **MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64**: This function multiplies a matrix by a
   scalar.

10. **TransposeMatrix(matrix [][]float64) [][]float64**: This function transposes a given matrix.

11. **FindMaxOffDiagonalElement(A [][]float64) (int, int)**: This function finds the maximum off-diagonal element in a
    matrix.

12. **Off(A [][]float64) float64**: This function returns the sum of the squares of the off-diagonal elements in a
    matrix.