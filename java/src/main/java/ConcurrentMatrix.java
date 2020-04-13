import java.util.ArrayList;

public class ConcurrentMatrix extends Thread {
    private int startingRow;
    private int[][] matrix1;
    private int[][] matrix2;
    private int[][] resultingMatrix;

    public ConcurrentMatrix(int[][] matrix1, int[][] matrix2, int startingRow) {
        this.startingRow = startingRow;
        this.matrix1 = matrix1;
        this.matrix2 = matrix2;

        this.resultingMatrix = new int[matrix1.length][matrix1[0].length];
    }

    public void run() {
        for (int row = 0; row < matrix1.length; row++) {
            for (int col = 0; col < matrix1[row].length; col++) {
                resultingMatrix[row][col] = matrix1[row][col] * matrix2[startingRow + row][col];
            }
        }
    }

    public int[][] getResultingMatrix() {
        return this.resultingMatrix;
    }

    public static int[][] Run(int[][] matrix1, int[][] matrix2) {
        ArrayList<ConcurrentMatrix> threads = new ArrayList<ConcurrentMatrix>();
        int[][] result = new int[matrix1.length][matrix1[0].length];

        for (int row = 0; row < matrix1.length; row += 2) {
            int[][] split = new int[2][matrix1[0].length];
            split[0] = matrix1[row];
            split[1] = matrix1[row+1];

            ConcurrentMatrix concurrentMatrix = new ConcurrentMatrix(split, matrix2, row);
            threads.add(concurrentMatrix);
            concurrentMatrix.start();
        }

        int currentRow = 0;
        for (int i = 0; i < threads.size(); i++) {
            try {
                ConcurrentMatrix thread = threads.get(i);
                thread.join();

                int[][] resultingMatrix = thread.getResultingMatrix();
                result[currentRow] = resultingMatrix[0];
                result[currentRow+1] = resultingMatrix[1];
                currentRow += 2;

            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        return result;
    }
}
