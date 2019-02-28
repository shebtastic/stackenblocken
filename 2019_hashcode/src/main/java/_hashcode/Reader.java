package _hashcode;

import _hashcode.models.HorizontalPicture;
import _hashcode.models.Picture;
import _hashcode.models.VerticalPicture;

import java.io.*;
import java.util.*;

public class Reader {
    public static ArrayList<Picture> read(String path) {
        ArrayList<Picture> pictures = new ArrayList<>();
        File file = new File(path);
        Integer tagCounter = 0;
        HashMap<String, Integer> tagDic = new HashMap<>();
        BufferedReader br = null;
        try {
            br = new BufferedReader(new FileReader("inputs/" + file));
            int pictureCount = Integer.parseInt(br.readLine());
            for (int i=0; i<pictureCount; i++) {
                String inputString = br.readLine();
                String[] inputArray = inputString.split(" ", 3);
                String orientation = inputArray[0];
                String[] tags = inputArray[2].split(" ");
                Set<Integer> intTags = new HashSet<>();
                for (String tag : tags) {
                    if (tagDic.containsKey(tag)) {
                        intTags.add(tagDic.get(tag));
                    } else {
                        intTags.add(++tagCounter);
                        tagDic.put(tag, tagCounter);
                    }
                }
                if (orientation.equals("H")) {
                    pictures.add(new HorizontalPicture(i, intTags));
                } else {
                    pictures.add(new VerticalPicture(i, intTags));
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
