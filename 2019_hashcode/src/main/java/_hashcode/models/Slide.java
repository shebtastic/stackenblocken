package _hashcode.models;

import java.util.Arrays;
import java.util.List;

public class Slide {
    public final List<Picture> PICTURES;

    public Slide(VerticalPicture verticalPicture, VerticalPicture secondVerticalPicture) {
        this.PICTURES = Arrays.asList(verticalPicture, secondVerticalPicture);
    }

    public Slide(HorizontalPicture horizontalPicture) {
        this.PICTURES = Arrays.asList(horizontalPicture);
    }

    public List<String> tags() {
        return PICTURES.size() == 2
            ? PICTURES.get(0).combineTags(PICTURES.get(1))
            : PICTURES.get(0).TAGS;
    }

    public int tagOverlap(Slide slide) {
        return Math.abs(tags().size() - slide.tags().size());
    }
}
