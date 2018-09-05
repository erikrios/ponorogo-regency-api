/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package luas_segitiga;
import java.util.Scanner;
/**
 *
 * @author user
 */
public class Luas_Segitiga {

    /**
     * @param args the command line arguments
     */
    static int alas,tinggi;
    static double luas;
    static Scanner n;
    public static void main(String[] args) {
        // TODO code application logic here
        n=new Scanner(System.in);
        System.out.println("Masukkan Alas Segitiga : ");
        alas=n.nextInt();
        System.out.println("Masukkan Tinggi Segitiga : ");
        tinggi=n.nextInt();
        luas=alas*tinggi/2;
        System.out.print("Jadi luas segitiga tersebut adalah "+luas);
    }
    
}
