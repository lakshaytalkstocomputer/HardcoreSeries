public class HelloGoodbye {
    public static void main(String[] args) {
        try {
            System.out.println("Hello " + args[0] + " and " + args[1]+".");
            System.out.println("GoodBye " + args[1] + " and " + args[0]+".");
        }
        catch (ArrayIndexOutOfBoundsException e){
            System.out.println("No Arguments Given.");
        }
    }
}
