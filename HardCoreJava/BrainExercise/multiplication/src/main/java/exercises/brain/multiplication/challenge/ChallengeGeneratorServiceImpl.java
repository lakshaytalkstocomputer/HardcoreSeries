package exercises.brain.multiplication.challenge;

import org.springframework.stereotype.Service;
import java.util.Random;

/**
 *  This class will implement the ChallengeGeneratorService Interface to generate challenges;
 */
@Service
public class ChallengeGeneratorServiceImpl implements ChallengeGeneratorService{

    private final Random random;
    private final static int MINIMUM_FACTOR = 11;
    private final static int MAXIMUM_FACTOR = 100;

    ChallengeGeneratorServiceImpl() {
        this.random = new Random();
    }

    protected ChallengeGeneratorServiceImpl(final Random random){
        this.random = random;
    }

    private int next(){
        return random.nextInt(MAXIMUM_FACTOR - MINIMUM_FACTOR) + MINIMUM_FACTOR;
    }
    @Override
    public Challenge randomChallenge() {
         return new Challenge(next(),next());
    }
}
