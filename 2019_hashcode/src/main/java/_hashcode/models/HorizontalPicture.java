package _hashcode.models;

import java.util.List;

public class HorizontalPicture extends Picture {

    public static final String ORIENTATION = "HORIZONTAL";

    public HorizontalPicture(int index, List<Integer> tags) {
        super(index, ORIENTATION, tags);
    }
}
