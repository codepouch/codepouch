import java.io.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

        // first line is the number of test cases
        int cases = Integer.parseInt(reader.readLine());
        for (int i = 0; i < cases; i++) {

            // first line of each case is the number of integers
            int count = Integer.parseInt(reader.readLine());

            int[] integers = new int[count];
            String[] strings = reader.readLine().trim().split("\\s+");
            assert integers.length == strings.length;

            for (int k = 0; k < count; k++) {
                integers[k] = Integer.parseInt(strings[k]);
            }

            Solution solution = new Solution();
            int output = solution.singleNumber(integers);

            System.out.println(output);
        }
    }
}
