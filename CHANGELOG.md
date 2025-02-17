# Changelog

## [v1.4.0] - 2023-06-15

**Important note:** This release require a study-service version >= 1.6.0 to work properly with participant messages.

### Added

- New env variable `GLOBAL_EMAIL_TEMPLATE_CONSTANTS_JSON` can be used to define global constants that can be used in email templates via a JSON file. This is useful to define global variables that are used in multiple email templates. The content of this file must be a valid JSON string. Example: `{"baseUrl": "https://example.com"}`

### Changed

- Add instance ID in logging information about generating auto messages.

- Before participant messages are generated, it is checked if there are any studies with pending participant messages to be sent now. If not, generating participant messages is skipped.

## [v1.3.2] - 2023-03-31

### Changed

- If a schedule is uploaded that contains invalid value for the `until` field, the schedule is not uploaded and an error is returned.
- Logging provides information about starting and finishing processes of generating and sending participant messages, researcher notifications, auto messages, low and high prio emails. Each process can be identified by its own thread ID.

## [v1.3.1] - 2023-02-22

### Changed

- When generating participant messages, the `profileId` field is also available in the template data. This can be used to generate a link to open a survey for a specific participant.

## [v1.3.0] - 2023-01-30

### BREAKING CHANGE

- [PR #12](https://github.com/influenzanet/messaging-service/pull/12)
  - Scheduled messages for participants are now fetched and sent regularly with specified frequency for all studies. Researcher notifications are also fetched and sent with specified frequency. Uploading auto message schedules for these message types is omitted. Existing schedules for these message types should be removed.
    - New env variables `MESSAGE_SCHEDULER_INTERVAL_PARTICIPANT_MESSAGE`, `MESSAGE_SCHEDULER_INTERVAL_RESEARCHER_NOTIFICATION` set the interval period of fetching participant messages and researcher notifications in seconds. An interval value of 0 indicates no query of the respective message type.
  - The next scheduled time for auto messages is ensured to be a valid future date in order to prevent them to be sent multiple times.
  - A new object ID is generated for emails moved from outgoing to send to prevent issues where messages could not be marked as sent.
  - Checks are added to prevent empty email templates and empty list of translations.

### Added

- [PR #10](https://github.com/influenzanet/messaging-service/pull/10) Checks if go template engine is able to parse and execute new email templates. This prevents uploading incorrect email templates.
- [PR #11](https://github.com/influenzanet/messaging-service/pull/11) Improvement on log messages and raise error when no instances are found to help discover issues during initial setup.
- Skip users with account types other than "email", and produce DEBUG log for them instead of errors.

## [v1.2.0] - 2022-10-06

### Added

- Add email-emulator service that will write emails onto the disk instead of sending them to an email server. This is a simple alternative to perform local tests without the need to setup actual email server (in case message sending is not needed).

### Changed

- Participant message generation will use the payload (participant flags) for the email template, so that these can be utilised in emails.

## [v1.1.1] - 2022-09-01

### Changed

- Replacing log.Print instances with custom logger to use log levels.
- Fixing issue, where participant messages did not inlcude a login token.

## [v1.1.0] - 2022-06-03

### Added

- New message type / message sending logic for researcher notifications. This messages can be generated through study rules to send a notification about specific topics to a specified list email addresses.

### Changed

- Updated dependencies (gRPC, study-service), and made necessary adaptations on the Makefile to be able to generate the new api files.

## [v1.0.0] - 2022-03-08

### Added

- New message type for participant messages.

### Changed

- Using new logger library with improved logging format and configurable log level.
  - For message-scheduler use the environment variable `LOG_LEVEL=<level>`. Valid values are `debug`, `info`, `warning`, `error` . Default (if not speficied) is `info`.

## [v0.9.3] - 2021-07-28

### Changed

- For newsletter message type, the weekday setting of the user can be ignored.
- API arguments for send messages to all users and study participants extended to use "IgnoreWeekday" (boolean), to control if for newsletter type, the filter should ignore reminder weekday of the user.

## [v0.9.2]

### Added

- Email templates can use the language attribute that would contain the preferred language from the user model. Example usage added to the [docs](docs/email-templates.md).

### Changed

- "Auto email" definitions can contain a label, so that admins can describe the intent for the specific config.
- Updated dependencies (reflected in go.mod).
- Email-templates documentation includes new possibilities related to the above changes of this release.
