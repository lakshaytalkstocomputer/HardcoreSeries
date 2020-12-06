package exercises.brain.multiplication.challenge;

import lombok.Value;

/**
 * Attempt coming form the user.
 */
@Value
public class ChallengeAttemptDTO {
    int guess;
    int factorA, factorB;
    String userAlias;
}
