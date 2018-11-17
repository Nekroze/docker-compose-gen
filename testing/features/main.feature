Feature: Root command

    Scenario: Can get help and usage information
        When I run `docker-compose-gen --help`

        Then it should pass with "Usage:"
