package exercises.brain.multiplication.challenge;


public interface ChallengeService {
    /**
     *  Verfies if an attempt from the presentation layer is correct or not.
     * @param resultAttempt
     * @return the resulting challengeAttempt Object
     */
    ChallengeAttempt verifyAttempt(ChallengeAttemptDTO resultAttempt);
}
