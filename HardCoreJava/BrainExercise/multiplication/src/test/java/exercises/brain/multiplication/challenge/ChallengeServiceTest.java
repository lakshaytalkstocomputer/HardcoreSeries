package exercises.brain.multiplication.challenge;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;
public class ChallengeServiceTest {

    private ChallengeService challengeService;

    @BeforeEach
    public void setUp(){
        challengeService = new ChallengeServiceImpl();
    }

    @Test
    public void checkCorrectAttemptTest(){
        // given
        ChallengeAttemptDTO attemptDTO = new ChallengeAttemptDTO(3000, 50, 60, "john_doe");

        // when
        ChallengeAttempt resultAttempt = challengeService.verifyAttempt(attemptDTO);

        // then
        then(resultAttempt.isCorrect()).isTrue();
    }

    @Test
    public void checkWrongAttemptTest(){
        // given
        ChallengeAttemptDTO attemptDTO = new ChallengeAttemptDTO(1000, 50, 60, "john_doe");

        // when
        ChallengeAttempt resultAttempt = challengeService.verifyAttempt(attemptDTO);

        // then
        then(resultAttempt.isCorrect()).isFalse();

    }
}
