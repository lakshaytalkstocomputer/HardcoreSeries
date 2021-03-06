package exercises.brain.multiplication.challenge;

import exercises.brain.multiplication.user.User;
import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

/**
 *  This class identifies an attempt from a User to solve a Challenge.
 */
@EqualsAndHashCode
@Getter
@ToString
@AllArgsConstructor
public class ChallengeAttempt {
    private Long id;
    private User user;
    private int factorA;
    private int factorB;
    private int resultAttempt;
    private boolean correct;
}
