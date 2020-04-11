public class FiveRandomStatistic {
    public static void main(String args[]){
        double avg=0, min=2, max=0;

        for(int i = 0 ;i<5; i++){
            double number = Math.random();
            min = Math.min(min, number);
            max = Math.max(max, number);
            avg += number;

            System.out.println(i+ ": " + number);

        }

        avg = avg/5;
        System.out.println("Average of Numbers: "+avg);
        System.out.println("Minimum of Numbers: "+min);
        System.out.println("Maximum of Numbers: "+max);

    }
}
