package _hashcode.models;

import java.util.List;

public class HorizontalPicture extends Picture {

    public static final String ORIENTATION = "HORIZONTAL";

    public HorizontalPicture(List<String> tags) {
        super(ORIENTATION, tags);
    }
}
