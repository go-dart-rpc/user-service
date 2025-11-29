# Security Setup Guide for Repository Administrators

This guide provides step-by-step instructions for configuring security settings to resolve all security audit issues (#17).

## Overview

This repository has 7 security audit issues that need to be addressed:
- ✅ CODEOWNERS file (Fixed - file created)
- ✅ Dependency vulnerabilities (Fixed - dependencies updated)
- ✅ Dependency Review workflow (Fixed - workflow added)
- ⚠️ Branch Protection (Requires admin configuration)
- ⚠️ Code Scanning (Requires admin configuration)
- ⚠️ Secret Scanning (Requires admin configuration)
- ⚠️ Signed Commits (Requires admin configuration)

## Steps to Complete Setup

### 1. Enable Branch Protection (Issue #9)

1. Navigate to **Settings** → **Branches**
2. Click **Add branch protection rule**
3. In "Branch name pattern", enter: `main` (or your default branch name)
4. Configure the following settings:
   - ✅ **Require a pull request before merging**
     - ✅ Require approvals: `1`
     - ✅ Dismiss stale pull request approvals when new commits are pushed
     - ✅ Require review from Code Owners
   - ✅ **Require status checks to pass before merging**
     - ✅ Require branches to be up to date before merging
     - Add status checks: `CodeQL`, `Dependency Review`, `OSV-Scanner`
   - ✅ **Require signed commits**
   - ✅ **Require linear history** (optional but recommended)
   - ✅ **Include administrators**
5. Click **Create** or **Save changes**

### 2. Enable GitHub Advanced Security (Issues #11, #12)

**Note**: This feature requires GitHub Advanced Security, which is:
- Free for public repositories
- Requires a license for private repositories

#### For Public Repositories:
1. Navigate to **Settings** → **Code security and analysis**
2. Enable the following features:
   - ✅ **Code scanning**
     - Click **Set up** → **Default** for CodeQL analysis
     - Or use **Advanced** to use the existing `.github/workflows/codeql.yml`
   - ✅ **Secret scanning**
     - Click **Enable**
   - ✅ **Secret scanning push protection** (recommended)
     - Click **Enable**

#### For Private Repositories:
1. Ensure your organization has GitHub Advanced Security licenses
2. Navigate to **Settings** → **Code security and analysis**
3. Enable **GitHub Advanced Security**
4. Then enable the features listed above

### 3. Configure Code Scanning (Issue #11)

The repository already has a CodeQL workflow at `.github/workflows/codeql.yml`. To activate it:

1. Navigate to **Settings** → **Code security and analysis**
2. Under **Code scanning**:
   - If using default setup: Click **Set up** → **Default**
   - If using advanced setup: Ensure the workflow is enabled
3. Configure the workflow to run on:
   - Pull requests to the default branch
   - Push events to the default branch
   - Weekly schedule (already configured)

### 4. Enable Secret Scanning (Issue #12)

1. Navigate to **Settings** → **Code security and analysis**
2. Under **Secret scanning**:
   - Click **Enable** for secret scanning
   - Click **Enable** for push protection (recommended)
3. Configure notifications:
   - Navigate to **Settings** → **Notifications**
   - Ensure security alerts are enabled

### 5. Require Signed Commits (Issue #14)

This setting is configured as part of branch protection (Step 1 above).

Additionally, you may want to:
1. Document signing requirements in contribution guidelines
2. Provide team members with GPG/SSH signing setup instructions
3. Monitor commit signature compliance

### 6. Configure Dependabot (Already Configured)

The repository already has:
- ✅ Dependabot configuration (`.github/dependabot.yml`)
- ✅ Dependency Review workflow (`.github/workflows/dependency-review.yml`)

Verify these are active:
1. Navigate to **Settings** → **Code security and analysis**
2. Ensure **Dependabot alerts** is enabled
3. Ensure **Dependabot security updates** is enabled

## Verification Steps

After completing the setup, verify the configuration:

### Branch Protection
```bash
# Check if branch is protected
gh api repos/{owner}/{repo}/branches/main/protection
```

### Security Features Status
1. Navigate to **Security** → **Overview**
2. Verify all security features show as "Enabled"
3. Check for any open security alerts

### Code Scanning
1. Navigate to **Security** → **Code scanning**
2. Verify CodeQL analysis is running
3. Check for any detected vulnerabilities

### Secret Scanning
1. Navigate to **Security** → **Secret scanning**
2. Verify the feature is active
3. Review any detected secrets

## Additional Recommendations

### Security Policies
- ✅ SECURITY.md file created with vulnerability reporting procedures
- ✅ CODEOWNERS file created for code review enforcement

### Monitoring
- Set up notifications for security alerts
- Regularly review security dashboard
- Monitor Dependabot pull requests

### Team Education
- Train team members on commit signing
- Document security practices in contribution guidelines
- Conduct regular security reviews

## Compliance

These settings help meet SLSA v1.2-rc2 requirements:
- **Source**: Branch protection and signed commits
- **Build**: Code scanning and dependency review
- **Provenance**: Signed commits and protected branches
- **Dependencies**: Dependency scanning and updates

## Support

For questions or issues with security configuration:
- Review GitHub documentation: https://docs.github.com/en/code-security
- Contact repository maintainers
- Refer to SECURITY.md for vulnerability reporting

## Closing Security Issues

Once all settings are configured, verify by:
1. Running a test pull request through the full CI/CD pipeline
2. Checking that all status checks pass
3. Verifying security scans complete successfully

Then close the following issues:
- Issue #9: Branch Protection
- Issue #10: CODEOWNERS (already resolved)
- Issue #11: Code Scanning
- Issue #12: Secret Scanning
- Issue #13: Dependency Review (already resolved)
- Issue #14: Signed Commits
- Issue #17: Main security audit issue
