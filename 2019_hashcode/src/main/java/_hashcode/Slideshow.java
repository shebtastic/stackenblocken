package _hashcode;

import _hashcode.models.Slide;

import java.util.ArrayList;
import java.util.List;

public class Slideshow {
    private List<Slide> slides = new ArrayList<>();
    
    public void addSlide(Slide slide) {
        slides.add(slide);
    }
}
