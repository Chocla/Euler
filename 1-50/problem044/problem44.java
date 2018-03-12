public class problem44 {
    public static final int UPPER_BOUND = 10000;
    public static void main(String[] args) {
        System.out.println(findValues());
    }

    public static int findValues() {
        int[] pentagons = generatePentagonNums(UPPER_BOUND);
        for(int i = 1; i < pentagons.length; i++) {
            for(int j = i; j < pentagons.length; j++) {
                if( isPentagonal(pentagons[i] + pentagons[j]) && isPentagonal(pentagons[j]-pentagons[i])) {
                    System.out.println(i + " " + j);
                    return pentagons[j] - pentagons[i];
                }
            }
        }
        System.out.println("Not found inside current upper bound");
        return -1;
    }
    public static int[] generatePentagonNums(int max) {
        int[] pentagons = new int[max];
        for(int i = 1; i < pentagons.length; i++) {
            pentagons[i] = (i * (3*i - 1) /2);
        }
        return pentagons;
    }
    public static boolean isPentagonal(int x) {
        double n = (Math.sqrt(24*x + 1) + 1 )/ 6;
        if (Math.abs(n - (int)n ) < 0.0001) {
            return true;
        } 
        return false;
    }
}