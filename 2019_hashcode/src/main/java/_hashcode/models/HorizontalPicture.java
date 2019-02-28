package _hashcode.models;

import java.util.List;
import java.util.Set;

public class HorizontalPicture extends Picture {

    public static final String ORIENTATION = "HORIZONTAL";

    public HorizontalPicture(int index, Set<Integer> tags) {
        super(index, ORIENTATION, tags);
    }
}
