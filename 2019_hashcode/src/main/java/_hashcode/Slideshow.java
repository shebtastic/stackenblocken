package _hashcode;

import _hashcode.models.Slide;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.*;

public class Slideshow {
    private static final Logger LOGGER = LoggerFactory.getLogger(Slideshow.class);

    private List<Slide> slides = new ArrayList<>();
    private final int THREAD_COUNT = 6;

    public void addSlide(Slide slide) {
        slides.add(slide);
    }

    public int interestScore() {
        int score = 0;

        if (slides.size() < 2) return 0;

        Slide currentSlide, nextSlide;
        for (int index = 0; index < slides.size() - 1; index++) {
            currentSlide = slides.get(index);
            nextSlide = slides.get(index + 1);

            score += Math.min(
                Math.min(
                    currentSlide.TAGS.size() - currentSlide.tagOverlap(nextSlide),
                    currentSlide.tagOverlap(nextSlide)
                ),
                nextSlide.TAGS.size() - currentSlide.tagOverlap(nextSlide)
            );
        }

        return score;
    }

    public void bringInOrder() {
        long beginTime = System.currentTimeMillis();
        int beginSlideSize = this.getSlides().size();
        LOGGER.debug("\n\nbring in Order:");
        for(Slide slide: this.getSlides()) {
            LOGGER.debug(slide.toString());
        }
        ExecutorService executor = Executors.newFixedThreadPool(THREAD_COUNT);
        List<Slide> newSlides = new ArrayList<>();
        newSlides.add(slides.remove(0));
        Slide currentSlide = newSlides.get(0);
        float tolerance = 0; // 10
        while (slides.size() > 1) {
            LOGGER.debug("----");
            int slideSize = slides.size();
            int slidesPerThread = Math.round(slideSize / THREAD_COUNT);
            LOGGER.debug("size: " + slideSize + " slidesPerThread: " + slidesPerThread + " tolerance: " + tolerance);
            SlidesSearchThread[] threads = new SlidesSearchThread[THREAD_COUNT];
            int fromIndex = 0;
            int toIndex = slidesPerThread;
            for (int i = 0; i < THREAD_COUNT; i++) {
                threads[i] = new SlidesSearchThread(currentSlide, Math.round(tolerance), slides.subList(fromIndex, toIndex), fromIndex);
                fromIndex = toIndex;
                toIndex = (i < THREAD_COUNT-2) ? fromIndex + slidesPerThread : slideSize;
            }
            int foundIndex = -1;

            try {
                foundIndex = executor.invokeAny(Arrays.asList(threads));
            } catch (InterruptedException e) {
                e.printStackTrace();
            } catch (ExecutionException e) {
                LOGGER.debug("no result found");
            }
            if (foundIndex >= 0) {
                LOGGER.debug("foundIndex: " + foundIndex);
                tolerance = -1;
                newSlides.add(slides.remove(foundIndex));
                currentSlide = newSlides.get(newSlides.size() - 1);
                if ((slides.size() % 10000) == 0) {
                    float slidesDone = beginSlideSize - slides.size();
                    float timePassed = System.currentTimeMillis() - beginTime;
                    float slidesPerMin = slidesDone / timePassed * 100 * 60;
                    LOGGER.info(slides.size() + " slides left (" + slidesPerMin + "/m)");
                }
            }


            tolerance++;
        }

        try {
            executor.shutdownNow();
            executor.awaitTermination(500, TimeUnit.MILLISECONDS);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        newSlides.add(slides.get(0));
        this.slides = newSlides;
        float timePassed = System.currentTimeMillis() - beginTime;
        LOGGER.info("ordered in " + timePassed + " ms");
    }

    public List<Slide> getSlides() {
        return slides;
    }
}
