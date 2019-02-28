package _hashcode;

import _hashcode.models.HorizontalPicture;
import _hashcode.models.Slide;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class SlideshowTest {

    private Slideshow underTest;

    private final List<String> TAGS_1 = Arrays.asList("cat", "garden");
    private final List<String> TAGS_2 = Arrays.asList("garden", "selfie", "smile");
    private final List<String> TAGS_3 = Arrays.asList("dings", "bums", "haus", "baum", "katze", "schuh");
    private final List<String> TAGS_4 = Arrays.asList("dings", "bums", "eins", "zwei", "katze", "drei");

    private final Slide SLIDE_1 = new Slide(new HorizontalPicture(0, TAGS_1));
    private final Slide SLIDE_2 = new Slide(new HorizontalPicture(1, TAGS_2));
    private final Slide SLIDE_3 = new Slide(new HorizontalPicture(2, TAGS_3));
    private final Slide SLIDE_4 = new Slide(new HorizontalPicture(2, TAGS_4));

    @Before
    public void setUp() {
        underTest = new Slideshow();
    }

    @Test
    public void testForCorrectInterestScore() {
        //given
        int expectedScore = 1;
        underTest.addSlide(SLIDE_1);
        underTest.addSlide(SLIDE_2);

        //when
        int actualScore = underTest.interestScore();

        //then
        assertEquals(expectedScore, actualScore);
    }

    @Test
    public void testForSecondInterestScore() {
        //given
        int expectedScore = 3;
        underTest.addSlide(SLIDE_3);
        underTest.addSlide(SLIDE_4);

        //when
        int actualScore = underTest.interestScore();

        //then
        assertEquals(expectedScore, actualScore);
    }

}
