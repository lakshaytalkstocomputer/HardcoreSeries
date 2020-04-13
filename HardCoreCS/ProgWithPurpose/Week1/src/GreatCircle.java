public class GreatCircle {
    public static void main(String[] args) {

        double x_1 = Math.toRadians(Float.parseFloat(args[0]));
        double y_1 = Math.toRadians(Float.parseFloat(args[1]));
        double x_2 = Math.toRadians(Float.parseFloat(args[2]));
        double y_2 = Math.toRadians(Float.parseFloat(args[3]));

        double distance = 2 * 6371.0 * Math.asin(
                        Math.sqrt(
                                Math.sin((x_2 - x_1)/2)* Math.sin((x_2 - x_1)/2) +
                                Math.cos(x_1) * Math.cos(x_2) * Math.sin((y_2 - y_1)/2)*Math.sin((y_2 - y_1)/2)
                        )
                );

        System.out.println(distance + " kilometers");
    }
}
