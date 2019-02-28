package _hashcode.models;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

public abstract class Picture {
    public final int INDEX;
    public final String ORIENTATION;
    public final Set<Integer> TAGS;

    public Picture(int index, String orientation, Set<Integer> tags) {
        this.INDEX = index;
        this.ORIENTATION = orientation;
        this.TAGS = tags;
    }

    public Set<Integer> combinedTags(Picture picture) {
        Set<Integer> set = new HashSet<>();
        set.addAll(this.TAGS);
        set.addAll(picture.TAGS);
        return set;
    }

    @Override
    public String toString() {
        return TAGS.stream().map(Object::toString).collect(Collectors.joining(" "));
    }

    public int tagOverlap(Picture picture) {
        return Math.abs(
            (this.TAGS.size() + picture.TAGS.size())
            - (combinedTags(picture).size())
        );
    }

    public String toOutput() {
        return Integer.toString(INDEX);
    }
}
