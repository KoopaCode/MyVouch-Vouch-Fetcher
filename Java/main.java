import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;
import com.google.gson.Gson;
import java.io.FileReader;
import java.io.IOException;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

class Config {
    Vouch Vouch;

    static class Vouch {
        String MyVouch_URL;
        int Request_Delay;
    }
}

public class Main {

    public static void main(String[] args) {
        Gson gson = new Gson();
        try {
            Config config = gson.fromJson(new FileReader("config.json"), Config.class);
            printVouchesCount(config);
            ScheduledExecutorService executor = Executors.newScheduledThreadPool(1);
            executor.scheduleAtFixedRate(() -> printVouchesCount(config), 0, config.Vouch.Request_Delay, TimeUnit.SECONDS);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static void printVouchesCount(Config config) {
        try {
            Document doc = Jsoup.connect(config.Vouch.MyVouch_URL).get();
            Elements vouchesElement = doc.select("p.social span:last-child");
            String vouchesText = vouchesElement.text().trim();
            int vouchesCount = Integer.parseInt(vouchesText.replaceAll("\\D", ""));
            System.out.println("Vouch count: " + vouchesCount);
        } catch (IOException e) {
            System.err.println("Failed to fetch the vouch count: " + e.getMessage());
        }
    }
}
