import java.util.Random;

public class Main {
    private static int matrixLen = 32;
    private static long fakesum = 0;
    private static long iterations = 100000;


    public static void main(String[] args) {
        // SEQUENTIAL
        long sequentialDuration = 0;

        for (int i = 0; i < iterations; i++) {
            int[][] matrix1 = generateMatrix();
            int[][] matrix2 = generateMatrix();

            long start = System.nanoTime();
            int[][] resultingMatrix = Matrix.Run(matrix1, matrix2);
            long end = System.nanoTime();

            sequentialDuration += end - start;
            fakesum += resultingMatrix.length;
        }

        long sequentialMean = sequentialDuration / iterations;


        // CONCURRENT
        long concurrentDuration = 0;

        for (int i = 0; i < iterations; i++) {
            int[][] matrix1 = generateMatrix();
            int[][] matrix2 = generateMatrix();

            long start = System.nanoTime();
            int[][] resultingMatrix = ConcurrentMatrix.Run(matrix1, matrix2);
            long end = System.nanoTime();

            concurrentDuration += end - start;
            fakesum += resultingMatrix.length;
        }

        long concurrentMean = concurrentDuration / iterations;

        System.out.printf("Fakesum: %d\n", fakesum);
        System.out.printf("Matrix size: %d\n", matrixLen);
        System.out.printf("Iterations: %d\n\n\n", iterations);

        System.out.println("---SEQUENTIAL---");
        System.out.printf("Mean time: %f microseconds\n", (float) sequentialMean / 1000);

        System.out.println("---CONCURRENT---");
        System.out.printf("Mean time: %f microseconds\n", (float) concurrentMean / 1000);
    }

    public static int[][] generateMatrix() {
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
