Feature: Without Background

    Scenario: Set time
        Given the clock is at "2020-01-02T03:04:05Z"

        Then the time is "2020-01-02T03:04:05Z"

        Given the clock is set to "2020-02-03T04:05:06Z"

        Then the time is "2020-02-03T04:05:06Z"

        Given Someone sets the clock to "2020-03-04T05:06:07Z"

        Then the time is "2020-03-04T05:06:07Z"

        Given now is "2020-04-05T06:07:08Z"

        Then the time is "2020-04-05T06:07:08Z"

    Scenario: Add time
        Given the clock is at "2020-01-02T03:04:05Z"
        And someone adds 1h5s to the clock

        Then the time is "2020-01-02T04:04:10Z"

        Given someone adds 2 days to the clock

        Then the time is "2020-01-04T04:04:10Z"

        Given someone adds 1 month to the clock

        Then the time is "2020-02-04T04:04:10Z"

        Given someone adds 3 years to the clock

        Then the time is "2023-02-04T04:04:10Z"

        Given someone adds 1 month 2 days to the clock

        Then the time is "2023-03-06T04:04:10Z"

        Given someone adds 2 year 3 days to the clock

        Then the time is "2025-03-09T04:04:10Z"

        Given someone adds 3 year 4 months to the clock

        Then the time is "2028-07-09T04:04:10Z"

        Given someone adds 4 year 5 months 6 days to the clock

        Then the time is "2032-12-15T04:04:10Z"

    Scenario: Freeze and Unfreeze
        Given the time is now

        When I freeze the clock
        And I wait for 50ms

        Then the time is not now

        When I wait for 50ms
        And I release the clock

        Then the time is now
