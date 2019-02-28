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
                    currentSlide.tags().size() - currentSlide.tagOverlap(nextSlide),
                    currentSlide.tagOverlap(nextSlide)
                ),
                nextSlide.tags().size() - currentSlide.tagOverlap(nextSlide)
            );
        }

        return score;
    }

    private void sort() {
        Collections.sort(slides, new Comparator<Slide>() {
            public int compare(Slide s1, Slide s2) {
                return Integer.compare(s1.tags().size(), s2.tags().size());
            }
        });
    }

    public void bringInOrder() {
        this.sort();
        List<Slide> newSlides = new ArrayList<>();
        for (int i=0; i<this.slides.size() / 2; i+=2) {
            newSlides.add(slides.get(i));
            newSlides.add(slides.get(slides.size() - (i + 1)));
        }
        this.slides = newSlides;
    }

}
