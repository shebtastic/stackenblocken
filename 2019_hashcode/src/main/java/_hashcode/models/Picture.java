package _hashcode.models;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public abstract class Picture {
    public final int INDEX;
    public final String ORIENTATION;
    public final List<String> TAGS;

    public Picture(int index, String orientation, List<String> tags) {
        this.INDEX = index;
        this.ORIENTATION = orientation;
        this.TAGS = tags;
    }

    public List<String> combineTags(Picture picture) {
        List<String> combinedSortedTags;
        Set<String> set = new HashSet<>();

        set.addAll(this.TAGS);
        set.addAll(picture.TAGS);

        combinedSortedTags = new ArrayList<>(set);
        combinedSortedTags.sort(String::compareTo);
        return combinedSortedTags;
    }

    public int tagOverlap(Picture picture) {
        return Math.abs(
            (this.TAGS.size() + picture.TAGS.size())
            - (combineTags(picture).size())
        );
    }
}
