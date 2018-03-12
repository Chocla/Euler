public class problem38 {
    public static void main(String[] args) {
        problem38 p = new problem38();
        int m = p.findNum();
        System.out.println(m);
    }

    public int findNum() {
        int max = 0;
        for(int i = 2; i < 100000; i++) {
            int concatNum = makeNum(i);
            if( isPandigital(concatNum) && concatNum > max) {
                max = concatNum;
            }
        }
        return max;
    }
    public int makeNum(int n) {
        String num = "";
        for(int i = 1; num.length() < 9; i++) {
            num += Integer.toString(n*i);
        }
        if(num.length() != 9) {
            return 0;
        }
        return Integer.parseInt(num);
    }

    public boolean isPandigital(int n) {
        boolean[] digits = new boolean[9];
        while(n > 0) {
            if(n % 10 == 0) {return false;}
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

}