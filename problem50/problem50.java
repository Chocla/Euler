/* https://projecteuler.net/problem=50
The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/
public class problem50 {
    public static final int MAX = 1000000;

    public static void main(String[] args) {
        System.out.println(findSpecialPrime(MAX));
    }

    public static int findSpecialPrime(int upperBound) {
        int maxPrime  = 0;
        int maxLength = 0;
        int[] primes = seive(upperBound);
        for(int i = 0; i < primes.length; i++) {
            int count = consecutiveSumLength(primes[i], primes);
            if(count > maxLength) {
                maxPrime = primes[i];
                maxLength = count;
            }
        }
        return maxPrime;
    }

    public static int consecutiveSumLength(int p, int[] primes) {
        int      count = 0;
        int partialSum = 0;
        for(int i = 0; primes[i] < p ;i++) {
            partialSum = primes[i];
            count = 1;
            for(int j = i+1; partialSum < p;j++, count++) {
                partialSum += primes[j];
            }
            if(partialSum == p) {
                return count;
            }
        }
        return 0;
    }
    //Sieve of Eratosthenes
    public static int[] seive(int max) {
        boolean[] seiveBools = new boolean[max];
        int       primeCount = 0;
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