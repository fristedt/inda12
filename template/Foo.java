/**
 * Typical comments describing this class
 * @author Snava Sne and his extended family
 * @version 2042.42.42
 */
public class Foo
{
    private int power;
    /**
     * Här kan man skriva nästan vad som helst, notera att även åäö funkar!
     */
    public Foo()
    {
		power = Integer.MAX_VALUE;
    }
    
    /** Inline comment.*/
    public void empty(){
        power = 0;
    }
    /** Visar att även kod wrappas om raderna är riktigt jäkla långa. Varför du nu borde behöva detta är en annan fråga....*/
    public void reallyLongLine() {
        power = Integer.MIN_VALUE + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1;
    }
    // This is a very long line which should be forced to wrap if this thing is set up correctly. I really hope it works because whenever code escapes and remains outside the page it's really hard to read.
    public static void main(String[] args) {
		System.out.println("Räksmörgås!");
    }
}
