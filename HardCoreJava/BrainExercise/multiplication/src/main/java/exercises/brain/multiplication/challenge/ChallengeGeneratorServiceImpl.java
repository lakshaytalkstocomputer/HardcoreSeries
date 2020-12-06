package exercises.brain.multiplication.challenge;

import org.springframework.stereotype.Service;
import java.util.Random;

/**
 *  This class will implement the ChallengeGeneratorService Interface to generate challenges;
 */
@Service
public class ChallengeGeneratorServiceImpl implements ChallengeGeneratorService{

    private final Random random;

    ChallengeGeneratorServiceImpl() {
        this.random = new Random();
    }

    protected ChallengeGeneratorServiceImpl(final Random random){
        this.random = random;
    }

    @Override
    public Challenge randomChallenge() {
        return null;
    }
}
