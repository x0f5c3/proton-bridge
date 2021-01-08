Feature: Import from EML files
  Background:
    Given there is connected user "user"
    And there is "user" with mailbox "Folders/Foo"
    And there is "user" with mailbox "Folders/Bar"
    And there are EML files
      | file              | from            | to                        | subject | time                |
      | Foo/one.eml       | foo@example.com | bridgetest@protonmail.com | one     | 2020-01-01T12:00:00 |
      | Foo/two.eml       | bar@example.com | bridgetest@protonmail.com | two     | 2020-01-01T13:00:00 |
      | Sub/Foo/three.eml | bar@example.com | bridgetest@protonmail.com | three   | 2020-01-01T12:30:00 |
    And there is EML file "Inbox/hello.eml"
      """
      Subject: hello
      From: Bridge Test <bridgetest@pm.test>
      To: Internal Bridge <test@protonmail.com>
      Received: by 2002:0:0:0:0:0:0:0 with SMTP id 0123456789abcdef; Wed, 30 Dec 2020 01:23:45 0000

      hello

      """

  Scenario: Import all
    When user "user" imports local files
    Then progress result is "OK"
    And transfer exported 4 messages
    And transfer imported 4 messages
    And transfer failed for 0 messages
    And API mailbox "INBOX" for "user" has messages
      | from               | to                  | subject |
      | bridgetest@pm.test | test@protonmail.com | hello   |
    And API mailbox "Folders/Foo" for "user" has messages
      | from            | to                        | subject |
      | foo@example.com | bridgetest@protonmail.com | one     |
      | bar@example.com | bridgetest@protonmail.com | two     |
      | bar@example.com | bridgetest@protonmail.com | three   |

  Scenario: Import only Foo to Bar with time limit
    When user "user" imports local files with rules
      | source | target | from                | to                  |
      | Foo    | Bar    | 2020-01-01T12:10:00 | 2020-01-01T13:00:00 |
    Then progress result is "OK"
    And transfer exported 2 messages
    And transfer imported 2 messages
    And transfer failed for 0 messages
    And API mailbox "Folders/Bar" for "user" has messages
      | from            | to                        | subject |
      | bar@example.com | bridgetest@protonmail.com | two     |
      | bar@example.com | bridgetest@protonmail.com | three   |

  Scenario: Import broken EML message
    Given there is EML file "Broken/broken.eml"
      """
      Content-type: multipart/mixed
      """
    When user "user" imports local files with rules
      | source | target |
      | Broken | Foo    |
    Then progress result is "OK"
    And transfer exported 1 messages
    And transfer imported 0 messages
    And transfer failed for 1 messages
