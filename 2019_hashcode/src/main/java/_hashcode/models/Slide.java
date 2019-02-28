package _hashcode.models;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Slide {
    public final List<Picture> PICTURES;

    public Slide(VerticalPicture verticalPicture, VerticalPicture secondVerticalPicture) {
        this.PICTURES = Arrays.asList(verticalPicture, secondVerticalPicture);
    }

    public Slide(HorizontalPicture horizontalPicture) {
        this.PICTURES = Arrays.asList(horizontalPicture);
    }

    public List<Integer> tags() {
        return PICTURES.size() == 2
            ? PICTURES.get(0).combineTags(PICTURES.get(1))
            : PICTURES.get(0).TAGS;
    }

    public int tagOverlap(Slide slide) {
        return (int) this.tags().stream().filter(tag -> slide.tags().contains(tag)).count();
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
