package _hashcode.models;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Slide {
    public final List<Picture> PICTURES;
    public final List<Integer> TAGS;

    public Slide(VerticalPicture verticalPicture, VerticalPicture secondVerticalPicture) {
        this.PICTURES = Arrays.asList(verticalPicture, secondVerticalPicture);
        this.TAGS = PICTURES.get(0).combineTags(PICTURES.get(1));
    }

    public Slide(HorizontalPicture horizontalPicture) {
        this.PICTURES = Arrays.asList(horizontalPicture);
        this.TAGS = PICTURES.get(0).TAGS;
    }

    public int tagOverlap(Slide slide) {
        return (int) this.TAGS.stream().filter(tag -> slide.TAGS.contains(tag)).count();
    }

    @Override
    public String toString() {
        return PICTURES.stream()
                .map(Picture::toOutput)
                .collect(Collectors.joining(" "));
    }

    public String toOutput() {
        return PICTURES.stream()
                .map(Picture::toOutput)
                .collect(Collectors.joining(" "));
    }
}
