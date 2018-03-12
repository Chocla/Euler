public class problem43 {
    public static void main(String[] args) {
        problem43 p = new problem43();
        long ans = p.findAnswer();
        System.out.println(ans);
    }

    public long findAnswer() {
        int[] DIVISORS = {2,3,5,7,11,13,17};
        long answer = 0;
        for(long i = 1023456789; i <= 9876543210L;i++) {
            if(!isPandigital(i)) {
                continue;
            }
            boolean flag = true;
            for(int j = 2; j <= 8; j++) {
                int intPart = subInt(i, j);
                if(intPart % DIVISORS[j-2] != 0) {
                    flag = false;
                }
            }
            if(flag) {
                answer += i;
                System.out.println(i);
            }
        }
        return answer;
    }
    public boolean isPandigital(long n) {
        boolean[] digits = new boolean[10];
        while(n > 0) {
            digits[(int)(n % 10)] = true;
            n /= 10;
        }
        for(int i = 0; i < digits.length; i++) {
            if(!digits[i]) {
                return false;
            }
        }
        return true;
    }
    public int subInt(long n, int digit) {
        String t = Long.toString(n).substring(digit-1, digit+2);
        return Integer.parseInt(t);
    }

}