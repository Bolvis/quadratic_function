import java.util.Scanner;

public class kwadratowa{

    static boolean positiveDelta = false;

    public static void main(String[] args){
        float a,b,c;
        Scanner scanner = new Scanner(System.in);
        do{
            System.out.print("Set a ->");
            a = scanner.nextFloat();
        }while(a == 0);
        System.out.print("Set b ->");
        b = scanner.nextFloat();
        System.out.print("Set c ->");
        c = scanner.nextFloat();
        scanner.close();

        float[] r = kwadratowa(a, b, c);

        if(positiveDelta){
            System.out.printf("x1 = %f\n",r[1]);
            System.out.printf("x2 = %f\n",r[2]);
        }
        System.out.printf("delta = %f\n",r[0]);
    }

    static float[] kwadratowa(float a, float b, float c){
        float[] r = new float[3];
        float delta = (float)Math.pow(b,2) + (-4*a*c);
        r[0] = delta;
        if(0 <= delta){
            r[1] = (float)(-1*b - Math.sqrt(delta));
            r[2] = (float)(-1*b + Math.sqrt(delta));
            positiveDelta = true;
        }else positiveDelta = false;
        return r;
    }
}