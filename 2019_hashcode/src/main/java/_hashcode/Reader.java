package _hashcode;

import _hashcode.models.HorizontalPicture;
import _hashcode.models.Picture;
import _hashcode.models.VerticalPicture;

import java.io.*;
import java.util.ArrayList;
import java.util.Arrays;

public class Reader {
    public static ArrayList<Picture> read(String path) {
        ArrayList<Picture> pictures = new ArrayList<>();
        File file = new File(path);
        BufferedReader br = null;
        try {
            br = new BufferedReader(new FileReader(file));
            int pictureCount = Integer.parseInt(br.readLine());
            for (int i=0; i<pictureCount; i++) {
                String inputString = br.readLine();
                String[] inputArray = inputString.split(" ", 3);
                String orientation = inputArray[0];
                String[] tags = inputArray[2].split(" ");
                if (orientation.equals("H")) {
                    pictures.add(new HorizontalPicture(Arrays.asList(tags)));
                } else {
                    pictures.add(new VerticalPicture(Arrays.asList(tags)));
                }
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return pictures;
    }
}
