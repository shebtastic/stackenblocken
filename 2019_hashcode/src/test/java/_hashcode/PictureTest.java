package _hashcode;

import _hashcode.models.Picture;
import _hashcode.models.VerticalPicture;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class PictureTest {

    private Picture underTest;

    private Picture pictureToTestWith;

    private final List<String> TAGS_1 = Arrays.asList("cat", "beach");
    private final List<String> TAGS_2 = Arrays.asList("dog", "beach");
    private final List<String> COMBINED_TAGS_1_2 = Arrays.asList("beach", "cat", "dog");

    @Before
    public void setUp() {
        underTest = new VerticalPicture(TAGS_1);
        pictureToTestWith = new VerticalPicture(TAGS_2);
    }

    @Test
    public void testIfCombinedLabelsAreCorrect() {
        //given
        List<String> expectedLabels = COMBINED_TAGS_1_2;

        //when
        List<String> actualLabels = underTest.combineTags(pictureToTestWith);

        //then
        assertEquals(actualLabels, expectedLabels);
    }

    @Test
    public void testIfOverlapIsCorrect() {
        //given
        int expectedOverlap = 1;

        //when
        int actualOverlap = underTest.tagOverlap(pictureToTestWith);

        //then
        assertEquals(actualOverlap, expectedOverlap);
    }
}
