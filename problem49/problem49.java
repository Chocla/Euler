/*https://projecteuler.net/problem=49
The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: 
    (i) each of the three terms are prime, and, 
    (ii) each of the 4-digit numbers are permutations of one another.
There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property,
but there is one other 4-digit increasing sequence.
What 12-digit number do you form by concatenating the three terms in this sequence?
*/
import java.util.Arrays;
public class problem49 {
    public static final int PRIME_MAX = 10000;
    public static final int NUM_LENGTH = 4;

    public static void main(String[] args) {
        long t0 = System.nanoTime();
        findSequences();
        System.out.println("Time: " + ((System.nanoTime() - t0)/1e9));
    }

    public static void findSequences() {
        int[] primes = seive(PRIME_MAX);
        for(int p = 0; p < primes.length;p++) {
            //we are given that no sequence exists < 1000
            if(primes[p] < 1000) {
                continue;
            }
            //odd increments would give an even number (composite)
            for(int i = 2; i < (PRIME_MAX - primes[p]) / 2; i+= 2) {
                if (isPrime(primes[p]+i, primes) && isPermutation(primes[p], primes[p]+i)
                && isPrime(primes[p]+i+i, primes) && isPermutation(primes[p], primes[p]+i+i)) {
                    System.out.println("Valid Sequence: (" + primes[p] + "," + i +"): " + primes[p] + ", " + (primes[p]+i) + ", " + (primes[p]+i+i));
                }

            }
        }
    }

    //turns both args into arrays of digits and sorts them
    //if the two arrays are equal, the numbers are permutations of each other
    public static boolean isPermutation(int n, int m) {
        int[] digitsN = digitify(n);
        int[] digitsM = digitify(m);
        Arrays.sort(digitsN);
        Arrays.sort(digitsM);
        return Arrays.equals(digitsN, digitsM);
    }

    public static int[] digitify(int n) {
        int[] digits = new int[NUM_LENGTH];
        for(int i = 0; i < digits.length; i++, n/= 10) {
            digits[i] = n % 10;
        }
        return digits;
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
}