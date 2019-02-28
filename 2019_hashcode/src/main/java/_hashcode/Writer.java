package _hashcode;

import _hashcode.models.Slide;

import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;
import java.util.List;

public class Writer {
    public static void write(List<Slide> slides, String name) {
        BufferedWriter writer;

        try {
            writer = new BufferedWriter(new FileWriter("outputs/" + name));
            writer.write(Integer.toString(slides.size()));
            writer.newLine();

            for (int i=0; i<slides.size(); i++) {
                writer.write(slides.get(i).toString());
                writer.newLine();
            }

            writer.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}
