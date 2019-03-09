package _hashcode;

import _hashcode.models.Slide;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.concurrent.Callable;
public class SlidesSearchThread implements Callable<Integer> {

    private static final Logger LOGGER = LoggerFactory.getLogger(SlidesSearchThread.class);

    private Slide CURRENT_SLIDE;
    private int TOLERANCE;
    private int FROM_INDEX;

    private List<Slide> SLIDES;

    public SlidesSearchThread(Slide currentSlide, int tolerance, List<Slide> slides, int fromIndex) {
        CURRENT_SLIDE = currentSlide;
        TOLERANCE = tolerance;
        SLIDES = slides;
        FROM_INDEX = fromIndex;
        LOGGER.debug("thread for " + slides.size() + " slides from " + fromIndex + " to " + (fromIndex+slides.size()-1) + " running");
    }

    @Override
    public Integer call() throws Exception {
        int result = -1;
        for (int i=0; i<SLIDES.size(); i++) {
            int distance = (CURRENT_SLIDE.TAGS.size() / 2) - CURRENT_SLIDE.tagOverlap(SLIDES.get(i));
            if (distance <= TOLERANCE) {
                result = i;
                break;
            }
        }


        if (result >= 0) {
            LOGGER.debug("thread from " + FROM_INDEX + " found: " + (result + FROM_INDEX) + " (intern: " + result + ")");
            return (result + FROM_INDEX);
        } else {
            throw new Exception("not found");
        }
    }
}
