public class problem34 {
    public static void main(String[] args) {
        final int max = 7 * factorial(9);

        int answer = 0;
        for(int i = 10; i < max; i++) {
            if(i == digitFactorial(i)) {
                answer+= i;
            }
        }
        System.out.println(answer);
    }
    
    public static int digitFactorial(int n) {
        int digitFact = 0;
        while( n > 0) {
            digitFact += factorial(n % 10);
            n /= 10;
        }
        return digitFact;
    }
    public static int factorial(int m) {
        int f = 1 ;
        for(; m > 0; m--) {
            f *= m;
        }
        return f;
    }
}