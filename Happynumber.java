class Solution {
    public boolean isHappy(int n) {
        HashSet set = new HashSet<>();
        int out = n;
        while (sumOfSquareOfDigits(out) != 1) {
            out = sumOfSquareOfDigits(out);
            // System.out.println(set);
            if (set.contains(out)) {
                return false;
            } else {
                set.add(out);
            }

        }
        return true;
    }

    public int sumOfSquareOfDigits(int input) {
        int output = 0;

        while (input != 0) {
            int rem = input % 10;
            output = output + (rem * rem);
            input = input / 10;
        }

        return output;
    }
}