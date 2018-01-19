public class problem35 {
    public static void main(String[] args) {
        final int MAX = 1000000;
        int[] primes = seive(MAX);
        int answer = checkPrimes(primes);
        System.out.println(answer);
    }
    
    public static int checkPrimes(int[] primes) {
        int count = 0;
        for(int i = 0; i < primes.length; i++) {
            boolean flag = false;
            int tmp = primes[i];

            //check all rotations of primes[i]
            do {
                tmp = rotateDigits(tmp);
                //if tmp contains a 0 (-1), or is composite, prime isn't circular
                if(tmp == -1 || !isPrime(tmp,primes)) {
                    flag = true;
                    break;
                }
                
            } while(tmp != primes[i]);
            //circular prime found
            if(!flag) {
                count++;
            }
        }
        return count; 
    }

    //it just works
    public static int rotateDigits(int n) {
        int nc = n;
        int m = 0;
        int i = 10;
        for(; i < nc ; i *= 10, n /= 10) {
            if (n % 10 == 0) {
                return -1;
            }
            m += (n % 10) * i;
        }
        return m + n;
    }

    public static boolean isPrime(int n, int[] primes) {
        for(int i = 0; i < primes.length; i++) {
            if (primes[i] == n) {
                return true;
            }
            if(primes[i] > n) {
                return false;
            }
        }
        //not found in prime list
        return false;
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