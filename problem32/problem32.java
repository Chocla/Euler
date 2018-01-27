import java.util.ArrayList;

public class problem32 {
    public static void main(String[] args) {
        System.out.println(findAllPandigitalProds());
    }

    public static int findAllPandigitalProds() {
        int sum = 0;
        ArrayList<Integer> list = new ArrayList<Integer>();
        for(int i = 2; i < 1000; i++) {
            for(int j = 2; j < 2000 && i*j < 100000000;j++) { 
                if(list.contains(i*j)) {
                    continue;
                }
                if(isPandigital(i, j)) {
                    System.out.println(i + " "+j + " "+i*j);
                    list.add(i*j);
                }
            }
        }
        for(int i = 0; i < list.size(); i++) {
            sum += list.get(i);
        }
        return sum;
    }
    //checks if together n,m,n*m have all digits 1-9
    public static boolean isPandigital(int n, int m) {
        int product = n*m;
        int length = String.valueOf(n).length() + 
                     String.valueOf(m).length() +
                     String.valueOf(product).length();
        if(length != 9) {
            return false;
        }
        boolean[] digits = new boolean[9];
        while(n > 0) {
            if(n % 10 == 0) {
                n /= 10;
                continue;
            } 
            digits[(n % 10) - 1] = true;
            n /= 10;
        }
        while(m > 0) {
            if(m% 10 == 0) {
                m /= 10;
                continue;
            } 
            digits[(m % 10) - 1] = true;
            m /= 10;
        }
        while(product > 0) {
            if(product % 10 == 0) {
                product /= 10;
                continue;
            } 
            digits[(product % 10) - 1] = true;
            product /= 10;
        }
        for(int i = 0; i < digits.length; i++) {
            if(!digits[i]) {
                return false;
            }
        }
        return true;
    }
}