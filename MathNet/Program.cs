using System;
using MathNet.Numerics.LinearAlgebra;

class Program
{
    static void Main(string[] args)
    {
        var matrixA = Matrix<double>.Build.DenseOfArray(new double[,] {
            {1,2,3},
            {4,5,6},
            {7,8,10}
        });

        var matrixB = Matrix<double>.Build.DenseOfArray(new double[,] {
            {3,2,1},
            {6,5,4},
            {9,8,7}
        });

        Console.WriteLine("Matrix A:");
        Console.WriteLine(matrixA);

        Console.WriteLine("\nMatrix B:");
        Console.WriteLine(matrixB);

        // Multiply matrices
        var result = matrixA * matrixB;
        Console.WriteLine("\nMultiplication of A and B:");
        Console.WriteLine(result);

        // Determinant of matrixA
        double determinantA = matrixA.Determinant();
        Console.WriteLine($"\nDeterminant of Matrix A: {determinantA}");

        // Inverse of matrixA
        if (matrixA.Determinant() != 0)
        {
            var inverseA = matrixA.Inverse();
            Console.WriteLine("\nInverse of Matrix A:");
            Console.WriteLine(inverseA);
        }
        else
        {
            Console.WriteLine("\nMatrix A is singular and doesn't have an inverse.");
        }
    }
}
