Feature: Test with background

    Background:
        Given now is "2020-04-05T06:07:08Z"

    Scenario: Add days
        Given someone adds 2 days to the clock

        Then the time is "2020-04-07T06:07:08Z"

    Scenario: Add months
        Given someone adds 1 month to the clock

        Then the time is "2020-05-05T06:07:08Z"
