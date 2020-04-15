import java.util.Random;

public class Main {
    public static void main(String[] args) {
        int matrixLen = 8;

        while (matrixLen <= 2048) {
            System.out.printf("%-20s %8s %11s\n\n", "name", "mean(ns)", "iterations");
            int iterations = 2;

            while (iterations <= 1048576) {
                int fakesum = 0;
                int[][] matrix1 = generateMatrix(matrixLen);
                int[][] matrix2 = generateMatrix(matrixLen);

                long start = System.nanoTime();
                for (int i = 0; i < iterations; i++) {
                    fakesum += ConcurrentMatrix.Run(matrix1, matrix2).length;
                }
                long end = System.nanoTime();
                long elapsed = end - start;

                long meanTime = elapsed / iterations;

                System.out.printf(
                        "%-20s %8d %11d\n",
                        "matrix length " + matrixLen,
                        meanTime,
                        iterations
                );

                iterations *= 2;
            }

            System.out.println("\n---MATRIX DOUBLED---\n");
            matrixLen *= 2;
        }
    }

    public static int[][] generateMatrix(int matrixLen) {
        Random rand = new Random();
        int[][] matrix = new int[matrixLen][matrixLen];

        for (int row = 0; row < matrix.length; row++) {
            for (int col = 0; col < matrix[row].length; col++) {
                matrix[row][col] = rand.nextInt(100) + 1;
            }
        }

        return matrix;
    }
}
