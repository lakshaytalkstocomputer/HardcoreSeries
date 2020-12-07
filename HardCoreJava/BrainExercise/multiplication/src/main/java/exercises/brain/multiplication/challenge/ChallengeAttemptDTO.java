package exercises.brain.multiplication.challenge;

import lombok.Value;

import javax.validation.constraints.Max;
import javax.validation.constraints.Min;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Positive;

/**
 * Attempt coming form the user.
 */
@Value
public class ChallengeAttemptDTO {
    @Positive(message = "this neeeds to be positive mate!")
    int guess;
    @Min(1) @Max(99)
    int factorA, factorB;
    @NotBlank
    String userAlias;
}
