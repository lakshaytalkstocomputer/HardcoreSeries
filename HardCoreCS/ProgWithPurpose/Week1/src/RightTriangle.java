public class RightTriangle {
    public static void main(String[] args) {
        assert args.length == 3 : "Number of Arguments should be 3";

        int a = Integer.parseInt(args[0]);
        int b = Integer.parseInt(args[1]);
        int c = Integer.parseInt(args[2]);

        if(a<0 || b<0 || c<0){
            System.out.println("false");
        } else if(a*a == (b*b + c*c) || c*c == (b*b + a*a) || b*b == (a*a + c*c)){
            System.out.println("true");
        } else{
            System.out.println("false");
        }

    }
}
