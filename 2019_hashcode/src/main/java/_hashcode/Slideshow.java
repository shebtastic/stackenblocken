package _hashcode;

import _hashcode.models.Slide;

import java.util.ArrayList;
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
}
