package _hashcode.models;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Slide {
    public final List<Picture> PICTURES;
    public final Set<Integer> TAGS;

    public Slide(VerticalPicture verticalPicture, VerticalPicture secondVerticalPicture) {
        this.PICTURES = Arrays.asList(verticalPicture, secondVerticalPicture);
        this.TAGS = PICTURES.get(0).combinedTags(PICTURES.get(1));
    }

    public Slide(HorizontalPicture horizontalPicture) {
        this.PICTURES = Arrays.asList(horizontalPicture);
        this.TAGS = PICTURES.get(0).TAGS;
    }

    public int tagOverlap(Slide slide) {
        Set<Integer> set = new HashSet<>(TAGS);
        set.removeAll(slide.TAGS);
        return TAGS.size() - set.size();

    }

    @Override
    public String toString() {
        String result = PICTURES.stream()
                .map(Picture::toOutput)
                .collect(Collectors.joining(" "));
        result += " TAGS: ";
        result += TAGS.stream().map(tag -> tag.toString()).collect(Collectors.joining(" "));
        return result;
    }

    public String toOutput() {
        return PICTURES.stream()
                .map(Picture::toOutput)
                .collect(Collectors.joining(" "));
    }
}
