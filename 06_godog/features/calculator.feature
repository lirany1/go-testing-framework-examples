Feature: Calculator Operations
  As a user of the calculator
  I want to perform basic arithmetic operations
  So that I can compute results accurately

  Background:
    Given a calculator

  Scenario: Add two positive numbers
    Given I have entered 2 into the calculator
    And I have entered 3 into the calculator
    When I press add
    Then the result should be 5 on the screen

  Scenario: Add negative numbers
    Given I have entered -5 into the calculator
    And I have entered -3 into the calculator
    When I press add
    Then the result should be -8 on the screen

  Scenario: Multiply numbers
    Given I have entered 4 into the calculator
    And I have entered 5 into the calculator
    When I press multiply
    Then the result should be 20 on the screen

  Scenario: Divide numbers
    Given I have entered 10 into the calculator
    And I have entered 2 into the calculator
    When I press divide
    Then the result should be 5 on the screen

  Scenario: Division by zero returns error
    Given I have entered 10 into the calculator
    And I have entered 0 into the calculator
    When I press divide
    Then I should see an error message "cannot divide by zero"

  Scenario Outline: Calculate with different values
    Given I have entered <first> into the calculator
    And I have entered <second> into the calculator
    When I press <operation>
    Then the result should be <result> on the screen

    Examples:
      | first | second | operation | result |
      | 1     | 1      | add       | 2      |
      | 10    | 5      | subtract  | 5      |
      | 3     | 7      | multiply  | 21     |
      | 20    | 4      | divide    | 5      |
      | 0     | 5      | add       | 5      |
      | 100   | 10     | divide    | 10     |
