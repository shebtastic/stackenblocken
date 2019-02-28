package _hashcode.models;

import java.util.List;
import java.util.Set;

public class VerticalPicture extends Picture {

    public static final String ORIENTATION = "VERTICAL";

    public VerticalPicture(int index, Set<Integer> tags) {
        super(index, ORIENTATION, tags);
    }
}
