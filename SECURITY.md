# Security Policy

## Supported Versions

We are currently supporting the following versions with security updates:

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability, please report it by creating a private security advisory through GitHub's Security tab.

**Please do not report security vulnerabilities through public GitHub issues.**

We will respond to your report within 48 hours and will keep you informed about the progress towards a fix and announcement.

## Security Features

This repository has the following security features configured:

### Automated Security Scanning

- **Dependabot**: Automatically checks for vulnerable dependencies and creates pull requests to update them
- **CodeQL Analysis**: Advanced semantic code analysis for security vulnerabilities
- **OSV-Scanner**: Scans for known vulnerabilities in open source dependencies
- **Dependency Review**: Blocks pull requests that introduce vulnerable dependencies

### Code Review

- **CODEOWNERS**: Enforces code review requirements from designated owners
- **Branch Protection**: Should be enabled on the default branch (requires repository admin configuration)

### Additional Security Measures

The following security features should be enabled in repository settings:

1. **Branch Protection Rules** - Protect the default branch
   - Require pull request reviews before merging
   - Require status checks to pass
   - Require signed commits
   - Include administrators in restrictions

2. **GitHub Advanced Security** - Enable for comprehensive security scanning
   - Code scanning with CodeQL
   - Secret scanning
   - Dependency review

3. **Secret Scanning** - Detect exposed credentials and secrets

4. **Signed Commits** - Require GPG/SSH signing for commits to prevent impersonation

## Security Best Practices

When contributing to this repository:

1. Keep dependencies up to date
2. Never commit secrets, credentials, or sensitive data
3. Sign your commits with GPG or SSH keys
4. Run security scans locally before pushing code
5. Follow secure coding practices
6. Review dependency updates carefully

## Security Contacts

For security-related questions or concerns, please contact the repository maintainers through GitHub.
