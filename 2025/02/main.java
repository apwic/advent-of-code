import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

public class main {
    public static void main(String[] args) throws Exception {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        System.out.println("Part 1: " + part1(lines));
        System.out.println("Part 2: " + part2(lines));
    }

    static long part1(List<String> lines) {
        String line = lines.get(0);
        String[] ids = line.split(",");
        long ans = 0;

        for (int i = 0; i < ids.length; i++) {
            String id = ids[i];
            String[] nums = id.split("-");
            long int1 = Long.parseLong(nums[0]);
            long int2 = Long.parseLong(nums[1]);

            for (long n : invalidIds(int1, int2))
                ans += n;
        }

        return ans;
    }

    static long part2(List<String> lines) {
        return 0;
    }

    static List<Long> invalidIds(long a, long b) {
        List<Long> ids = new ArrayList<Long>();
        for (long x = a; x <= b; x++) {
            String xStr = Long.toString(x);

            if (isInvalid(xStr)) {
                ids.add(x);
            }
        }

        return ids;
    }

    static Boolean isInvalid(String x) {
        if (x.length() % 2 == 1) {
            return false;
        }

        int mid = x.length() / 2;
        String firstPart = x.substring(0, mid);
        String secondPart = x.substring(mid, x.length());
        return firstPart.equals(secondPart);
    }
}
