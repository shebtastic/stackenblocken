package _hashcode;

import _hashcode.models.Slide;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class Slideshow {
    private List<Slide> slides = new ArrayList<>();

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
        List<Slide> newSlides = new ArrayList<>();
        newSlides.add(slides.remove(0));
        Slide currentSlide = newSlides.get(0);
        int tolerance = 0;
        while (slides.size() > 1) {
            System.out.println(slides.size());
            for (int i = 0; i < slides.size() - 1; i++) {
                int distance = (currentSlide.TAGS.size() / 2) - currentSlide.tagOverlap(slides.get(i));
                if (distance == tolerance) {
                    newSlides.add(slides.remove(i));
                    tolerance = 0;
                }
            }
            tolerance++;
        }
        newSlides.add(slides.get(0));
        this.slides = newSlides;
    }

}
