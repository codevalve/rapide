---
description: Pull latest GitHub issues and evaluate them against project philosophy.
---

# Issue Triage & Engineering Workflow

This workflow is designed to automate the process of reviewing new issues or feature requests for Rapide.

## 1. Retrieve Latest Issues
// turbo
- [ ] Run `gh issue list --limit 10` to see the most recent activity.
- [ ] For each new ID, run `gh issue view <id>` to understand the full context.

## 2. Philosophical Evaluation
Against **AGENTS.md**:
- [ ] Does this request add unnecessary dependencies? (Minimalist)
- [ ] Is it "Local First" by default? (Private)
- [ ] If it's a TUI change, does it follow the "Charm" aesthetic? (Lipgloss usage)

Against **ROADMAP.md**:
- [ ] Should this feature be slated for v2.7 (Hashtags/Links) or v3.0 (MCP)?
- [ ] Does it align with the core journal vision?

## 3. Engineering Plan
For each triage action:
- [ ] Draft a `plan.md` artifact with:
    - **Context**: Summary of the issue.
    - **Technical Approach**: Specific functions in `internal/` or `cmd/` to modify.
    - **Test Cases**: New test requirements for `parser_test.go` or similar.
- [ ] Respond to the issue on GitHub using `gh issue comment <id> -b "..."` with the proposed plan.

## 4. Work Dispatch
- [ ] Update `ROADMAP.md` if the issue is high-priority for the next release.
- [ ] Ask the USER to approve the plan before execution.
- [ ] **Closing**: After merging the fix to `main`, close the issue using `gh issue close <id> --comment "Fixed in vX.Y.Z"`.

## 5. Environment Check
// turbo
- [ ] Ensure local git config is correct:
      `git config user.email "john.lovell@codevalve.com" && git config user.name "codevalve"`
