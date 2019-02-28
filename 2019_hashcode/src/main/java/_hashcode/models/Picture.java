package _hashcode.models;

public abstract class Picture {
    public final String ORIENTATION;
    public final String[] TAGS;

    public Picture(String orientation, String[] tags) {
        this.ORIENTATION = orientation;
        this.TAGS = tags;
    }

    public String[] combineTags(String[] tags) {
        return null;
    }

    public int tagOverlap(Picture picture) {
        return Math.abs(this.TAGS.length - picture.TAGS.length);
    }
}
