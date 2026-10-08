Feature: F1 Greeting

  @story-F1.1
  Rule: Requesting a greeting with a name returns that greeting

    Scenario: Greeting a named caller
      When an API caller requests a greeting for "Ada"
      Then the response is a JSON greeting of "Hello, Ada!"

  @story-F1.2
  Rule: A missing or empty name is refused

    @negative
    Scenario: No name is given
      When an API caller requests a greeting without a name
      Then the request is refused with a 400 error

    @negative
    Scenario: An empty name is given
      When an API caller requests a greeting with an empty name
      Then the request is refused with a 400 error
