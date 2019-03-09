package _hashcode;

import _hashcode.models.HorizontalPicture;
import _hashcode.models.Picture;
import _hashcode.models.VerticalPicture;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;
import java.util.*;

public class Reader {
    private static final Logger LOGGER = LoggerFactory.getLogger(Reader.class);

    public static ArrayList<Picture> read(String path) {
        ArrayList<Picture> pictures = new ArrayList<>();
        File file = new File(path);
        Integer tagCounter = 0;
        int tagsPerPictureSum = 0;
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
                tagsPerPictureSum += tags.length;
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

            int tagsPerPicture = tagsPerPictureSum / pictureCount;
            LOGGER.info("read " + path + " with " + pictureCount + " pictures and " + tagCounter + " tags and ~" + tagsPerPicture + " tags/picture");
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return pictures;
    }
}
