public class Matrix {
    public static int[][] Run(int[][] matrix1, int[][] matrix2) {
        int[][] result = new int[matrix1.length][matrix1.length];

        for (int row = 0; row < matrix1.length; row++) {
            for (int col = 0; col < matrix1[row].length; col++) {
                result[row][col] = matrix1[row][col] * matrix2[row][col];
            }
        }

        return result;
    }
}
