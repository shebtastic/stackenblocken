package _hashcode;

import _hashcode.models.HorizontalPicture;
import _hashcode.models.Picture;
import _hashcode.models.Slide;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class SlideTest {

    private Slide underTest;

    private Slide slideToTestWith;

    private final List<String> TAGS_1 = Arrays.asList("cat", "garden");
    private final List<String> TAGS_2 = Arrays.asList("garden", "selfie", "smile");

    private final HorizontalPicture PICTURE_1 = new HorizontalPicture(0, TAGS_1);
    private final HorizontalPicture PICTURE_2 = new HorizontalPicture(1, TAGS_2);

    @Before
    public void setUp() {
        underTest = new Slide(PICTURE_1);
        slideToTestWith = new Slide(PICTURE_2);
    }

    @Test
    public void testForCorrectOverlap() {
        //given
        int expectedOverlap = 1;

        //when
        int actualOverlap = underTest.tagOverlap(slideToTestWith);

        //then
        assertEquals(expectedOverlap, actualOverlap);
    }
}
