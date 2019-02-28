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

    private final Slide SLIDE_1 = new Slide(new HorizontalPicture(0, TAGS_1));
    private final Slide SLIDE_2 = new Slide(new HorizontalPicture(1, TAGS_2));

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

}
