## TODO
- Implement errors.* function in Error interface
- Add WithError() extension
- Support nested error
- Add support for JSON (marshal,unmarshal, etc)

## Severity and priority concept
The concepts of severity and priority are often used in error handling, issue tracking, and incident management, but they refer to different aspects of an issue. Understanding the difference between these two terms is important for effectively managing and addressing problems in software or systems.

---

1. Severity

- Definition: Severity refers to the impact or seriousness of the issue on the system or application. It describes how critical the error is in terms of its effect on the functionality of the system or the user experience.

- Focus: The focus of severity is how bad the issue is in terms of technical or operational impact.
Common Metrics: Severity is often categorized into levels like Critical, Major, Minor, or Trivial (or, as in our previous discussion, Warning, Recoverable, Critical, Fatal, Panic).

- Example:
Critical severity: A database failure that makes the application unavailable.
Minor severity: A small UI glitch that doesn't affect functionality.

---

2. Priority

- Definition: Priority refers to the urgency or order of importance for addressing the issue. It indicates how soon the issue should be fixed or resolved relative to other tasks, usually based on business needs or customer impact.

- Focus: The focus of priority is how soon the issue needs to be addressed, often driven by factors like customer expectations, deadlines, or resource availability.
Common Metrics: Priority is often categorized into levels such as High, Medium, or Low (sometimes with additional gradation like Urgent).

- Example:
High priority: A critical bug preventing users from logging in.
Low priority: A minor cosmetic issue that does not affect functionality or user experience.

## Key Differences Between Severity and Priority

| Aspect             | Severity                                                                                | Priority                                                                                                          |
| ------------------ | --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| Focus              | The impact or seriousness of the issue                                                  | The urgency or timeliness for resolving the issue                                                                 |
| How it is Measured | Based on technical impact (e.g., system failure, performance degradation)               | Based on business context, customer needs, or team resource availability                                          |
| Example            | A system crash is high severity because it halts all operations.                        | A cosmetic UI bug is low priority because it does not impede workflow, but it should still be fixed eventually.   |
| Impact             | Severity determines how much the issue affects functionality or the system’s stability. | Priority determines how quickly the issue needs to be addressed based on business needs or customer expectations. |
| Who Decides        | Typically determined by technical teams (developers, testers, etc.).                    | Typically determined by product managers, business teams, or support teams.                                       |


## Relation Between Severity and Priority

While severity and priority are different, they are related, and the relationship can guide the decision-making process in managing issues. Here’s how they are related:

- High Severity, High Priority:

If an issue has high severity (e.g., the system crashes), it typically also has high priority because it significantly impacts the functionality and requires an immediate fix.
Example: A database server going down is both a high severity and high priority issue, and it requires urgent resolution to restore the system's availability.

- High Severity, Low Priority:

Occasionally, an issue can have high severity but low priority if it does not need to be addressed immediately or if it affects a small portion of users or is less critical in the short term.
Example: A critical bug affecting a minor feature used by only a few users could be high severity but low priority, meaning it’s important but doesn’t need to be fixed immediately.

- Low Severity, High Priority:

Some issues may have low severity (e.g., cosmetic or non-functional issues) but still have high priority if they need to be addressed quickly due to business reasons, user experience, or customer complaints.
Example: A minor UI issue on a promotional landing page could be low severity but high priority because the company is running a marketing campaign, and a clean, polished look is essential for the campaign’s success.

- Low Severity, Low Priority:

These are the issues that are neither critical nor urgent. They have low impact and can be addressed at a later time or when resources permit.
Example: A typo on a rarely visited page could be low severity and low priority, and while it should eventually be fixed, it’s not urgent.

### Summary

Severity is about how bad an issue is (impact on the system or functionality).
Priority is about how soon the issue needs to be addressed (urgency based on business or customer needs).
Relation: While high severity usually leads to high priority, the two are not always aligned. An issue can have high severity but low priority if it’s not immediately urgent, or low severity but high priority if it needs to be fixed quickly for business reasons.
In summary, severity affects the impact on the system or user, while priority affects the urgency of fixing it. They work together to determine the most appropriate response to an issue in terms of addressing it both technically and business-wise.