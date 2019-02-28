package _hashcode.models;

import java.util.List;

public class VerticalPicture extends Picture {

    public static final String ORIENTATION = "VERTICAL";

    public VerticalPicture(int index, List<String> tags) {
        super(index, ORIENTATION, tags);
    }
}
