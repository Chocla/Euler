/*https://projecteuler.net/problem=39

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of p ≤ 1000, is the number of solutions maximised?
*/

public class problem39 {
    public static void main(String[] args) {
        long t0 = System.nanoTime();
        System.out.println(maximizeSolutions());
        System.out.println("Time: " + ((System.nanoTime() - t0)/1e9)); //~1 minute
        
    }
    //brute force checks every perimeter below 1000 for all possible right triangles
    public static int maximizeSolutions() {
        int max = 0;
        int maxP = 3;
        for( int p = 3; p <= 1000; p++) {
            
            int tmpSoln = findTriangleSolutions(p);
            if(tmpSoln > max) {
                max = tmpSoln;
                maxP = p;
            }
        }
        return maxP;
    }
    public static int findTriangleSolutions(int perimeter) {
        int solnCount = 0;
        for(int c = 1; c < perimeter / 2 ;c++) {
            for(int a = 1; a < perimeter - c ; a++) {
                for(int b = 1; b < a; b++) {
                    if((a + b + c == perimeter) && isRightTriangle(a, b, c)) {
                        
                        solnCount++;
                    }
                }
            }
        }

        return solnCount;
    }

    public static boolean isRightTriangle(int a, int b, int c) {
        return a*a + b*b == c*c;
    }
}