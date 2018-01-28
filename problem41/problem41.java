public class problem41 {
    public static void main(String[] args) {
        problem41 p = new problem41();
        int ans = p.findAnswer();
        System.out.println(ans);
    }

    public int findAnswer() {
        int max = 0;
        int[] primes = seive(100000000);
        for(int i = 0; i < primes.length; i++) {
            if(isPandigital(primes[i]) && primes[i] > max) {
                max = primes[i];
            }
        }
        return max;

    }

    public boolean isPandigital(int n) {
        int len = Integer.toString(n).length();
        boolean[] digits = new boolean[len];
        while(n > 0) {
            if(n % 10 > len || n % 10 == 0) {return false;}
            digits[(n % 10) - 1] = true;
            n /= 10;
        }
        for(int i = 0; i < digits.length; i++) {
            if(!digits[i]) {
                return false;
            }
        }
        return true;
    }

    public int[] seive(int max) {
        boolean[] bools = new boolean[max];
        int count = 0;
        for(int i = 2; i < bools.length; i++) {
            if(!bools[i]) {
                count++;
                for(int j = i+i; j < bools.length; j+=i) {
                    bools[j] = true;
                }
            }
        }
        int[] primes = new int[count];
        for (int i = 2,j = 0; i < max; i++) {
            if(!bools[i]) {
                primes[j] = i;
                j++;
            }
        }
        return primes;
    }
}