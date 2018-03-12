public class problem47 {
    public static final int PRIME_MAX = 1000000;

    public static void main(String[] args) {
        int answer = findConsecutiveNums(4,4);
        System.out.println("Answer: " + answer);
    }

    public static int distinctFactors(int n, int[] primes) {
        int factorCount = 0;
        for(int i = 0; n > 1; i++) {
            if(n % primes[i] == 0) { //distinct factor found
                factorCount++;
                do { //factor it out as many times as possible
                    n /= primes[i];
                } while(n % primes[i] == 0);
            }
        }
        return factorCount;
    }
    public static int findConsecutiveNums(int length, int factorCount) {
        int[] primes = seive(PRIME_MAX);
        boolean found = false;
        int count = 0, i = 2;
        while( !found ) {
            int fCount = distinctFactors(i, primes);
            if (fCount == factorCount) {
                count++;
            } else {
                count = 0; //count resets when a num doesn't satisfy requirement
            }
            if(count == length) {
                found = true;
            }
            i++;
        }
        return i - length;
    }

    //Sieve of Eratosthenes
    public static int[] seive(int max) {
        boolean[] seiveBools = new boolean[max];
        int primeCount = 0;
        for(int i = 2; i < max; i++) {
            if (!seiveBools[i]) {
                primeCount++;
                for(int j = 2*i; j < max; j += i) {
                    seiveBools[j] = true;
                }   
            }
        }
        int[] primes = new int[primeCount];
        for (int i = 2,j = 0; i < max; i++) {
            if(!seiveBools[i]) {
                primes[j] = i;
                j++;
            }
        }
        return primes;
    }
}