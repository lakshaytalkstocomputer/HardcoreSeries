package exercises.brain.multiplication.challenge;

import exercises.brain.multiplication.user.User;
import org.springframework.stereotype.Service;

@Service
public class ChallengeServiceImpl implements ChallengeService {

    @Override
    public ChallengeAttempt verifyAttempt(ChallengeAttemptDTO resultAttempt) {

        // Check if the attempt is correct
        boolean isCorrect = resultAttempt.getGuess() == resultAttempt.getFactorA() * resultAttempt.getFactorB();

        // Forget the identifier
        User user = new User(null, resultAttempt.getUserAlias());

        // Build the domain object, Null id for now.
        ChallengeAttempt checkAttempt = new ChallengeAttempt(null,
                user,
                resultAttempt.getFactorA(),
                resultAttempt.getFactorB(),
                resultAttempt.getGuess(),
                isCorrect);

        return checkAttempt;
    }
}
