public class RightTriangle {
    public static void main(String[] args) {

        double a = Double.parseDouble(args[0]);
        double b = Double.parseDouble(args[1]);
        double c = Double.parseDouble(args[2]);

        System.out.println( ( c*c == a*a + b*b || b*b == a*a + c*c || a*a == b*b + c*c) && (c > 0 && b > 0 && a > 0 ));
    }
}
